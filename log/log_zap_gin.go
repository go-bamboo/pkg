package log

//func (s *ZapLogger) Gin(conf gin.LoggerConfig) gin.HandlerFunc {
//	logger := s.Skip(3)
//	notlogged := conf.SkipPaths
//
//	var skip map[string]struct{}
//	if length := len(notlogged); length > 0 {
//		skip = make(map[string]struct{}, length)
//
//		for _, path := range notlogged {
//			skip[path] = struct{}{}
//		}
//	}
//	return func(c *gin.Context) {
//		// Start timer
//		start := time.Now()
//		path := c.Request.URL.Path
//		raw := c.Request.URL.RawQuery
//
//		// Process request
//		c.Next()
//
//		// Log only when path is not being skipped
//		if _, ok := skip[path]; !ok {
//			param := gin.LogFormatterParams{
//				Request: c.Request,
//				//isTerm:  isTerm,
//				Keys: c.Keys,
//			}
//
//			// Stop timer
//			param.TimeStamp = time.Now()
//			param.Latency = param.TimeStamp.Sub(start)
//
//			param.ClientIP = c.ClientIP()
//			param.Method = c.Request.Method
//			param.StatusCode = c.Writer.Status()
//			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
//
//			param.BodySize = c.Writer.Size()
//
//			if raw != "" {
//				path = path + "?" + raw
//			}
//
//			param.Path = path
//
//			logger.Log(log.LevelInfo, "path", param.Path)
//		}
//	}
//}
