package dto

type ProfileResponse struct {
	ID         string `json:"id"`
	FullName   string `json:"full_name"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Phone      string `json:"phone"`
}

type UpdateProfileRequest struct {
	FullName string `json:"full_name" binding:"required,max=150"`
	Phone    string `json:"phone" binding:"min=8,max=30"`
}

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
