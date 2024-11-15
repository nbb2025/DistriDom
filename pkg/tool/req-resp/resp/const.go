package resp

type RspCode int64

// 错误码
const (
	CODE_INVALID_PARAMETER       RspCode = 10000 + iota // 无效的参数
	CODE_EMAIL_EXISTS                                   // 邮箱重复
	CODE_INVALID_CAPTCHA                                // 验证码错误
	CODE_LOGIN_FAILED                                   // 登录失败
	CODE_GROUP_NOT_EXIST                                // 组不存在
	CODE_USER_IS_DOCTOR                                 // 目标用户为医生账号
	CODE_PARTNER_EXIST                                  // 已存在合作关系
	CODE_PARTNER_NOT_ESTABLISHED                        // 合作关系未建立
	CODE_ORDER_EXIST                                    // 订单已存在
	CODE_INVITATION_EXPIRED                             // 邀请过期
	CODE_FIND_FAILED                                    // 查询失败
	CODE_ATTACHMENT_EXISTS                              // 附件已存在
	CODE_ORDER_DOWNLOADED                               // 订单已被技工所接收
	CODE_INCORRECT_PASSWORD                             // 密码错误
	CODE_CAPTCHA_EXPIRED                                // 验证码过期
	CODE_SEND_FREQUENTLY                                // 发送过于频繁
	CODE_FILE_TOO_LARGE                                 // 文件过大
	CODE_INVALID_FILE_STRUCT                            // 文件类型错误
	CODE_FILE_INFO_EXIST                                // 文件信息已存在
	CODE_OPEN_FILE_FAILED                               // 打开文件错误

	CODE_SUCCESS        RspCode = 200   // 成功
	CODE_FAILED         RspCode = 400   // 失败
	CODE_NO_PERMISSIONS RspCode = 403   // 无权限
	CODE_ERR_BUSY       RspCode = 500   // 系统繁忙
	CODE_TOKEN_EXPIRED  RspCode = 20000 //token过期

)

// 翻译文件中的key
//
//	var codeToName = map[RspCode]string{
//		CODE_NO_PERMISSIONS:          "no permissions",
//		CODE_ERR_MSG:                 "error occurred",
//		CODE_ERR_BUSY:                "system is busy",
//		CODE_INVALID_PARAMETER:       "invalid parameter",
//		CODE_INVALID_CAPTCHA:         "invalid captcha",
//		CODE_ADD_FAILED:              "failed to add record",
//		CODE_DELETE_FAILED:           "failed to delete item",
//		CODE_UPDATE_FAILED:           "failed to update item",
//		CODE_FIND_FAILED:             "failed to find record",
//		CODE_LOGIN_FAILED:            "login failed",
//		CODE_VERIFY_SUCCESS:          "verify success",
//		CODE_OPERATION_FAILED:        "operation failed",
//		CODE_EMAIL_EXISTS:            "Email already exists",
//		CODE_SUCCESS:                 "Operation successful", // 成功
//		CODE_FAILED:                  "Operation failed",     // 失败
//		CODE_TOKEN_EXPIRED:           "Token expired",
//		CODE_GROUP_NOT_EXIST:         "Group does not exist",
//		CODE_USER_IS_DOCTOR:          "The target user is a doctor account",
//		CODE_PARTNER_EXIST:           "The partner relationship already exists",
//		CODE_PARTNER_NOT_ESTABLISHED: "Partnership not established",
//		CODE_ORDER_EXIST:             "Order already exists",
//		CODE_INVITATION_EXPIRED:      "Invitation has expired",
//		CODE_ATTACHMENT_EXISTS:       "Attachment already exists",
//		CODE_ORDER_DOWNLOADED:        "The order has been received by the mechanist",
//		CODE_INCORRECT_PASSWORD:      "Password incorrect",
//		CODE_SEND_FREQUENTLY:         "Sending emails too frequently",
//		CODE_CAPTCHA_EXPIRED:         "Captcha code expired",
//	}
var codeToName = map[RspCode]string{
	CODE_NO_PERMISSIONS:          "CODE_NO_PERMISSIONS",
	CODE_ERR_BUSY:                "CODE_ERR_BUSY",
	CODE_INVALID_PARAMETER:       "CODE_INVALID_PARAMETER",
	CODE_INVALID_CAPTCHA:         "CODE_INVALID_CAPTCHA",
	CODE_FIND_FAILED:             "CODE_FIND_FAILED",
	CODE_LOGIN_FAILED:            "CODE_LOGIN_FAILED",
	CODE_EMAIL_EXISTS:            "CODE_EMAIL_EXISTS",
	CODE_SUCCESS:                 "CODE_SUCCESS",
	CODE_FAILED:                  "CODE_FAILED",
	CODE_TOKEN_EXPIRED:           "CODE_TOKEN_EXPIRED",
	CODE_GROUP_NOT_EXIST:         "CODE_GROUP_NOT_EXIST",
	CODE_USER_IS_DOCTOR:          "CODE_USER_IS_DOCTOR",
	CODE_PARTNER_EXIST:           "CODE_PARTNER_EXIST",
	CODE_PARTNER_NOT_ESTABLISHED: "CODE_PARTNER_NOT_ESTABLISHED",
	CODE_ORDER_EXIST:             "CODE_ORDER_EXIST",
	CODE_INVITATION_EXPIRED:      "CODE_INVITATION_EXPIRED",
	CODE_ATTACHMENT_EXISTS:       "CODE_ATTACHMENT_EXISTS",
	CODE_ORDER_DOWNLOADED:        "CODE_ORDER_DOWNLOADED",
	CODE_INCORRECT_PASSWORD:      "CODE_INCORRECT_PASSWORD",
	CODE_SEND_FREQUENTLY:         "CODE_SEND_FREQUENTLY",
	CODE_CAPTCHA_EXPIRED:         "CODE_CAPTCHA_EXPIRED",
	CODE_FILE_INFO_EXIST:         "CODE_FILE_INFO_EXIST",
	CODE_INVALID_FILE_STRUCT:     "CODE_INVALID_FILE_STRUCT",
	CODE_OPEN_FILE_FAILED:        "CODE_OPEN_FILE_FAILED",
	CODE_FILE_TOO_LARGE:          "CODE_FILE_TOO_LARGE",
}

func (context RspCode) Msg() string {
	return codeToName[context]
}

func GetErr(s RspCode) Err {
	return Err{
		Code: s,
		Msg:  s.Msg(),
	}
}
