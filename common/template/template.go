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
	beego.AddFuncMap("ProcessStatus", ProcessStatus)
	beego.AddFuncMap("ProcessIsRecvGoods", ProcessIsRecvGoods)
	beego.AddFuncMap("ProcessFundRet", ProcessFundRet)
	beego.AddFuncMap("UnixTimeForFormat", UnixTimeForFormat)
}

//时间轴转时间字符串
func UnixTimeForFormat(timeUnix int) string {
	//转化所需模板
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(int64(timeUnix), 0).Format(timeLayout)
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

//退款订单状态
func ProcessStatus(t int8) string {
	switch t {
	case 0:
		return "等待卖家确认"
	case 1:
		return "卖家已同意"
	case 2:
		return "卖家拒绝"
	case 3:
		return "等待买家邮寄"
	case 4:
		return "等待卖家收货"
	case 5:
		return "卖家已经发货"
	case 6:
		return "等待买家收货"
	case 7:
		return "已完成"
	default:
		return "未知"
	}
}

func ProcessIsRecvGoods(t int8) string {
	switch t {
	case 0:
		return "未收到货物"
	case 1:
		return "已收到货物"
	default:
		return "未知"
	}
}

func ProcessFundRet(t int8) string {
	switch t {
	case 0:
		return "返回到平台钱包"
	case 1:
		return "原路返回"
	default:
		return "未知"
	}
}
