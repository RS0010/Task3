package schemas

type (
	Reply struct {
		Status int         `json:"status" binding:"required"`
		Data   interface{} `json:"data" binding:"required"`
		Error  string      `json:"error" binding:"required"`
	}
	TokenGetting struct {
		AuthToken    string `json:"authToken" binding:"required"`
		RefreshToken string `json:"refreshToken" binding:"required"`
	}
	UserGetting struct {
		Userid   int    `json:"userid"`
		Username string `json:"username"`
	}
)
