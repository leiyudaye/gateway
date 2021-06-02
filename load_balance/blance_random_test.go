/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-02 23:09:26
 * @LastEditors: lly
 * @LastEditTime: 2021-06-02 23:13:07
 */

package gateway

import (
	"fmt"
	"testing"
)

func TestRandomBalance(t *testing.T) {
	t.Log("TestRandomBalance begin")

	balance := RandomBalance{}
	balance.Add("127.0.0.1:8800")
	balance.Add("127.0.0.1:8801")
	balance.Add("127.0.0.1:8802")
	balance.Add("127.0.0.1:8803")
	balance.Add("127.0.0.1:8804")

	for i := 0; i < 5; i++ {
		fmt.Println(balance.Get(""))
	}

	t.Log("TestRandomBalance end")
}
