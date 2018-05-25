package testFileDir

import "errorDocPrinter/util"

var (
	ParamsError    		= util.GetCommonErr(3110,"error params")
	SqlUpdateError 		= util.GetCommonErr(3111,"update failed")
	YellowError    		= util.GetCommonErr(3112,"yellow 内容涉黄")
	ForbidError    		= util.GetCommonErr(3113,"forbid 禁止访问")
	EmptyIdError   		= util.GetCommonErr(3114,"empty id")
	SqlOpenTraError		= util.GetCommonErr(3115,"服务端开启事务失败")
	SqlCommitTraError	= util.GetCommonErr(3116,"服务端事务提交失败")
	EffectLessError 	= util.GetCommonErr(3117,"update effect row <= 0")
	SqlGetEffectError   = util.GetCommonErr(3118,"RowsAffected 失败")
	SqlPartSuccessError = util.GetCommonErr(3119,"更新只有部分成功")

	EmptyUserIdError    = util.GetCommonErr(3120,"empty userId")
	TooLagerError       = util.GetCommonErr(3121,"too lager")

	UserNotExitError    = util.GetCommonErr(3122,"user not exits")

	SqlInvalidUpdateError = util.GetCommonErr(3123,"非法更新")

	SizeTooManyError    = util.GetCommonErr(3124,"参数个数长度限制")
	SqlCommitError      = util.GetCommonErr(3126,"服务端事务提交失败")

	PayMoneyFormatError = util.GetCommonErr(3127,"invalid money")

	MoneyNotEnoughError = util.GetCommonErr(3128,"money not enough")

	CreateConsumeError  = util.GetCommonErr(3129,"创建消费记录失败")
)
