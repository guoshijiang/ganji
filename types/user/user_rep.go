package user

// 用户登陆成功数据返回形式
type UserLoginRet struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}