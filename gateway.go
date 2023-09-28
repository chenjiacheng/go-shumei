package shumei

const (
	RegionBeijing       = "bj"   // 北京地区
	RegionShanghai      = "sh"   // 上海地区
	RegionGuangzhou     = "gz"   // 广州地区
	RegionVirginia      = "fjny" // 美国（弗吉尼亚）地区
	RegionSingapore     = "xjp"  // 新加坡地区
	RegionSiliconValley = "gg"   // 硅谷地区
	RegionIndia         = "yd"   // 印度地区
)

const (
	ServiceText        = "text"
	ServiceImageSync   = "image_sync"
	ServiceImageAsync  = "image_async"
	ServiceImagesSync  = "images_sync"
	ServiceImagesAsync = "images_async"
	ServiceImageQuery  = "image_query"
	ServiceAudioSync   = "audio_sync"
	ServiceAudioAsync  = "audio_async"
	ServiceAudioQuery  = "audio_query"
)

var GatewayMap = map[string]map[string]string{
	ServiceText: {
		RegionBeijing:   "http://api-text-bj.fengkongcloud.com/text/v4",   // 北京
		RegionShanghai:  "http://api-text-sh.fengkongcloud.com/text/v4",   // 上海
		RegionGuangzhou: "http://api-text-gz.fengkongcloud.com/text/v4",   // 广州
		RegionVirginia:  "http://api-text-fjny.fengkongcloud.com/text/v4", // 美国（弗吉尼亚）
		RegionSingapore: "http://api-text-xjp.fengkongcloud.com/text/v4",  // 新加坡
	},
	ServiceImageSync: {
		RegionBeijing:       "http://api-img-bj.fengkongcloud.com/image/v4",  // 北京
		RegionShanghai:      "http://api-img-sh.fengkongcloud.com/image/v4",  // 上海
		RegionSingapore:     "http://api-img-xjp.fengkongcloud.com/image/v4", // 新加坡
		RegionIndia:         "http://api-img-yd.fengkongcloud.com/image/v4",  // 印度
		RegionSiliconValley: "http://api-img-gg.fengkongcloud.com/image/v4",  // 硅谷
	},
	ServiceImageAsync: {
		RegionBeijing:  "http://api-img-bj.fengkongcloud.com/v4/saas/async/img", // 北京
		RegionShanghai: "http://api-img-sh.fengkongcloud.com/v4/saas/async/img", // 上海
	},
	ServiceImagesSync: {
		RegionBeijing:       "http://api-img-bj.fengkongcloud.com/images/v4",  // 北京
		RegionShanghai:      "http://api-img-sh.fengkongcloud.com/images/v4",  // 上海
		RegionSingapore:     "http://api-img-xjp.fengkongcloud.com/images/v4", // 新加坡
		RegionIndia:         "http://api-img-yd.fengkongcloud.com/images/v4",  // 印度
		RegionSiliconValley: "http://api-img-gg.fengkongcloud.com/images/v4",  // 硅谷
	},
	ServiceImagesAsync: {
		RegionBeijing:  "http://api-img-bj.fengkongcloud.com/v4/saas/async/imgs", // 北京
		RegionShanghai: "http://api-img-sh.fengkongcloud.com/v4/saas/async/imgs", // 上海
	},
	ServiceImageQuery: {
		RegionBeijing: "http://api-img-active-query.fengkongcloud.com/v4/image/query", // 北京
	},
	ServiceAudioSync: {
		RegionShanghai: "http://api-audio-sh.fengkongcloud.com/audiomessage/v4", // 上海
	},
	ServiceAudioAsync: {
		RegionShanghai:      "http://api-audio-sh.fengkongcloud.com/audio/v4",  // 上海
		RegionSiliconValley: "http://api-audio-gg.fengkongcloud.com/audio/v4",  // 硅谷
		RegionSingapore:     "http://api-audio-xjp.fengkongcloud.com/audio/v4", // 新加坡
	},
	ServiceAudioQuery: {
		RegionShanghai:      "http://api-audio-sh.fengkongcloud.com/query_audio/v4",  // 上海
		RegionSiliconValley: "http://api-audio-gg.fengkongcloud.com/query_audio/v4",  // 硅谷
		RegionSingapore:     "http://api-audio-xjp.fengkongcloud.com/query_audio/v4", // 新加坡
	},
}
