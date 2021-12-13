package Schemas

type (
	Reply struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
		Error  string      `json:"error"`
	}
)
