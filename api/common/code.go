package common

const (
	RetCodeQRCodeExpired           = -106  // 二维码过期
	RetCodeSendPhoneCodeFrequently = -3101 // 发送验证码频率过快
	RetCodeGameHasSigned           = -5003 // 已签到
	RetCodeForumHasSigned          = 1008  // 打卡失败或重复打卡
	RetCodeForumNeedVerification   = 1034  // 需要验证码
)
