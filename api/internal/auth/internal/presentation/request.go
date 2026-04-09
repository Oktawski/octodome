package http

type AuthenticateCredentialsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SendMagicCodeRequest struct {
	Email string `json:"email"`
}

type AuthenticateMagicCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
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
