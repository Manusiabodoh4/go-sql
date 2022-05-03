package entity

type (
	RequestAccountLogin struct {
		Email    string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}

	RequestAccountRegister struct {
		Email    string `json:"email"`
		Nama     string `json:"nama"`
		Password string `json:"password"`
		Age      uint32 `json:"age"`
	}

	RequestAccountRegisterMany struct {
		Data []AccountEntity `json:"data"`
	}
)
