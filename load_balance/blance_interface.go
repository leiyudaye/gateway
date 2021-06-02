/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-02 23:01:39
 * @LastEditors: lly
 * @LastEditTime: 2021-06-02 23:08:44
 */
package gateway

type BlanceInterface interface {
	Add(params ...string) error
	Get(string) (string, error)
}
