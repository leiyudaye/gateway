/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-14 02:47:47
 * @LastEditors: lly
 * @LastEditTime: 2021-06-14 02:48:07
 */
package lib

type GatewayParams struct {
	Comm   interface{} `form:"comm" json:"comm"`
	Module ModuleBody  `form:"module" json:"module"`
}
type ModuleBody struct {
	Module string      `form:"module" json:"module"`
	Method string      `form:"method" json:"method"`
	Param  interface{} `form:"param" json:"param"`
}
