// models/cashback_filters.go
package models

import "time"

type CashbackHistoryFilters struct {
	OutletIDs   []int      `form:"outlet_ids[]" json:"outlet_ids"`
	IsCancelled *bool      `form:"is_cancelled" json:"is_cancelled"`
	DateFrom    *time.Time `form:"date_from" json:"date_from"`
	DateTo      *time.Time `form:"date_to" json:"date_to"`
	Page        int        `form:"page" json:"page"`
	PerPage     int        `form:"per_page" json:"per_page"`
}
