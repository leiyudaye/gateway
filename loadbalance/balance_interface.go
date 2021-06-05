/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-02 23:01:39
 * @LastEditors: lly
 * @LastEditTime: 2021-06-05 21:12:05
 */
package loadbalance

const (
	BlanceTypeRandom = 1
	BlanceTypeWhile  = 2
)

type BalanceInterface interface {
	Add(params ...string) error
	Get(string) (string, error)
}

func NewLoadBlance(bleType int) BalanceInterface {
	switch bleType {
	case BlanceTypeRandom:
		return new(RandomBalance)
	case BlanceTypeWhile:
		return new(WhileBalance)
	default:
		return nil
	}
}
