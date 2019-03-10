package model

type TestRequest struct {
	ID int `form:"id"`
}

type TestResponse struct {
	Data string `json:"data"`
}
