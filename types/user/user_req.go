package user

import (
	rds_conn "ganji/redis"
	"ganji/types"
	"github.com/pkg/errors"
	"regexp"
)

const (
	PhoneNumRule = "^(1[3|4|5|6|7|8|9][0-9]\\d{4,8})$"
	EmailPattern = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	IcardPattern = `( ^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)`
)

type PhoneNumberCheck struct {
	Phone string `json:"phone"`
}

func (this PhoneNumberCheck) PhoneNumberParamValidate() (int, error) {
	result, _ := regexp.MatchString(PhoneNumRule, this.Phone)
	if !result {
		return types.PhoneFormatError, errors.New("手机号码格式不正确")
	}
	return types.ReturnSuccess, nil
}


type PhoneCodeCheck struct {
	PhoneNumberCheck
	PhoneCode string `json:"phone_code"`
}

func (this PhoneCodeCheck) ReqPhoneCodeCheckParamValidate() (int, error) {
	code, err := this.PhoneNumberParamValidate()
	if err != nil {
		return code, err
	}
	if this.PhoneCode == "" {
		return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
	}
	phone_code := rds_conn.RdsConn.Get(this.Phone).Val()
	if phone_code != this.PhoneCode {
		return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
	}
	return types.ReturnSuccess, nil
}

type PhoneRigsterCheck struct {
	PhoneNumberCheck
	LoginRisger int8 `json:"login_risger"`  // 1: 登陆， 2：注册
}

func (this PhoneRigsterCheck) PhoneRigsterCheckParamValidate() (int, error) {
	code, err := this.PhoneNumberParamValidate()
	if err != nil {
		return code, err
	}
	if this.LoginRisger != 1 && this.LoginRisger != 2 {
		return types.NoThisLoginRisgerWay, errors.New("没有这种验证方式，请选择 1 或者 2; 1表示登陆，2：表示注册")
	}
	return types.ReturnSuccess, nil
}


type EmailNumberCheck struct {
	Email string `json:"email"`
}

func (this EmailNumberCheck) EmailNumberCheckParamValidate() (int, error) {
	result, _ := regexp.MatchString(EmailPattern, this.Email)
	if !result {
		return types.EmailFormatError, errors.New("邮箱格式不正确")
	}
	return types.ReturnSuccess, nil
}

type EmailCodeCheck struct {
	EmailNumberCheck
	EmailCode string `json:"email_code"`
}

func (this EmailCodeCheck) EmailCodeCheckParamValidate() (int, error) {
	code, err := this.EmailNumberCheckParamValidate()
	if err != nil {
		return code, err
	}
	if this.EmailCode == "" {
		return types.EmailVerifyCodeEmptyError, errors.New("邮箱验证码为空")
	}
	email_code := rds_conn.RdsConn.Get(this.Email).Val()
	if email_code != this.EmailCode {
		return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
	}
	return types.ReturnSuccess, nil
}


type EamilRigsterCheck struct {
	EmailNumberCheck
	LoginRisger int8 `json:"login_risger"`  // 1: 登陆， 2：注册
}


func (this EamilRigsterCheck) EamilRigsterCheckParamValidate() (int, error) {
	code, err := this.EmailNumberCheckParamValidate()
	if err != nil {
		return code, err
	}
	if this.LoginRisger != 1 && this.LoginRisger != 2 {
		return types.NoThisLoginRisgerWay, errors.New("没有这种验证方式，请选择 1 或者 2; 1:表示登陆，2:表示注册")
	}
	return types.ReturnSuccess, nil
}


type UserRegisterCheck struct {
	VerifyWay     int8    `json:"verify_way"`    // 1：手机号码验证； 2：邮箱验证
	PhoneEmail    string  `json:"phone_email"`
	PhonEmailCode string  `json:"phon_email_code"`
	Password1     string  `json:"password1"`
	Password2     string  `json:"password2"`
	InviteCode    string  `json:"invite_code"`
}


func (this UserRegisterCheck) UserRegisterCheckParamValidate() (int, error) {
	if this.VerifyWay == 1 { // 手机号码验证
		if this.PhonEmailCode == "" {
			return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
		}
		result, _ := regexp.MatchString(PhoneNumRule, this.PhoneEmail)
		if !result {
			return types.PhoneFormatError, errors.New("手机号码格式不正确")
		}
		phone_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if phone_code != this.PhonEmailCode {
			return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
		}
	} else if this.VerifyWay == 2 { // 邮箱验证
		result, _ := regexp.MatchString(EmailPattern, this.PhoneEmail)
		if !result {
			return types.EmailFormatError, errors.New("邮箱格式不正确")
		}
		email_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if email_code != this.PhonEmailCode {
			return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
		}
	} else {
		return types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
	if this.Password1 == "" || this.Password2 == "" {
		return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
	}
	if this.Password1 != this.Password2 {
		return types.TwicePasswordNotEqual, errors.New("两次输入的密码不一样")
	}
	return types.ReturnSuccess, nil
}


type UserLoginCheck struct {
	VerifyWay     int8    `json:"verify_way"`   // 1：手机号码验证； 2：邮箱验证
	PhoneEmail    string  `json:"phone_email"`
	PhonEmailCode string  `json:"phon_email_code"`
	Password      string  `json:"password"`
}


func (this UserLoginCheck) UserLoginCheckParamValidate() (int, error) {
	if this.VerifyWay == 1 {  // 手机号码验证
		if this.PhonEmailCode == "" {
			return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
		}
		result, _ := regexp.MatchString(PhoneNumRule, this.PhoneEmail)
		if !result {
			return types.PhoneFormatError, errors.New("手机号码格式不正确")
		}
		phone_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if phone_code != this.PhonEmailCode {
			return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
		}
		if this.Password == "" {
			return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
		}
	} else if this.VerifyWay == 2 { // 邮箱验证
		result, _ := regexp.MatchString(EmailPattern, this.PhoneEmail)
		if !result {
			return types.EmailFormatError, errors.New("邮箱格式不正确")
		}
		email_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if email_code != this.PhonEmailCode {
			return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
		}
		if this.Password == "" {
			return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
		}
	} else {
		return types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
	return types.ReturnSuccess, nil
}


type UpdatePasswordCheck struct {
	VerifyWay     int8    `json:"verify_way"`   // 1：手机号码验证； 2：邮箱验证
	PhoneEmail    string  `json:"phone_email"`
	PhonEmailCode string  `json:"phon_email_code"`
	OldPassword   string  `json:"old_password"`
	NewPassword   string  `json:"new_password"`
}


func (this UpdatePasswordCheck) UpdatePasswordCheckParamValidate() (int, error) {
	if this.VerifyWay == 1 { // 手机号码验证
		if this.PhoneEmail == "" {
			return types.ParamEmptyError, errors.New("手机号码为空")
		}
		result, _ := regexp.MatchString(PhoneNumRule, this.PhoneEmail)
		if !result {
			return types.PhoneFormatError, errors.New("手机号码格式不正确")
		}
		if this.PhonEmailCode == "" {
			return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
		}
		phone_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if phone_code != this.PhonEmailCode {
			return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
		}
	} else if this.VerifyWay == 2 { // 邮箱验证
		if this.PhoneEmail == "" {
			return types.ParamEmptyError, errors.New("邮箱为空")
		}
		result, _ := regexp.MatchString(EmailPattern, this.PhoneEmail)
		if !result {
			return types.EmailFormatError, errors.New("邮箱格式不正确")
		}
		email_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if email_code != this.PhonEmailCode {
			return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
		}
	} else {
		return types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
	if this.NewPassword == "" || this.OldPassword == "" {
		return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
	}
	if this.NewPassword == this.OldPassword {
		return types.NewOldPasswordEqual, errors.New("新密码和老密码相等")
	}
	return types.ReturnSuccess, nil
}


// 修改手机号码, 邮箱或者绑定手机号码邮箱
type UpdateCreatePhoneEmailCheck struct {
	UpdateWay     int8    `json:"update_way"`  // 1: 修改手机号码  2: 绑定手机号码;  3:修改邮箱; 4: 绑定邮箱
	PhoneEmail    string  `json:"phone_email"`
	PhonEmailCode string  `json:"phon_email_code"`
}

func (this UpdateCreatePhoneEmailCheck) UpdateCreatePhoneEmailCheckParamValidate() (int, error) {
	if this.UpdateWay == 1 || this.UpdateWay ==2 { // 修改手机号码
		if this.PhoneEmail == "" {
			return types.ParamEmptyError, errors.New("手机号码为空")
		}
		result, _ := regexp.MatchString(PhoneNumRule, this.PhoneEmail)
		if !result {
			return types.PhoneFormatError, errors.New("手机号码格式不正确")
		}
		if this.PhonEmailCode == "" {
			return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
		}
		phone_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if phone_code != this.PhonEmailCode {
			return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
		}
	} else if this.UpdateWay == 3 || this.UpdateWay == 4 { // 修改邮箱
		if this.PhoneEmail == "" {
			return types.ParamEmptyError, errors.New("邮箱为空")
		}
		result, _ := regexp.MatchString(EmailPattern, this.PhoneEmail)
		if !result {
			return types.EmailFormatError, errors.New("邮箱格式不正确")
		}
		email_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if email_code != this.PhonEmailCode {
			return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
		}
	} else {
		return types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
	return types.ReturnSuccess, nil
}

type ForgetPasswordCheck struct {
	VerifyWay     int8    `json:"verify_way"`   // 1：手机号码验证； 2：邮箱验证
	PhoneEmail    string  `json:"phone_email"`
	PhonEmailCode string  `json:"phon_email_code"`
	NewPassword1  string  `json:"new_password1"`
	NewPassword2  string  `json:"new_password2"`
}

func (this ForgetPasswordCheck) ForgetPasswordCheckParamValidate() (int, error) {
	if this.VerifyWay == 1 { // 手机号码验证
		if this.PhonEmailCode == "" {
			return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
		}
		result, _ := regexp.MatchString(PhoneNumRule, this.PhoneEmail)
		if !result {
			return types.PhoneFormatError, errors.New("手机号码格式不正确")
		}
		phone_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if phone_code != this.PhonEmailCode {
			return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
		}
	} else if this.VerifyWay == 2 { // 邮箱验证
		result, _ := regexp.MatchString(EmailPattern, this.PhoneEmail)
		if !result {
			return types.EmailFormatError, errors.New("邮箱格式不正确")
		}
		email_code := rds_conn.RdsConn.Get(this.PhoneEmail).Val()
		if email_code != this.PhonEmailCode {
			return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
		}
	} else {
		return types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
	if this.NewPassword1 == "" || this.NewPassword2 == "" {
		return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
	}
	if this.NewPassword1 != this.NewPassword1 {
		return types.TwicePasswordNotEqual, errors.New("两次输入的密码不一样")
	}
	return types.ReturnSuccess, nil
}


type UpdateUserInfoCheck struct {
	ImageId  int64  `json:"image_id"`
	UserName string `json:"user_name"`
}



// 用户实名制度认证
type UserAuthCheck struct {
	RealName       string `json:"real_name"`
	IdCard         string `json:"id_card"`
	IdCardPosImgId int64  `json:"id_card_pos_img_id"`
	IdCardNegImgId int64  `json:"id_card_neg_img_id"`
}

func (ua UserAuthCheck) UserAuthCheckParamValidate() (int, error) {
	if ua.RealName == "" {
		return types.RealNameEmpty, errors.New("real name is empty")
	}
	if ua.IdCard == "" {
		return types.IdCardEmpty, errors.New("id card is empty")
	}
	result, _ := regexp.MatchString(IcardPattern, ua.IdCard)
	if !result {
		return types.IdCardFormatError, errors.New("id card format is error")
	}
	return types.ReturnSuccess, nil
}
