package repositories

import (
	"context"

	"flypro-assessment/internal/models"
)

type ReportRepository interface {
	CreateExpenseReport(ctx context.Context, instance *models.ExpenseReport) (*models.ExpenseReport, error)
	AddExpensesToReport(ctx context.Context, expenses []models.Expense) (*models.ExpenseReport, error)
	ListExpenseReports(ctx context.Context, page int) ([]models.ExpenseReport, error)
	SubmitReportForApproval(ctx context.Context, id int) (*models.ExpenseReport, error)
}

type reportRepository struct {
	db *DBConnection
}

func (r reportRepository) CreateExpenseReport(ctx context.Context, instance *models.ExpenseReport) (*models.ExpenseReport, error) {
	return nil, nil
}
func (r reportRepository) AddExpensesToReport(ctx context.Context, expenses []models.Expense) (*models.ExpenseReport, error) {
	return nil, nil
}
func (r reportRepository) ListExpenseReports(ctx context.Context, page int) ([]models.ExpenseReport, error) {
	return nil, nil
}
func (r reportRepository) SubmitReportForApproval(ctx context.Context, id int) (*models.ExpenseReport, error) {
	return nil, nil
}

func NewReportRepository(db *DBConnection) ReportRepository {
	return &reportRepository{
		db: db,
	}
}
