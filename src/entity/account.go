package entity

type AccountEntity struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint32 `json:"age"`
}
