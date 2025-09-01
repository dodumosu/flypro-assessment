package handlers

const (
	HealthCheckPath      = "/health"
	APIPrefix            = "/api"
	DocsPath             = APIPrefix + "/docs"
	UserListPath         = APIPrefix + "/users"
	UserCreatePath       = UserListPath
	UserDetailPath       = UserListPath + "/{id}"
	ExpenseCreatePath    = APIPrefix + "/expenses"
	ExpenseListPath      = ExpenseCreatePath
	ExpenseDetailPath    = ExpenseListPath + "/{id}"
	ExpenseUpdatePath    = ExpenseDetailPath
	ExpenseDeletePath    = ExpenseDetailPath
	ReportListPath       = APIPrefix + "/reports"
	ReportCreatePath     = ReportListPath
	ReportApprovalPath   = ReportListPath + "/{id}/submit"
	ReportExpenseAddPath = ReportListPath + "/{id}/expenses"
)
