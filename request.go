package shumei

// CommonReq 公共请求
type CommonReq struct {
	AccessKey string `json:"accessKey"`
	AppID     string `json:"appId"`
	EventID   string `json:"eventId"`
}

// TextDetectReq 文本检测请求
type TextDetectReq struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// ImageSyncDetectReq 图片单条同步检测请求
type ImageSyncDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	Data         map[string]interface{} `json:"data"`
}

// ImageAsyncDetectReq 图片单条异步检测请求
type ImageAsyncDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	Data         map[string]interface{} `json:"data"`
	Callback     string                 `json:"callback"`
}

// ImagesSyncDetectReq 图片批量同步检测请求
type ImagesSyncDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	Data         map[string]interface{} `json:"data"`
}

// ImagesAsyncDetectReq 图片批量异步检测请求
type ImagesAsyncDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	Data         map[string]interface{} `json:"data"`
	Callback     string                 `json:"callback"`
}

// ImageQueryReq 图片查询请求
type ImageQueryReq struct {
	RequestIds []struct {
		RequestID string `json:"requestId"`
		BtID      string `json:"btId"`
	} `json:"requestIds"`
}

// AudioSyncDetectReq 音频同步检测请求
type AudioSyncDetectReq struct {
	Type        string                 `json:"type"`
	ContentType string                 `json:"contentType"`
	Content     string                 `json:"content"`
	Data        map[string]interface{} `json:"data"`
	BtID        string                 `json:"btId"`
}

// AudioAsyncDetectReq 音频异步检测请求
type AudioAsyncDetectReq struct {
	Type         string                 `json:"type"`
	BusinessType string                 `json:"businessType"`
	ContentType  string                 `json:"contentType"`
	Content      string                 `json:"content"`
	Data         map[string]interface{} `json:"data"`
	BtID         string                 `json:"btId"`
	Callback     string                 `json:"callback"`
}

// AudioQueryReq 音频查询请求
type AudioQueryReq struct {
	BtID string `json:"btId"`
}
