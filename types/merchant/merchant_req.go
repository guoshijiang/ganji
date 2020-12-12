package merchant

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type MerchantListCheck struct {
	types.PageSizeData
	MerchantName    string `json:"merchant_name"`
	MerchantAddress string `json:"merchant_address"`
}

func (this MerchantListCheck) MerchantListCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	return types.ReturnSuccess, nil
}


type MerchantDetailCheck struct {
	MerchantId int64 `json:"merchant_id"`
}

func (this MerchantDetailCheck) MerchantDetailCheckParamValidate() (int, error) {
	if this.MerchantId <= 0 {
		return types.ParamLessZero, errors.New("MerchantId 不能小于 0")
	}
	return types.ReturnSuccess, nil
}