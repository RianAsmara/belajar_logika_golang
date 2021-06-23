package webservice

// standarize response body

type ResBody struct {
	Message string  `json:"message"`
	Data    float64 `json:"data"`
	Status  int     `json:"status"`
	IsError bool    `json:"isError"`
}
