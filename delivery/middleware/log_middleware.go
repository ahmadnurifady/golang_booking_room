package middleware

import (
	"final-project-booking-room/utils/common"
	"final-project-booking-room/utils/modelutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogMiddleware interface {
	LogRequest() gin.HandlerFunc
}

type logMiddleware struct {
	logService common.MyLogger
}

func (l *logMiddleware) LogRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := l.logService.InitializeLogger()
		if err != nil {
			logrus.Error("err:", err.Error())
		}
		t := time.Now()

		ctx.Next()

		logString := modelutil.RequestLog{
			AcessTime: t,
			Latency:   time.Since(t),
			ClientIP:  ctx.ClientIP(),
			Method:    ctx.Request.Method,
			Code:      ctx.Writer.Status(),
			Path:      ctx.Request.URL.Path,
			UserAgent: ctx.Request.UserAgent(),
		}

		switch {
		case ctx.Writer.Status() >= 500:
			l.logService.LogError(logString)
		case ctx.Writer.Status() >= 400:
			l.logService.LogWarn(logString)
		default:
			l.logService.LogInfo(logString)
		}
	}
}

func NewLogMiddleware(logService common.MyLogger) LogMiddleware {
	return &logMiddleware{logService: logService}
}
