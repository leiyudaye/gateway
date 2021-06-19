/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-05 21:00:25
 * @LastEditors: lly
 * @LastEditTime: 2021-06-19 23:50:22
 */
package log

import (
	"log"
)

type GoLogger struct {
}

func (l *GoLogger) Debug(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *GoLogger) Info(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *GoLogger) Error(format string, v ...interface{}) {
	log.Printf(format, v...)
}
