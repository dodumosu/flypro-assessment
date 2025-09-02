package models

import "time"

type ExpenseStatus string

const (
	ExpenseStatusPending  ExpenseStatus = "pending"
	ExpenseStatusApproved ExpenseStatus = "approved"
	ExpenseStatusRejected ExpenseStatus = "rejected"
)

type Expense struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Currency    string  `gorm:"not null"`
	Category    string  `gorm:"not null"`
	Description string
	Receipt     string
	Status      ExpenseStatus `gorm:"default:'pending'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User `gorm:"foreignKey:UserID"`
}

type ExpenseListFilter struct {
	Category string
	Page     int
	Status   string
}
