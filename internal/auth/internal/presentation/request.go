package http

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AssignRoleRequest struct {
	Role   string `json:"role"`
	UserID uint   `json:"user_id"`
}
