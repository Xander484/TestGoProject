package handler

import (
	"fmt"
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CashbackReportHandler) UserHandler(c *gin.Context) {
	rows, err := h.Service.DB.Query(`
        SELECT id, name, login, hash, phone, active, is_superadmin, created_by, updated_by, created_at, updated_at, deleted_at, outlet_id, last_visit_date, language, can_login, owner_id, team_id, "order", salary_currency_id, salary, type
        FROM users
    `)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Login,
			&user.Hash,
			&user.Phone,
			&user.Active,
			&user.IsSuperadmin,
			&user.CreatedBy,
			&user.UpdatedBy,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.OutletID,
			&user.LastVisitDate,
			&user.Language,
			&user.CanLogin,
			&user.OwnerID,
			&user.TeamID,
			&user.Order,
			&user.SalaryCurrencyID,
			&user.Salary,
			&user.Type,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
			return
		}
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
	}
}
