package ecodes

const ILLEGAL_OPERATION = 1
const PARAMS_ERROR = 2

// 作品
const TITLE_NOT_NULL = 1001
const PRICE_GREATER_0 = 1002
const COVER_NOT_NULL = 1003
const CATEGORY_NOT_NULL = 1004
const WORKS_NOT_EXISTS = 1005
const CASTING_CAST_FAIL = 1006
const SELL_FAIL = 1007
const NOT_REPEAT_SELL = 1008
const NO_AUTH_OPERATE = 1009
const CASTING_NOT_EXISTS = 1010

const UNFOLLOW_NOT_UNFOLLOW = 2001

const USER_NOT_EXISTS = 3001

const UNLIKE_NOT_UNLIKE = 4001
const TEST1 = 4002
const WORKS_SELL_OUT = 4003

const COLLECTION_NOT_EXISTS = 5001

var statusText = map[int]string{

	ILLEGAL_OPERATION: "非法操作",
	PARAMS_ERROR:      "参数错误",

	TITLE_NOT_NULL:     "标题不能为空",
	PRICE_GREATER_0:    "价格必须大于0",
	COVER_NOT_NULL:     "封面不能为空",
	CATEGORY_NOT_NULL:  "类目不能为空",
	WORKS_NOT_EXISTS:   "作品不存在",
	CASTING_CAST_FAIL:  "铸造失败",
	SELL_FAIL:          "挂售失败",
	NOT_REPEAT_SELL:    "请勿重复售卖",
	NO_AUTH_OPERATE:    "无权操作",
	CASTING_NOT_EXISTS: "铸件不存在",
	TEST1:              "签名错误",

	UNFOLLOW_NOT_UNFOLLOW: "尚未关注不能取消关注",
	UNLIKE_NOT_UNLIKE:     "没有喜欢不能取消",

	USER_NOT_EXISTS:       "user not exists",
	WORKS_SELL_OUT:        "works is sell out",
	COLLECTION_NOT_EXISTS: "collection not exists",
}

func ErrorMsg(code int) string {
	return statusText[code]
}
