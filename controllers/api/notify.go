package api

import (
	"fmt"
	"ganji/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type NotifyController struct {
	beego.Controller
}

/*
gmt_create=2021-01-09+19%3A34%3A19&charset=utf-8&seller_email=mujiangzi2018%40163.com&
subject=%E5%B8%82%E9%9B%86%E5%95%86%E5%9F%8E&
sign=D05cx3FwaXBmNPabzto76xbsdCMbfLsBnE%2FE5sTs8Ftx8omJnAYHo6l%2F3YJV2tZj7tF5wA48fUPDkyhyNW2huHqqKigiz3LY2J3znQ26%2F906WGVGQ1GHhJBT1UIhdKUivMqHRV8E2HTAW7UF92ugkcWvNtl2eKQR%2FVMADlwkUfL0YmGBcIyR1kMFPOZClBDEzldqJELeyaRrtnMCpdMRkgnS66pLT3bjvPGf4PcOwaygskEmfBEak%2FUZequcV8iztMx4775Zvu2XQ34rqdr69PYEnyzHiMvfNu9%2FtJ%2BOyMTHA7qNbxFjUz803PYDGFRsAZupQS%2BEU5ZP8ulYnBh4wA%3D%3D
&buyer_id=2088012348746586&invoice_amount=10.00&notify_id=2021010900222193420046581454685498&
fund_bill_list=%5B%7B%22amount%22%3A%2210.00%22%2C%22fundChannel%22%3A%22ALIPAYACCOUNT%22%7D%5D&
notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&
receipt_amount=10.00&app_id=2021002118614531&buyer_pay_amount=10.00&
sign_type=RSA2&seller_id=2088631877885139&gmt_payment=2021-01-09+19%3A34%3A20
&notify_time=2021-01-09+19%3A48%3A38&version=1.0&out_trade_no=3d923dcf-2bd1-48a1-b4e3-67a630acdc6b&total_amount=10.00
&trade_no=2021010922001446581403295652&auth_app_id=2021002118614531
&buyer_logon_id=guo***%40163.com&point_amount=0.00
 */

// @Title ZhifubaoNotify
// @Description 支付支付成功回调函数 ZhifubaoNotify
// @Success 200 status bool, data interface{}, msg string
// @router /zfb_notify [post]
func (this *NotifyController) ZhifubaoNotify() {
	trade_status := this.GetString("trade_status")
	trade_no := this.GetString("trade_no")
	out_trade_no := this.GetString("out_trade_no")
	total_amount := this.GetString("total_amount")
	fmt.Println(trade_status, trade_no, out_trade_no, total_amount)
	if trade_status != "TRADE_SUCCESS" {
		return
	}
	if strings.Contains(out_trade_no, "deposit-") { // 充值
		w_r, _, _ := models.GetWRByOrderNumber(out_trade_no)
		w_r.Status = 1
		err := w_r.Update()
		if err != nil {
			return
		}
		u_w, _ := models.GetWalletByUserId(w_r.UserId)
		t_amount, _ := strconv.ParseFloat(total_amount, 64)
		u_w.TotalAmount = u_w.TotalAmount + t_amount
		err = u_w.Update()
		if err != nil {
			return
		}
	} else { // 支付
		order, _, _ := models.GetGoodsOrderByOrderNumber(out_trade_no)
		order.OrderStatus = 2
		err := order.Update()
		if err != nil {
			return
		}
	}
	this.Data["json"] = RetResource(true, 200, nil, "success")
	this.ServeJSON()
	return
}