package database

type (
	User struct {
		Id   int
		Name string
		Code string
	}
	Data struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Done      bool   `json:"status"`
		StartTime int    `json:"startTime" gorm:"autoCreateTime:milli"`
		Deadline  int    `json:"deadline"`
		UserId    int    `gorm:"column:userid"`
	}
)
