package dto

import "time"

type CreateExpenseReportDTO struct {
	Title string `json:"title" doc:"Title of the expense report"`
	// might be better served by using just the IDs, but whatevs
	Expenses []ExpenseDetailDTO `json:"expenses" doc:"the expenses to add to the report"`
}

type CreateExpenseReportRequest struct {
	Body CreateExpenseReportDTO
}

type ExpenseReportDetailResponse struct {
	Body ExpenseDetailDTO
}

type AddExpensesToReportDTO struct {
	Expenses []ExpenseDetailDTO
}

type AddExpensesToReportRequest struct {
	ID   int                    `path:"id" doc:"ID of the report to add expenses to"`
	Body AddExpensesToReportDTO `json:"expenses" doc:"the expenses to add to the report"`
}

type ExpenseReportDTO struct {
	ID        int           `json:"id" doc:"ID of the expense report"`
	UserID    int           `json:"user_id" doc:"The ID of the owner of the expense report"`
	Title     string        `json:"title" doc:"The title of the expense report"`
	Status    string        `json:"status" enum:"approved,draft,submitted" doc:"Status of the report"`
	Total     float64       `json:"total" doc:"The total amount of the expenses, in USD"`
	CreatedAt time.Time     `json:"created_at" format:"date-time" doc:"The timestamp the report was created"`
	UpdatedAt time.Time     `json:"updated_at" format:"date-time" doc:"The last time the report was updated"`
	User      UserDetailDTO `json:"user" doc:"The report owner"`
}

type ExpenseReportListDTO struct {
	PageInfo
	Reports []ExpenseReportDTO `json:"reports" doc:"The expense reports"`
}

type ExpenseReportListRequest struct {
	PaginationParams
}

type ExpenseReportListResponse struct {
	Body ExpenseReportListDTO
}

type ExpenseReportApprovalRequest struct {
	ID int `path:"id" doc:"ID of the report submitted for approval"`
}
