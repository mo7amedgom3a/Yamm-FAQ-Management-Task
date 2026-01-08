package dto

import "github.com/google/uuid"

type SignupRequest struct {
	Email    string
	Password string
	Role     string
}

type LoginRequest struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID    uuid.UUID
	Email string
	Role  string
}
