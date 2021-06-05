/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-05 21:00:25
 * @LastEditors: lly
 * @LastEditTime: 2021-06-05 23:46:31
 */
package log

import (
	"os"

	"github.com/go-kit/kit/log"
)

var logger log.Logger

func init() {
	// 创建日志组件
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "Ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "Func", log.DefaultCaller)
}

func Debug(keyvals ...interface{}) {
	logger.Log(keyvals...)
}

func Info(keyvals ...interface{}) {
	logger.Log(keyvals...)
}

func Error(keyvals ...interface{}) {
	logger.Log(keyvals...)
}
