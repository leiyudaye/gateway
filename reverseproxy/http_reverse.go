/*
 * @Descripttion:
	http反向代理，使用httputil.
 * @Author: lly
 * @Date: 2021-05-31 19:43:36
 * @LastEditors: lly
 * @LastEditTime: 2021-06-06 00:22:32
*/

package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/leiyudaye/gateway/discover"
)

func NewHttpReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		reqPath := req.URL.Path
		if reqPath == "" {
			return
		}
		pathArray := strings.Split(reqPath, "/")
		srviceName := pathArray[1]
		disConn, err := discover.NewDiscoverClient("127.0.0.1", 8500)
		if err != nil {
			return
		}
		tagAddr, err := disConn.Discover(srviceName)
		if err != nil {
			return
		}

		req.URL.Scheme = "http"
		req.URL.Host = tagAddr
		req.URL.Path = "/" + strings.Join(pathArray[2:], "/")
	}
	return &httputil.ReverseProxy{Director: director}
}
