/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-05 21:00:25
 * @LastEditors: lly
 * @LastEditTime: 2021-06-19 23:51:16
 */
package log

type LogInterface interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
}

var logger LogInterface

func init() {
	logger = new(GoLogger)
}

func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}
