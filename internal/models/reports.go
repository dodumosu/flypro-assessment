package models

import "time"

type ReportStatus string

const (
	ReportStatusDraft     ReportStatus = "draft"
	ReportStatusSubmitted ReportStatus = "submitted"
	ReportStatusApproved  ReportStatus = "approved"
)

type ExpenseReport struct {
	ID        uint         `gorm:"primaryKey"`
	UserID    uint         `gorm:"not null"`
	Title     string       `gorm:"not null"`
	Status    ReportStatus `gorm:"default:'draft'"`
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User      `gorm:"foreignKey:UserID"`
	Expenses  []Expense `gorm:"many2many:report_expenses"`
}
