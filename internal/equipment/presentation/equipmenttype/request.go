package http

type CreateRequest struct {
	Name string `json:"name"`
}

type UpdateRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
