package entity

type AccountEntity struct {
	Id       uint64 `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint32 `json:"age"`
}
