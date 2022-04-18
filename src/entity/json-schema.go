package entity

type (
	RequestAccountLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RequestAccountRegister struct {
		Email    string `json:"email"`
		Nama     string `json:"nama"`
		Password string `json:"password"`
		Age      uint32 `json:"age"`
	}
)
