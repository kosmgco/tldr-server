package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	REQUEST_ID_NAME = "x-request-id"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		reqID := c.Request.Header.Get(REQUEST_ID_NAME)

		c.Set(REQUEST_ID_NAME, reqID)

		fields := logrus.Fields{
			"tag":       "access",
			"log_id":    reqID,
			"remote_ip": c.ClientIP(),
			"method":    c.Request.Method,
			"pathname":  c.Request.URL.Path,
		}
		logrus.SetFormatter(&logrus.JSONFormatter{})

		c.Next()

		fields["status"] = c.Writer.Status()
		fields["request_time"] = time.Now().Sub(now).String()

		logger := logrus.WithFields(fields)

		if len(c.Errors) > 0 {
			logger.WithField("errors", c.Errors).Warningf("")
		} else {
			logger.Info("")
		}
	}
}
