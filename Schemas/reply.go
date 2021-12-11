package Schemas

type (
	Reply struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
		Error  string      `json:"error"`
	}

	Data struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Done      bool   `json:"status"`
		StartTime int    `json:"startTime"`
		EndTime   int    `json:"endTime"`
	}
)
