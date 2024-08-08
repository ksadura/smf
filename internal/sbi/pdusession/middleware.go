package pdusession

import (
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

var (
	RequestCount int
	Rate         float64
	Mutex        sync.Mutex
)

func CalcRateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" && strings.HasSuffix(c.FullPath(), "/sm-contexts") {
			Mutex.Lock()
			RequestCount++
			Mutex.Unlock()
		}
		c.Next()
	}
}
