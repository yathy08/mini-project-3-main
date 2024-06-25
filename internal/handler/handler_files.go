package handler

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Users struct {
	Data []User `json:"data"`
}
