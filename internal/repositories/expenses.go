package repositories

import (
	"context"

	"flypro-assessment/internal/models"
)

type ExpenseRepository interface {
	CreateExpense(ctx context.Context, input *models.Expense) (*models.Expense, error)
	ListExpenses(ctx context.Context, filter models.ExpenseListFilter) ([]models.Expense, error)
	GetExpense(ctx context.Context, id int) (*models.Expense, error)
	UpdateExpense(ctx context.Context, id int, input *models.Expense) (*models.Expense, error)
	DeleteExpense(ctx context.Context, id int) error
}

type expenseRepository struct {
	db *DBConnection
}

func (e expenseRepository) CreateExpense(ctx context.Context, instance *models.Expense) (*models.Expense, error) {
	return nil, nil
}

func (e expenseRepository) ListExpenses(ctx context.Context, filter models.ExpenseListFilter) ([]models.Expense, error) {
	return nil, nil
}

func (e expenseRepository) GetExpense(ctx context.Context, id int) (*models.Expense, error) {
	return nil, nil
}

func (e expenseRepository) UpdateExpense(ctx context.Context, id int, instance *models.Expense) (*models.Expense, error) {
	return nil, nil
}

func (e expenseRepository) DeleteExpense(ctx context.Context, id int) error {
	return nil
}

func NewExpenseRepository(db *DBConnection) ExpenseRepository {
	return &expenseRepository{
		db: db,
	}
}
