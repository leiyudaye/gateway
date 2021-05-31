/*
 * @Descripttion:
	http反向代理，使用httputil.
 * @Author: lly
 * @Date: 2021-05-31 19:43:36
 * @LastEditors: lly
 * @LastEditTime: 2021-05-31 20:45:17
*/

package gateway

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
)

func NewHttpReverseProxy(client *api.Client, logger log.Logger) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		reqPath := req.URL.Path
		if reqPath == "" {
			return
		}
		pathArray := strings.Split(reqPath, "/")
		srviceName := pathArray[1]
		result, _, err := client.Catalog().Service(srviceName, "", nil)
		if err != nil {
			logger.Log("ReverseProxy failed", "query service instace error", err.Error())
			return
		}

		if len(result) == 0 {
			logger.Log("ReverseProxy failed", "not service instace error", srviceName)
			return
		}

		destPath := strings.Join(pathArray[2:], "/")

		tgt := result[rand.Int()%len(result)]
		logger.Log("service id", tgt.ServiceID)

		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", tgt.Address, tgt.ServicePort)
		req.URL.Path = "/" + destPath
	}
	return &httputil.ReverseProxy{Director: director}
}
