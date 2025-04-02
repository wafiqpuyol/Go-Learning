package types

type Student struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
