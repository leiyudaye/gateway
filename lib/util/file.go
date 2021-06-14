/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-14 20:35:16
 * @LastEditors: lly
 * @LastEditTime: 2021-06-14 21:24:57
 */

package util

import (
	"io/ioutil"

	"github.com/leiyudaye/gateway/log"
)

// 获取文件路径下的所有文件，不包含目录
func GetFileNameByPath(path string) []string {
	var (
		filenames []string
	)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error("no found path, err=%v", err)
		return []string{}
	}
	for _, f := range files {
		if !f.IsDir() {
			filenames = append(filenames, path+"/"+f.Name())
		}
	}
	return filenames
}
