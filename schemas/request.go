package schemas

type (
	LoginToken struct {
		Userid   int    `json:"userid" binding:"required"`
		Passcode string `json:"passcode" binding:"required"`
	}
	User struct {
		Username string `json:"username" binding:"required"`
		Passcode string `json:"passcode" binding:"required"`
	}
	TodoAdding struct {
		Title     string `json:"title" binding:"required"`
		Content   string `json:"content"`
		StartTime int    `json:"startTime"`
		Deadline  int    `json:"deadline"`
	}
	TodoDeleting struct {
		Id   int `form:"id"`
		Type int `form:"type"`
	}
	TodoGetting struct {
		Id      int    `form:"id"`
		Type    int    `form:"type"`
		Keyword string `form:"keyword"`
		Page    int    `form:"page"`
	}
	TodoDone struct {
		Id int `form:"id"`
	}
)
