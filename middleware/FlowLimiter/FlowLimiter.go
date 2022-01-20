package FlowLimiter

import (
	"github.com/gin-gonic/gin"
	"time"
)

type value int

const (
	FlagUser     value = 1 << iota
	FlagMethod   value = 1 << iota
	FlagClientIP value = 1 << iota
)

type Limiter struct {
	Rate struct {
		Frequency int
		Duration  time.Duration
	}
	Policy value
}

func FlowLimiter() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
