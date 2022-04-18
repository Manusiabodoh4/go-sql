package entity

type TemplateResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TemplateChannelResponse struct {
	Data  interface{}
	Error error
}
