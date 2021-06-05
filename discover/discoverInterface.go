/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-06 00:01:53
 * @LastEditors: lly
 * @LastEditTime: 2021-06-06 00:10:26
 */
package discover

type DiscoverClient interface {
	Register(serviceName, intstanceID, healthCheckUrl string, instanceHost string,
		instancePort int, meta map[string]string) bool

	DeRegister(instanceID string) bool

	Discover(serviceName string) (string, error)
}
