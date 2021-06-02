/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-02 23:09:26
 * @LastEditors: lly
 * @LastEditTime: 2021-06-02 23:21:55
 */

package gateway

import (
	"fmt"
	"testing"
)

func TestWhileBalance(t *testing.T) {
	t.Log("TestWhileBalance begin")

	balance := WhileBalance{}
	balance.Add("127.0.0.1:8800")
	balance.Add("127.0.0.1:8801")
	balance.Add("127.0.0.1:8802")
	balance.Add("127.0.0.1:8803")
	balance.Add("127.0.0.1:8804")

	for i := 0; i < 1111; i++ {
		fmt.Println(balance.Get(""))
	}

	t.Log("TestWhileBalance end")
}
