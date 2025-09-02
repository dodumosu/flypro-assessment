package dto

import "time"

type UserCreateDTO struct {
	Email string `json:"email" required:"true" format:"email"`
	Name  string `json:"name" required:"true"`
}

type UserCreateRequest struct {
	Body UserCreateDTO
}

type UserCreateResponse struct {
	Body   UserDetailDTO
	Status int
}

type UserDetailRequest struct {
	ID int `path:"id"`
}

type UserDetailDTO struct {
	ID        int       `json:"id"`
	Email     string    `json:"email" format:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDetailResponse struct {
	Body UserDetailDTO
}
