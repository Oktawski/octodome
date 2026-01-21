package http

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AssignRoleRequest struct {
	Role   string `json:"role"`
	UserID uint   `json:"user_id"`
}

type UnassignRoleRequest struct {
	Role   string `json:"role"`
	UserID uint   `json:"user_id"`
}

type SyncRolesRequest struct {
	Roles  []string `json:"roles"`
	UserID uint     `json:"user_id"`
}
