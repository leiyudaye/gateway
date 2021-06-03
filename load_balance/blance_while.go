/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-02 23:00:00
 * @LastEditors: lly
 * @LastEditTime: 2021-06-02 23:20:47
 */
package gateway

import (
	"errors"
)

type WhileBalance struct {
	pool  []string
	index int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (b *WhileBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param invaild")
	}
	b.pool = append(b.pool, params[0])
	return nil
}

func (b *WhileBalance) Next() string {
	if len(b.pool) == 0 {
		return ""
	}
	addr := b.pool[abs(b.index%len(b.pool))]
	b.index++
	return addr
}

func (b *WhileBalance) Get(string) (string, error) {
	return b.Next(), nil
}
