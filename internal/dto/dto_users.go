package dto

import "time"

// response
type UserResponse struct {
	ID          string     `json:"id"`
	TenantID    string     `json:"tenant_id"`
	TenantName  string     `json:"tenant_name"`
	RoleID      string     `json:"role_id"`
	RoleName    string     `json:"role_name"`
	FullName    string     `json:"full_name"`
	Username    string     `json:"username"`
	Phone       string     `json:"phone"`
	IsActive    bool       `json:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at"`
}

// create request
type CreateUserRequest struct {
	RoleID   string `json:"role_id" binding:"required,uuid"`
	FullName string `json:"full_name" binding:"required,min=3,max=150"`
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required,min=8,max=30"`
	IsActive bool   `json:"is_active"`
}

// update request
type UpdateUserRequest struct {
	RoleID   *string `json:"role_id" binding:"omitempty,uuid"`
	FullName *string `json:"full_name" binding:"omitempty,min=3,max=150"`
	Username *string `json:"username" binding:"omitempty,min=3,max=100"`
	// Password *string `json:"password" binding:"omitempty"` -> endpoint update password sebaiknya dipisahkan
	Phone    *string `json:"phone" binding:"omitempty,min=8,max=30"`
	IsActive *bool   `json:"is_active" binding:"omitempty"`
}
