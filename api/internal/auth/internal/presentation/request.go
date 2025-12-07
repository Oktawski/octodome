package http

type AuthRequest struct {
	Username string `json:"username"`
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
