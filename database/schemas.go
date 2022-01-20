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
		View      int    `json:"view"`
		CreatedAt int    `json:"createdAt" gorm:"autoCreateTime:milli"`
		StartTime int    `json:"startTime"`
		Deadline  int    `json:"deadline"`
		UserId    int    `gorm:"column:userid"`
	}
)
