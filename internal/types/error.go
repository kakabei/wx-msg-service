package types

type ErrorResp struct {
	Ret  HTTPCommonHead `json:"ret"`
	Body interface{}    `json:"body"`
}

type CodeErrorResponse struct {
	Ret  HTTPCommonHead `json:"ret"`
	Body interface{}    `json:"body"`
}

type HTTPCommonHead struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

const defaultCode = 35840001

var (
	HttpSucc                        HTTPCommonHead = HTTPCommonHead{Code: 0, Msg: "OK"}
	HttpCheckParamError             HTTPCommonHead = HTTPCommonHead{Code: 3584002, Msg: "参数错误"}
	HttpJosnMarshalErr              HTTPCommonHead = HTTPCommonHead{Code: 3584003, Msg: "json序列化失败"}
	HttpGetWxQrcode                 HTTPCommonHead = HTTPCommonHead{Code: 3584301, Msg: "创建二维码失败"}
	HttpSendPhoneCaptcha            HTTPCommonHead = HTTPCommonHead{Code: 3584302, Msg: "发送手机验证码失败"}
	HttpVerifyPhoneCaptcha          HTTPCommonHead = HTTPCommonHead{Code: 3584303, Msg: "验证码失效或错误"}
	HttpPhoneCaptchaERROR           HTTPCommonHead = HTTPCommonHead{Code: 3584304, Msg: "验证码错误"}
	HttpUpdatePhone                 HTTPCommonHead = HTTPCommonHead{Code: 3584305, Msg: "更新手机号失败"}
	HttpGetSaasBiz                  HTTPCommonHead = HTTPCommonHead{Code: 3584306, Msg: "查询租户失败"}
	HttpCreateNewUser               HTTPCommonHead = HTTPCommonHead{Code: 3584307, Msg: "创建新用户失败"}
	HttpIDCardCheck                 HTTPCommonHead = HTTPCommonHead{Code: 3584308, Msg: "调用检测实名认证失败"}
	HttpIDCardCheckNO               HTTPCommonHead = HTTPCommonHead{Code: 3584309, Msg: "名字和身份证号不匹配"}
	HttpIDCardCheckFail             HTTPCommonHead = HTTPCommonHead{Code: 3584310, Msg: "实名认证失败"}
	HttpWxQrcodeSetRedis            HTTPCommonHead = HTTPCommonHead{Code: 3584311, Msg: "设置缓存失败"}
	HttpUserInvalid                 HTTPCommonHead = HTTPCommonHead{Code: 3584312, Msg: "用户未生效"}
	HttpUserDelete                  HTTPCommonHead = HTTPCommonHead{Code: 3584313, Msg: "用户已删除"}
	HttpUserLoginErr                HTTPCommonHead = HTTPCommonHead{Code: 3584314, Msg: "用户没有登录"}
	HttpUserNotExist                HTTPCommonHead = HTTPCommonHead{Code: 3584315, Msg: "用户没有注册"}
	HttpZoreError                   HTTPCommonHead = HTTPCommonHead{Code: 3584316, Msg: "用户不拥有该区域资源"}
	HttpGetSaleInfoErr              HTTPCommonHead = HTTPCommonHead{Code: 3584317, Msg: "查询商品信息失败"}
	HttpNativePrePayErr             HTTPCommonHead = HTTPCommonHead{Code: 3584318, Msg: "请求商品二维码失败"}
	HttpRedisGet                    HTTPCommonHead = HTTPCommonHead{Code: 3584319, Msg: "查询缓存失败"}
	HttpRqcodeNotExpire             HTTPCommonHead = HTTPCommonHead{Code: 3584320, Msg: "二维码过期了，请刷新"}
	HttpUidEmpty                    HTTPCommonHead = HTTPCommonHead{Code: 3584321, Msg: "Uid不存在"}
	HttpWaitForScan                 HTTPCommonHead = HTTPCommonHead{Code: 3584322, Msg: "请扫描二维码"}
	HttpCoinNotEnough               HTTPCommonHead = HTTPCommonHead{Code: 3584323, Msg: "nika币不足,请充值"}
	HttpPhoneExist                  HTTPCommonHead = HTTPCommonHead{Code: 3584324, Msg: "该手机号已被注册"}
	HttpGetOrderList                HTTPCommonHead = HTTPCommonHead{Code: 3584325, Msg: "查询用户历史订单失败"}
	HttpPhoneNoLogin                HTTPCommonHead = HTTPCommonHead{Code: 3584326, Msg: "手机号没有注册"}
	HttpUserResStateApply           HTTPCommonHead = HTTPCommonHead{Code: 3584327, Msg: "资源正在申请中，请稍后再试"}
	HttpQueryResFail                HTTPCommonHead = HTTPCommonHead{Code: 3584501, Msg: "查询资源失败"}
	HttpNotifyResStatus             HTTPCommonHead = HTTPCommonHead{Code: 3584502, Msg: "申请资源已被释放"}
	HttpQueryResDirectIpv4Address   HTTPCommonHead = HTTPCommonHead{Code: 3584503, Msg: "查询资源直连地址失败"}
	HttpQueryResResourceIpv4Address HTTPCommonHead = HTTPCommonHead{Code: 3584504, Msg: "查询资接入机地址失败"}
	HttpAlreadyFirstPurchase        HTTPCommonHead = HTTPCommonHead{Code: 3584505, Msg: "你已经参加过首购，请购买其他商品"}
	HttpAddCostInfo                 HTTPCommonHead = HTTPCommonHead{Code: 3584506, Msg: "添加消费信息失败"}
	HttpLockError                   HTTPCommonHead = HTTPCommonHead{Code: 3584507, Msg: "加锁失败"}
	HttpDalGetError                 HTTPCommonHead = HTTPCommonHead{Code: 3584902, Msg: "查询数据错误"}
	HttpDalUpdateError              HTTPCommonHead = HTTPCommonHead{Code: 3584902, Msg: "更新数据错误"}
	HttpDalInsertError              HTTPCommonHead = HTTPCommonHead{Code: 3584903, Msg: "更新数据错误"}
)

/////

func NewResultError(requestId string, ret HTTPCommonHead) error {
	ret.RequestID = requestId
	return &ret
}

func NewCodeError(requestId string, code int, msg string) error {
	return &HTTPCommonHead{RequestID: requestId, Code: code, Msg: msg}
}

func NewDefaulResultError(ret HTTPCommonHead) error {
	return &ret
}

func NewDefaultError(requestId string, msg string) error {
	return NewCodeError(requestId, defaultCode, msg)
}

func (e *HTTPCommonHead) Error() string {
	return e.Msg
}

func (e *HTTPCommonHead) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Ret: *e,
	}
}
