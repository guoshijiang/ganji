//自定义模板函数
package template

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	beego.AddFuncMap("TimeForFormat", TimeForFormat)
	beego.AddFuncMap("WalletRecordType", WalletRecordType)
	beego.AddFuncMap("WalletRecordIsHandle", WalletRecordIsHandle)
	beego.AddFuncMap("WalletRecordSource", WalletRecordSource)
	beego.AddFuncMap("WalletRecordStatus", WalletRecordStatus)
}

func TimeForFormat(t interface{}) string {
	timeLayout := "2006-01-02 15:04:05"
	return t.(time.Time).Format(timeLayout)
}

//资金类型
func WalletRecordType(t int8) string{
	switch t {
	case 0:
		return "充值"
	case 1:
		return "提现"
	case 2:
		return "积分兑换"
	case 3:
		return "消费"
	default:
		return "未知"
	}
}

//资金处理状态
func WalletRecordIsHandle(t int8) string{
	switch t {
	case 0:
		return "审核中"
	case 1:
		return "审核通过"
	case 2:
		return "已打款"
	case 3:
		return "审核拒绝"
	default:
		return "未知"
	}
}

//来源平台 0：支付宝 1:微信; 2:积分兑换
func WalletRecordSource(t int8) string{
	switch t {
	case 0:
		return "支付宝"
	case 1:
		return "微信"
	case 2:
		return "积分兑换"
	default:
		return "未知"
	}
}

//来源平台 0:入账中; 1:入账成功; 2:入账失败
func WalletRecordStatus(t int8) string{
	switch t {
	case 0:
		return "入账中"
	case 1:
		return "入账成功"
	case 2:
		return "入账失败"
	default:
		return "未知"
	}
}
