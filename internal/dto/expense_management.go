package dto

import "time"

type ExpenseCreateDTO struct {
	// commented out because it doesn't make sense to send this like this
	// unforunately, the exercise doesn't require tokens
	// UserID      int     `json:"user_id" required:"true"`
	Amount      float64 `json:"amount" exclusiveMinimum:"0" required:"true"`
	Currency    string  `json:"currency" required:"true"`
	Category    string  `json:"category" required:"true"`
	Description string  `json:"description"`
	Receipt     string  `json:"receipt" format:"uri"`
	// would a user create a rejected expense off the bat?
	// Status      string  `json:"status" enum:"approved,pending,rejected"`
}

type ExpenseDetailDTO struct {
	ID          int       `json:"id" doc:"ID of the expense"`
	UserID      int       `json:"user_id" doc:"ID of the user who created the expense"`
	Amount      float64   `json:"amount" doc:"The expense amount"`
	Currency    string    `json:"currency" doc:"The symbol for the expense currency" example:"USD"`
	Category    string    `json:"category" doc:"The expense category"`
	Description string    `json:"description" doc:"A description for the expense"`
	Receipt     string    `json:"receipt" format:"uri" doc:"The URI/path to the receipt"`
	Status      string    `json:"status" enum:"approved,pending,rejected"`
	CreatedAt   time.Time `json:"created_at" format:"date-time" doc:"The timestamp for the expense creation"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time" doc:"The timestamp for the last expense update"`
}

type ExpenseUpdateDTO struct {
	Amount      float64 `json:"amount" exclusiveMinimum:"0" required:"true"`
	Currency    string  `json:"currency" required:"true"`
	Category    string  `json:"category" required:"true"`
	Description string  `json:"description"`
	Receipt     string  `json:"receipt" format:"uri"`
	// who approves/rejects the expenses?
	// looks like a multi-user scenario, but no roles are specified :(
	// Status      string  `json:"status" enum:"approved,pending,rejected"`
}

type ExpenseCreateResponse struct {
	Body ExpenseDetailDTO
}

type ExpenseDetailResponse struct {
	Body ExpenseDetailDTO
}

type ExpenseListRequestParams struct {
	Category string `query:"category"`
	Status   string `query:"status" enum:"approved,pending,rejected"`
}

type ExpenseListRequest struct {
	ExpenseListRequestParams
	PaginationParams
}

type ExpenseListResponseDTO struct {
	Data     []ExpenseDetailDTO `json:"expenses"`
	PageInfo `json:"page_info"`
}

type ExpenseListResponse struct {
	Body ExpenseListResponseDTO
}

type ExpenseDetailRequest struct {
	ID int `path:"id"`
}

type ExpenseUpdateRequest struct {
	ID   int              `path:"id"`
	Body ExpenseUpdateDTO `json:"expense"`
}
