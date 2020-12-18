package models

import (
	"encoding/base64"
	"ganji/common"
	"ganji/types"
	type_user "ganji/types/user"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	BaseModel
	Id             int64         `json:"id"`
	Phone          string        `orm:"size(64);index" json:"phone"`
	UserName       string        `orm:"size(128)" json:"user_name"`
	Avator         string        `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"`
	Password       string        `orm:"size(128)" json:"password"`
	FundPassword   string        `orm:"size(128)" json:"fund_password"`           // 钱包资金密码
	Email          string        `orm:"size(128);index" json:"email"`
	LoginCount     int64         `orm:"default(0);index" json:"login_count"`
	Token          string        `orm:"size(128)" json:"token"`
	IsAuth         int8          `orm:"default(0);index" json:"is_auth"`          // 0 未实名认证，1: 实名认证中；2:实名认证成功；3实名认证失败
	MemberLevel    int8          `orm:"default(1);index" json:"member_level"`     // 0:v0:普通会员 1:V1:白银会员，2:V2:白金会员，3:V3:黄金会员; 4:V4:砖石会有; 5:V5:皇冠会员
	MyInviteCode   string        `orm:"size(128)" json:"my_invite_code"`          // 用户自己网体邀请码
	InviteMeUserId int64         `orm:"size(64);index" json:"invite_me_user_id"`  // 网体上级用户id
}

func (this *User) TableName() string {
	return common.TableName("user")
}

func (this *User) SearchField() []string {
	return []string{"user_name"}
}

func (this *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *User) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *User) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *User) ExistByPhone(phone string) bool {
	return orm.NewOrm().QueryTable(this).Filter("phone", phone).Exist()
}

func (this *User) ExistByEmail(email string) bool {
	return orm.NewOrm().QueryTable(this).Filter("email", email).Exist()
}

func (this *User) GetInviteMeUser(inviteCode string) (*User, error) {
	var user User
	err := user.Query().Filter("my_invite_code", inviteCode).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (u *User) GetUserByPhone(phone string) (User, error) {
	var query_user User
	err := query_user.Query().Filter("phone", phone).Limit(1).One(&query_user)
	if err != nil {
		return query_user, errors.New("user is not exist")
	}
	return query_user, nil
}

func (u *User) GetUserByEmail(email string) (User, error) {
	var query_user User
	err := query_user.Query().Filter("email", email).Limit(1).One(&query_user)
	if err != nil {
		return query_user, errors.New("user is not exist")
	}
	return query_user, nil
}

func (u *User) AddLoginCount() error {
	u.LoginCount += 1
	u.UpdatedAt = time.Now()
	return u.Update("LoginCount", "UpdatedAt")
}

func GetUserByToken(token string) (*User, error) {
	u := User{}
	if err := orm.NewOrm().QueryTable(u.TableName()).RelatedSel().Filter("token", token).One(&u); err != nil {
		return nil, errors.Wrap(err, "error in GetUserByToken")
	}
	return &u, nil
}

// 通过手机号码或者邮箱注册
func RegisterByPhoneOrEmail(register_parm type_user.UserRegisterCheck) (success bool, code int, err error) {
	u := User{}
	var phone, email string
	if register_parm.VerifyWay == 1 {       // 1：手机号码验证
		if u.ExistByPhone(register_parm.PhoneEmail) {
			return false, types.UserIsExist, errors.New("用户已经注册")
		}
		phone = register_parm.PhoneEmail
	} else if register_parm.VerifyWay == 2 {  // 2：邮箱验证
		if u.ExistByEmail(register_parm.PhoneEmail) {
			return false, types.UserIsExist, errors.New("用户已经注册")
		}
		email = register_parm.PhoneEmail
	} else {
		return false, types.InvalidVerifyWay, errors.New("无效的校验方式")
	}
	var inviteMeUserID int64
	if len(register_parm.InviteCode) > 0 {
		inviteMeUser, err := u.GetInviteMeUser(register_parm.InviteCode)
		if err != nil {
			return false, types.InviteCodeNotExist, errors.New("对不起，没有这个邀请码，请核对后输入")
		}
		inviteMeUserID = inviteMeUser.Id
	} else {
		inviteMeUserID = 0
	}

	token,inviteCode := uidCode()
	//uid,_:= uuid.NewV4()
	//uuid,_:= uuid.NewV4()
	//hex_uuid := base64.RawURLEncoding.EncodeToString(uid.Bytes())
	user_reg_data := User {
		Phone:          phone,
		UserName:      "小鱼儿",
		Password:       common.ShaOne(register_parm.Password1),
		Email:          email,
		Token:          token,
		InviteMeUserId: inviteMeUserID,
		MyInviteCode:   inviteCode,
	}
	err, user_id := user_reg_data.Insert()
	if err != nil {
		return false, types.CreateUserFail, errors.New("创建用户失败")
	}
	user_info_query := UserInfo {
		UserId: user_id,
	}
	if err := user_info_query.Insert(); err != nil {
		return false, types.CreateUserFail, errors.New("创建用户信息失败")
	}
	user_wallet := UserWallet{
		UserId:    user_id,
		AssetName: "人民币",
		TotalAmount: 0,
	}
	if err := user_wallet.Insert(); err != nil {
		return false, types.CreateUserWalletFail, errors.New("创建用户钱包失败")
	}
	crfr_integral := UserIntegral{
		UserId: user_id,
		IntegralName: "商城积分",
		TotalIg: 0,
		UsedIg:  0,
		TodayIg: 0,
	}
	if err := crfr_integral.Insert(); err != nil {
		return false, types.InsertIntegralFail, errors.New("插入积分失败")
	}
	return true, types.ReturnSuccess, nil
}

func uidCode() (string,string){
	uid := uuid.NewV4()
	tuid := uuid.NewV4()
	hex_uuid := base64.RawURLEncoding.EncodeToString(uid.Bytes())
	return tuid.String(),hex_uuid
}

//管理后台添加用户
func (Self *User) RegisterUserByAdmin() error {
	token,inviteCode := uidCode()
	user_reg_data := User {
		Phone:          Self.Phone,
		UserName:      Self.UserName,
		Password:       common.ShaOne(Self.Password),
		Email:          Self.Email,
		Token:          token,
		InviteMeUserId: 0,
		MyInviteCode:   inviteCode,
		Avator: Self.Avator,
	}

	err, user_id := user_reg_data.Insert()
	if err != nil {
		return  errors.New("创建用户失败")
	}
	user_info_query := UserInfo {
		UserId: user_id,
	}
	if err := user_info_query.Insert(); err != nil {
		return errors.New("创建用户信息失败")
	}
	user_wallet := UserWallet{
		UserId:    user_id,
		AssetName: "人民币",
		TotalAmount: 0,
	}
	if err := user_wallet.Insert(); err != nil {
		return errors.New("创建用户钱包失败")
	}
	crfr_integral := UserIntegral{
		UserId: user_id,
		IntegralName: "商城积分",
		TotalIg: 0,
		UsedIg:  0,
		TodayIg: 0,
	}
	if err := crfr_integral.Insert(); err != nil {
		return errors.New("插入积分失败")
	}
	return nil
}


// 通过手机号码,邮箱登陆
func LoginByPhoneOrEmail(login_parm type_user.UserLoginCheck) (bool, *type_user.UserLoginRet, int, error) {
	u := User{}
	if login_parm.VerifyWay == 1 { // 1：手机号码验证
		ret_user, err := u.GetUserByPhone(login_parm.PhoneEmail)
		if err != nil {
			return false, nil, types.UserNotRegister, errors.New("用户没有注册")
		}
		if ret_user.Password != common.ShaOne(login_parm.Password) {
			return false, nil, types.PasswordError, errors.New("输入密码错误")
		}
		if err := u.AddLoginCount(); err != nil {
			return false, nil, types.AddLoginTimesError, errors.New("添加登陆次数错误")
		}
		return true, &type_user.UserLoginRet{
			Id:       ret_user.Id,
			UserName: ret_user.UserName,
			Token:    ret_user.Token,
			Phone:    ret_user.Phone,
		}, types.ReturnSuccess, nil
	} else if login_parm.VerifyWay == 2 {  // 2：邮箱验证
		ret_user, err := u.GetUserByEmail(login_parm.PhoneEmail)
		if err != nil {
			return false, nil, types.UserNotRegister, errors.New("用户没有注册")
		}
		if ret_user.Password != common.ShaOne(login_parm.Password) {
			return false, nil, types.PasswordError,  errors.New("输入密码错误")
		}
		if err := u.AddLoginCount(); err != nil {
			return false, nil, types.AddLoginTimesError, errors.New("添加登陆次数错误")
		}
		return true, &type_user.UserLoginRet{
			Id:       ret_user.Id,
			Token:    ret_user.Token,
			Phone:    ret_user.Phone,
		}, types.ReturnSuccess, nil
	} else {
		return false, nil, types.InvalidVerifyWay, errors.New("没有这种验证方式")
	}
}


// 修改密码
func UpdatePassword(u_params type_user.UpdatePasswordCheck, user_id int64) (success bool, code int, err error) {
	u := User{}
	if err := orm.NewOrm().QueryTable(u.TableName()).RelatedSel().Filter("id", user_id).One(&u); err != nil {
		return false, types.UserIsNotExist, errors.New("用户不存在")
	}
	if u.Password == common.ShaOne(u_params.NewPassword) {
		return false, types.NewOldPasswordEqual, errors.New("新老密码一样")
	}
	u.Password = common.ShaOne(u_params.NewPassword)
	err = u.Update()
	if err != nil {
		return false, types.SystemDbErr, errors.New("修改密码数据库操作失败")
	}
	return true, types.ReturnSuccess, nil
}


// 修改手机号码，邮箱或者 绑定手机号码邮箱
func UpdateOrCrearePhoneEmail(u_params type_user.UpdateCreatePhoneEmailCheck, user_id int64) (success bool, code int, err error) {
	u := User{}
	if err := orm.NewOrm().QueryTable(u.TableName()).RelatedSel().Filter("id", user_id).One(&u); err != nil {
		return false, types.UserIsNotExist, errors.New("用户不存在")
	}
	if u_params.UpdateWay == 1 { // 修改手机号码
		if u.Phone == u_params.PhoneEmail {
			return false, types.OldNewPhoneIsEqual, errors.New("新旧手机号码一样")
		} else {
			u.Phone = u_params.PhoneEmail
			err := u.Update()
			if err != nil {
				return false, types.SystemDbErr, errors.New("数据库操作失败")
			}
			return true, types.ReturnSuccess, nil
		}
	} else if u_params.UpdateWay == 2 { // 绑定手机号码
		if u.Phone != "" {
			return false, types.PhoneIsAlreadyBind, errors.New("手机号码已经绑定")
		} else {
			u.Phone = u_params.PhoneEmail
			err := u.Update()
			if err != nil {
				return false, types.SystemDbErr, errors.New("数据库操作失败")
			}
			return true, types.ReturnSuccess, nil
		}
	} else if u_params.UpdateWay == 3 { // 修改邮箱
		if u.Email == u_params.PhoneEmail {
			return false, types.OldNewEmailEqual, errors.New("新旧邮箱一样")
		} else {
			u.Email = u_params.PhoneEmail
			err := u.Update()
			if err != nil {
				return false, types.SystemDbErr, errors.New("数据库操作失败")
			}
			return true, types.ReturnSuccess, nil
		}
	} else if u_params.UpdateWay == 4 { // 绑定邮箱
		if u.Email != "" {
			return false, types.EamilAlreadyBind, errors.New("邮箱已经绑定")
		} else {
			u.Email = u_params.PhoneEmail
			err := u.Update()
			if err != nil {
				return false, types.SystemDbErr, errors.New("数据库操作失败")
			}
			return true, types.ReturnSuccess, nil
		}
	} else {
		return false, types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
}


// 找回登陆密码
func ForgetPassword(u_params type_user.ForgetPasswordCheck) (success bool, code int, err error) {
	u := User{}
	if u_params.VerifyWay == 0 { // 手机号码
		user_one, err := u.GetUserByPhone(u_params.PhoneEmail)
		if err != nil {
			return false, types.BindPhoneError, errors.New("没有绑定手机号码")
		}
		if user_one.Password == common.ShaOne(u_params.NewPassword1) {
			return false, types.NewOldPasswordEqual, errors.New("新密码和原密码一样")
		}
		user_one.Password = common.ShaOne(u_params.NewPassword1)
		err = user_one.Update()
		if err != nil {
			return false, types.SystemDbErr, errors.New("新密码和原密码一样")
		}
		return true, types.ReturnSuccess, nil
	} else if u_params.VerifyWay == 1 {
		user_two, err := u.GetUserByEmail(u_params.PhoneEmail)
		if err != nil {
			return false, types.BindEamilError, errors.New("没有绑定邮箱")
		}
		if user_two.Password == common.ShaOne(u_params.NewPassword1) {
			return false, types.NewOldPasswordEqual, errors.New("新密码和原密码一样")
		}
		user_two.Password = common.ShaOne(u_params.NewPassword1)
		err = user_two.Update()
		if err != nil {
			return false, types.SystemDbErr, errors.New("数据库操作失败")
		}
		return true, types.ReturnSuccess, nil
	} else {
		return false, types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
}

// 修改用户信息
func UpdateUserInfo(id int64, user_info type_user.UpdateUserInfoCheck) (success bool, code int, err error) {
	var user_data User
	if err := orm.NewOrm().QueryTable(User{}).RelatedSel().Filter("id", id).One(&user_data); err != nil {
		return false, types.UserIsNotExist, errors.New("用户不存在")
	}
	if user_info.UserName != "" {
		user_data.UserName = user_info.UserName
		err := user_data.Update()
		if err != nil {
			return false, types.UpdateUserInfoFail, errors.New("更新用户信息失败")
		}
	}
	if user_info.ImageId > 0 {
		var imgfile ImageFile
		img, code, err := imgfile.GetImageById(user_info.ImageId)
		if err != nil {
			return false, code, errors.New("文件不存在")
		}
		user_data.Avator = img.Url
		err = user_data.Update()
		if err != nil {
			return false, types.UpdateUserInfoFail, errors.New("更新用户信息失败")
		}
	}
	return true, types.ReturnSuccess, nil
}