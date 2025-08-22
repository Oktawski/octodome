package user

type UserGetResponse struct {
	ID    uint
	Name  string
	Email string
}

type UserCreateRequest struct {
	Name     string
	Email    string
	Password string
}

type UserCreateResponse struct {
	ID    uint
	Name  string
	Email string
}
