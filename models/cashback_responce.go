package models

type CashbackResponse struct {
	CurrentPage  int         `json:"current_page"`
	Data         []Cashback  `json:"data"`
	FirstPageURL string      `json:"first_page_url"`
	From         int         `json:"from"`
	LastPage     int         `json:"last_page"`
	LastPageURL  string      `json:"last_page_url"`
	Links        []Link      `json:"links"`
	NextPageURL  *string     `json:"next_page_url"` // nullable
	Path         string      `json:"path"`
	PerPage      int         `json:"per_page"`
	PrevPageURL  *string     `json:"prev_page_url"` // nullable
	To           int         `json:"to"`
	Total        int         `json:"total"`
	Totals       []TotalData `json:"totals"`
}

type Cashback struct {
	Type                  int     `json:"type"`
	AmountInBaseCurrency  float64 `json:"amount_in_base_currency"`
	CurrencyCode          string  `json:"currency_code"`
	IsCancelled           bool    `json:"is_cancelled"`
	CreatedAt             string  `json:"created_at"`
	OutletName            string  `json:"outlet_name"`
	OutletID              int     `json:"outlet_id"`
	DocumentID            int     `json:"document_id"`
	DocumentNumber        string  `json:"document_number"`
	DocumentDate          string  `json:"document_date"`
	DocumentStatus        int     `json:"document_status"`
	DocumentType          int     `json:"document_type"`
	DocumentCurrencyCode  string  `json:"document_currency_code"`
	DocumnetTotalCashback string  `json:"documnet_total_cashback"`
}

type Link struct {
	URL    *string `json:"url"` // nullable
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

type TotalData struct {
	Name   string        `json:"name"`
	Values []TotalAmount `json:"values"`
}

type TotalAmount struct {
	AmountInBaseCurrency float64 `json:"amount_in_base_currency"`
	CurrencyCode         string  `json:"currency_code"`
}
