/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-05 20:31:49
 * @LastEditors: lly
 * @LastEditTime: 2021-06-09 23:36:25
 */
package discover

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/leiyudaye/gateway/loadbalance"
	"github.com/leiyudaye/gateway/log"
)

type ConsulDiscoverClient struct {
	client consul.Client
	ld     loadbalance.BalanceInterface
}

func NewDiscoverClient(addr string) (DiscoverClient, error) {
	conf := api.DefaultConfig()
	conf.Address = addr
	apiClient, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &ConsulDiscoverClient{
		client: client,
		ld:     loadbalance.NewLoadBlance(loadbalance.BlanceTypeWhile),
	}, nil
}

// 服务注册
func (consulClient *ConsulDiscoverClient) Register(serviceName, instanceId, healthCheckUrl, instanceHost string,
	instancePort int, meta map[string]string) bool {
	// 构建服务实例元数据
	registration := &api.AgentServiceRegistration{
		ID:      instanceId,
		Name:    serviceName,
		Address: instanceHost,
		Port:    instancePort,
		Meta:    meta,
		// Check: &api.AgentServiceCheck{
		// 	DeregisterCriticalServiceAfter: "30s",
		// 	HTTP:                           healthCheckUrl,
		// 	Interval:                       "15s",
		// },
	}

	// 注册服务
	err := consulClient.client.Register(registration)
	if err != nil {
		log.Error("Register Service Error", err)
		return false
	}
	log.Info("Register Service Success")
	return true
}

// 服务注销
func (c *ConsulDiscoverClient) DeRegister(instanceID string) bool {
	return true
}

// 服务发现
func (c *ConsulDiscoverClient) Discover(serviceName string) (string, error) {
	// 以http的形式去注册中心拉取服务配置
	result, _, err := c.client.Service(serviceName, "", false, nil)
	if err != nil {
		log.Error("discover service failed, err=%v", err)
		return "", errors.New("discover service failedd")
	}

	if len(result) == 0 {
		log.Error("no found server, serverName=%v", serviceName)
		return "", errors.New("no found")
	}

	// 把服务地址添加到负载均衡器中
	for _, svr := range result {
		log.Debug(fmt.Sprintf("%s:%d", svr.Service.Address, svr.Service.Port))
		c.ld.Add(fmt.Sprintf("%s:%d", svr.Service.Address, svr.Service.Port))
	}

	return c.ld.Get("")
}
