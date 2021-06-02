/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-02 23:00:00
 * @LastEditors: lly
 * @LastEditTime: 2021-06-02 23:12:55
 */
package gateway

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	pool []string
}

func (b *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param invaild")
	}
	b.pool = append(b.pool, params[0])
	return nil
}

func (b *RandomBalance) Next() string {
	if len(b.pool) == 0 {
		return ""
	}
	return b.pool[rand.Intn(len(b.pool))]
}

func (b *RandomBalance) Get(string) (string, error) {
	return b.Next(), nil
}
