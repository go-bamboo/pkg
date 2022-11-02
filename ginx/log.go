package ginx

// import (
// 	"fmt"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-bamboo/pkg/log"
// )

// // extractArgs returns the string of the req
// func extractArgs(req interface{}) string {
// 	if stringer, ok := req.(fmt.Stringer); ok {
// 		return stringer.String()
// 	}
// 	return fmt.Sprintf("%+v", req)
// }

// func Logger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Start timer
// 		start := time.Now()
// 		path := c.Request.URL.Path
// 		// raw := c.Request.URL.RawQuery

// 		// Process request
// 		c.Next()

// 		log.Infow(
// 			"logger",
// 			"kind", "client",
// 			"path", path,
// 			"method", c.Request.Method,
// 			// "args", extractArgs(c.),
// 			"code", c.Writer.Status(),
// 			// "reason", reason,
// 			// "stack", stack,
// 			"latency", time.Since(start).Seconds(),
// 		)
// 	}
// }
