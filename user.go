package wechat

import (
	"fmt"

	"github.com/esap/wechat/util"
)

// WXAPI 企业号用户列表接口
const (
	WXAPI_GETUSER     = WXAPI_ENT + "user/getuserinfo?access_token=%s&code=%s"
	WXAPI_GETUSERINFO = WXAPI_ENT + "user/get?access_token=%s&userid=%s"
	WXAPI_USERLIST    = WXAPI_ENT + `user/list?access_token=%s&department_id=1&fetch_child=1&status=0`
	WXAPI_USERADD     = WXAPI_ENT + `user/create?access_token=`
	WXAPI_DEPT        = WXAPI_ENT + `department/list?access_token=%s&id=1`
)

// UserOauth 用户鉴权信息
type UserOauth struct {
	WxErr
	UserId   string
	DeviceId string
	OpenId   string
}

// GetUserOauth 通过code鉴权
func GetUserOauth(code string) (userOauth UserOauth, err error) {
	url := fmt.Sprintf(WXAPI_GETUSER, GetAccessToken(), code)
	if err = util.GetJson(url, &userOauth); err != nil {
		return
	}
	if userOauth.ErrCode != 0 {
		err = fmt.Errorf("GetUserId error : errcode=%v , errmsg=%v", userOauth.ErrCode, userOauth.ErrMsg)
	}
	return
}

// UserInfo 用户信息
type UserInfo struct {
	WxErr
	UserId     string
	Name       string
	Department []int
	Position   string
	Mobile     string
	Gender     string
	Email      string
	WeixinId   string
	Avatar     string
	Status     int
	ExtAttr    struct {
		Attrs []struct {
			Name  string
			Value string
		}
	}
}

// GetUserInfo 通过userId获取用户信息
func GetUserInfo(userId string) (userInfo UserInfo, err error) {
	url := fmt.Sprintf(WXAPI_GETUSERINFO, GetAccessToken(), userId)
	if err = util.GetJson(url, &userInfo); err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserId error : errcode=%v , errmsg=%v", userInfo.ErrCode, userInfo.ErrMsg)
	}
	return
}

// UserList 用户列表
type UserList struct {
	WxErr
	UserList []UserInfo
}

// GetUserList 获取用户列表
func GetUserList() (userList UserList, err error) {
	url := fmt.Sprintf(WXAPI_USERLIST, GetAccessToken())
	if err = util.GetJson(url, &userList); err != nil {
		return
	}
	if userList.ErrCode != 0 {
		err = fmt.Errorf("MediaUpload error : errcode=%v , errmsg=%v", userList.ErrCode, userList.ErrMsg)
	}
	return
}

type DeptList struct {
	WxErr
	Department []struct {
		Id       int
		Name     string
		ParentId int
		Order    int
	}
}

// GetDeptList 获取部门列表
func GetDeptList() (deptList DeptList, err error) {
	url := fmt.Sprintf(WXAPI_DEPT, GetAccessToken())
	if err = util.GetJson(url, &deptList); err != nil {
		return
	}
	if deptList.ErrCode != 0 {
		err = fmt.Errorf("MediaUpload error : errcode=%v , errmsg=%v", deptList.ErrCode, deptList.ErrMsg)
	}
	return
}
