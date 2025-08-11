package models

import "time"

type User struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	Login            string     `json:"login"`
	Hash             string     `json:"hash"`
	Phone            *string    `json:"phone"`
	Active           bool       `json:"active"`
	IsSuperadmin     bool       `json:"is_superadmin"`
	CreatedBy        *int       `json:"created_by"`
	UpdatedBy        *int       `json:"updated_by"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	OutletID         *int       `json:"outlet_id"`
	LastVisitDate    *time.Time `json:"last_visit_date"`
	Language         *string    `json:"language"`
	CanLogin         bool       `json:"can_login"`
	OwnerID          *int       `json:"owner_id"`
	TeamID           *int       `json:"team_id"`
	Order            *int       `json:"order"`
	SalaryCurrencyID *int       `json:"salary_currency_id"`
	Salary           *float64   `json:"salary"`
	Type             *string    `json:"type"`
}
