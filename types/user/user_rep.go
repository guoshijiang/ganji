package user

import "time"

// 用户登陆成功数据返回形式
type UserLoginRet struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}


type UserInfoRet struct {
	UserId       int64   `json:"user_id"`
	Token        string  `json:"token"`
	UserName     string  `json:"user_name"`
	IgAmount     float64 `json:"integral"`
	CnyAmount    float64 `json:"cny_amount"`
	Phone        string `json:"phone"`
	Eamil        string `json:"eamil"`
	Sex          int8   `json:"sex"`
	MemberLevel  int8   `json:"member_level"`
	InviteCode   string `json:"invite_code"`
	Avator       string `json:"avator"`
	RealName     string `json:"real_name"`
	WeiChat      string `json:"wei_chat"`
	QQ           string `json:"qq"`
}

type UserConponRet struct {
	ConponId    int64     `json:"id"`
	ConponName  string    `json:"conpon_name"`
	IsUsed      int8      `json:"is_used"`
	TotalAmount float64   `json:"total_amount"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}