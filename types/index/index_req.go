package index

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type IndexDownCheck struct {
	types.PageSizeData
	IndexCatId  int8 `json:"index_cat_id"` // 0:爆款产品; 1:活动优选；2:积分兑换；3:拼团送
}

func (this IndexDownCheck) IndexDownCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.IndexCatId < 0 || this.IndexCatId > 3 {
		return types.PhoneFormatError, errors.New("没有这种类型的商品")
	}
	return types.ReturnSuccess, nil
}
