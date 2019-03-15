package model

type TestRequest struct {
	ID      int `form:"id"`
	RouteID int `route:"route_id"`
}

type TestResponse struct {
	Data string `json:"data"`
}
