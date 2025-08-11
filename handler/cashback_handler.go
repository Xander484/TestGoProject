package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"myapp/models"

	"github.com/gin-gonic/gin"
)

type CashbackReportHandler struct {
	DB *sql.DB
}

// buildWhereClause builds WHERE conditions and args for queries
func buildWhereClause(filters models.CashbackHistoryFilters) (string, []interface{}) {
	whereClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if len(filters.OutletIDs) > 0 {
		placeholders := []string{}
		for _, id := range filters.OutletIDs {
			placeholders = append(placeholders, fmt.Sprintf("$%d", argPos))
			args = append(args, id)
			argPos++
		}
		whereClauses = append(whereClauses, fmt.Sprintf("documents.outlet_id IN (%s)", strings.Join(placeholders, ",")))
	}

	if filters.IsCancelled != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("cashback.is_cancelled = $%d", argPos))
		args = append(args, *filters.IsCancelled)
		argPos++
	}

	if filters.DateFrom != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("documents.date >= $%d", argPos))
		args = append(args, *filters.DateFrom)
		argPos++
	}

	if filters.DateTo != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("documents.date <= $%d", argPos))
		args = append(args, *filters.DateTo)
		argPos++
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	return whereSQL, args
}
func (h *CashbackReportHandler) CashbackHistory(c *gin.Context) {
	var filters models.CashbackHistoryFilters
	if err := c.ShouldBind(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pagination params
	page := filters.Page
	if page <= 0 {
		page = 1
	}
	perPage := filters.PerPage
	if perPage <= 0 {
		perPage = 10
	}
	offset := (page - 1) * perPage

	// WHERE clause
	whereSQL, args := buildWhereClause(filters)

	// Count query
	countQuery := `
		SELECT COUNT(*)
		FROM cashback
		LEFT JOIN documents ON documents.id = cashback.document_id
		LEFT JOIN outlets ON outlets.id = documents.outlet_id
		LEFT JOIN currencies ON documents.currency_id = currencies.id
		` + whereSQL
	var total int
	if err := h.DB.QueryRow(countQuery, args...).Scan(&total); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Data query
	argsWithLimit := append(args, perPage, offset)
	dataQuery := `
		SELECT cashback.type, cashback.amount_in_base_currency, currencies.code AS currency_code,
		       cashback.is_cancelled, cashback.created_at,
		       outlets.name AS outlet_name, outlets.id AS outlet_id,
		       documents.id AS document_id, documents.number AS document_number,
		       documents.date AS document_date, documents.status AS document_status,
		       documents.type AS document_type, currencies.code AS document_currency_code,
		       cashback.amount_in_base_currency AS document_total_cashback
		FROM cashback
		LEFT JOIN documents ON documents.id = cashback.document_id
		LEFT JOIN outlets ON outlets.id = documents.outlet_id
		LEFT JOIN currencies ON documents.currency_id = currencies.id
		` + whereSQL + `
		ORDER BY cashback.id DESC
		LIMIT $` + fmt.Sprintf("%d", len(argsWithLimit)-1) + ` OFFSET $` + fmt.Sprintf("%d", len(argsWithLimit))

	rows, err := h.DB.Query(dataQuery, argsWithLimit...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var cashbacks []models.Cashback
	for rows.Next() {
		var cb models.Cashback
		err := rows.Scan(
			&cb.Type, &cb.AmountInBaseCurrency, &cb.CurrencyCode,
			&cb.IsCancelled, &cb.CreatedAt, &cb.OutletName, &cb.OutletID,
			&cb.DocumentID, &cb.DocumentNumber, &cb.DocumentDate,
			&cb.DocumentStatus, &cb.DocumentType, &cb.DocumentCurrencyCode,
			&cb.DocumnetTotalCashback,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cashbacks = append(cashbacks, cb)
	}

	// Totals
	totals := h.getTotals(filters)

	// Pagination links
	lastPage := (total + perPage - 1) / perPage
	var prevPageURL *string
	var nextPageURL *string
	path := "/cashback-history" // change to your actual endpoint path
	if page > 1 {
		url := fmt.Sprintf("%s?page=%d", path, page-1)
		prevPageURL = &url
	}
	if page < lastPage {
		url := fmt.Sprintf("%s?page=%d", path, page+1)
		nextPageURL = &url
	}

	// Build links array
	links := []models.Link{}
	for i := 1; i <= lastPage; i++ {
		url := fmt.Sprintf("%s?page=%d", path, i)
		links = append(links, models.Link{
			URL:    &url,
			Label:  fmt.Sprintf("%d", i),
			Active: i == page,
		})
	}

	// Build final response
	response := models.CashbackResponse{
		CurrentPage:  page,
		Data:         cashbacks,
		FirstPageURL: fmt.Sprintf("%s?page=1", path),
		From:         offset + 1,
		LastPage:     lastPage,
		LastPageURL:  fmt.Sprintf("%s?page=%d", path, lastPage),
		Links:        links,
		NextPageURL:  nextPageURL,
		Path:         path,
		PerPage:      perPage,
		PrevPageURL:  prevPageURL,
		To:           offset + len(cashbacks),
		Total:        total,
		Totals:       totals,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CashbackReportHandler) getTotals(filters models.CashbackHistoryFilters) []models.TotalData {
	whereSQL, args := buildWhereClause(filters)

	query := `
        SELECT COALESCE(SUM(cashback.amount_in_base_currency), 0)
        FROM cashback
        JOIN documents ON documents.id = cashback.document_id
        ` + whereSQL

	var totalAmount float64
	_ = h.DB.QueryRow(query, args...).Scan(&totalAmount)

	return []models.TotalData{
		{
			Name: "amount_in_base_currency",
			Values: []models.TotalAmount{
				{
					AmountInBaseCurrency: totalAmount,
					CurrencyCode:         "USD", // adjust if needed
				},
			},
		},
	}
}
