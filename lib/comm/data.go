/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-14 02:47:47
 * @LastEditors: lly
 * @LastEditTime: 2021-06-20 00:32:14
 */
package comm

type GatewayParams struct {
	Comm   interface{} `form:"comm" json:"comm"`
	Module ModuleBody  `form:"moduleKey" json:"moduleKey"`
}
type ModuleBody struct {
	Module string      `form:"module" json:"module"`
	Method string      `form:"method" json:"method"`
	Param  interface{} `form:"param" json:"param"`
}

type GatewayRsp struct {
	Data string // json数据
	Code int32  // 错误码
}
