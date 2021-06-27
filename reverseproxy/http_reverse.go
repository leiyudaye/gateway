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
	"github.com/leiyudaye/gateway/log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/leiyudaye/gateway/discover"
)

func NewHttpReverseProxy() *httputil.ReverseProxy {
	disConn, err := discover.NewDiscoverClient("127.0.0.1:8500")
	if err != nil {
		log.Error("discover connect failed. err=%v", err)
		return nil
	}

	director := func(req *http.Request) {
		reqPath := req.URL.Path
		if reqPath == "" {
			log.Error("no found url path")
			return
		}
		pathArray := strings.Split(reqPath, "/")
		if len(pathArray) == 0 {
			log.Error("url parse failed.")
			return
		}

		tagAddr, err := disConn.Discover(pathArray[1])
		if err != nil {
			return
		}

		req.URL.Scheme = "http"
		req.URL.Host = tagAddr
		req.URL.Path = "/" + strings.Join(pathArray[2:], "/")
	}
	return &httputil.ReverseProxy{Director: director}
}
