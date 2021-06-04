/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-02 23:01:39
 * @LastEditors: lly
 * @LastEditTime: 2021-06-02 23:08:44
 */
package gateway

const (
	BlanceTypeRandom = 1
	BlanceTypeWhile  = 2
)

type BlanceInterface interface {
	Add(params ...string) error
	Get(string) (string, error)
}

func NewLoadBlance(bleType int) BlanceInterface {
	switch bleType {
	case BlanceTypeRandom:
		return new(RandomBalance)
	case BlanceTypeWhile:
		return new(WhileBalance)
	default:
		return nil
	}
}
