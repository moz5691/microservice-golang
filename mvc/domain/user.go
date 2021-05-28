package domain

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"firsName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
