/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-06 22:24:33
 * @LastEditors: lly
 * @LastEditTime: 2021-06-10 00:07:57
 */
package middleware

import (
	"sync/atomic"
	"time"

	"github.com/leiyudaye/gateway/log"
	"google.golang.org/grpc"
)

var coter *Counter

type Counter struct {
	total uint64
	count uint64
}

func (c *Counter) Increase() {
	atomic.AddUint64(&c.count, 1)
}

func (c *Counter) GetCount() uint64 {
	return c.count
}

func NewCounter() *Counter {
	if coter != nil {
		return coter
	} else {
		coter = new(Counter)

		for range time.Tick(1 * time.Second) {
			log.Debug(coter.GetCount())
			coter.total += coter.count
			log.Debug("qps=%v", coter.GetCount())
			atomic.StoreUint64(&coter.count, 0)
		}

		return coter
	}
}

func GrpcFlowCount() func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Debug(time.Now(), "+1")

		coter := NewCounter()
		coter.Increase()
		log.Debug("counter = %v", coter.GetCount())

		if err := handler(srv, ss); err != nil {
			log.Debug("GrpcFlowCountMiddleware failed with error %v\n", err)
			return err
		}

		return nil
	}
}
