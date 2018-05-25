package testFileDir

import "errorDocPrinter/util"

func CreateUser()  {
	util.GetErrWithTips(-1,"create user failed","创建用户失败")

	// 注释会被过滤
	// util.GetCommonErr(-2,"create user failed")
	 /*
	  * util.GetCommonErr(-3,"create user failed")
	  */
	util.GetCommonErr(-4,"invalid create user")
}
















