package shumei

// CommonResp 公共响应
type CommonResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

// TextDetectResp 文本检测响应
type TextDetectResp struct {
	CommonResp
	AllLabels []struct {
		Probability     interface{} `json:"probability"`
		RiskDescription string      `json:"riskDescription"`
		RiskDetail      struct {
		} `json:"riskDetail"`
		RiskLabel1 string `json:"riskLabel1"`
		RiskLabel2 string `json:"riskLabel2"`
		RiskLabel3 string `json:"riskLabel3"`
		RiskLevel  string `json:"riskLevel"`
	} `json:"allLabels"`
	AuxInfo struct {
		ContactResult []struct {
			ContactString string `json:"contactString"`
			ContactType   int    `json:"contactType"`
		} `json:"contactResult"`
		FilteredText string `json:"filteredText"`
	} `json:"auxInfo"`
	BusinessLabels  []interface{} `json:"businessLabels"`
	RiskDescription string        `json:"riskDescription"`
	RiskDetail      struct {
	} `json:"riskDetail"`
	RiskLabel1 string `json:"riskLabel1"`
	RiskLabel2 string `json:"riskLabel2"`
	RiskLabel3 string `json:"riskLabel3"`
	RiskLevel  string `json:"riskLevel"`
}

// ImageSyncDetectResp 图片单条同步检测响应
type ImageSyncDetectResp struct {
	CommonResp
	RiskLevel       string `json:"riskLevel"`
	RiskLabel1      string `json:"riskLabel1"`
	RiskLabel2      string `json:"riskLabel2"`
	RiskLabel3      string `json:"riskLabel3"`
	RiskDescription string `json:"riskDescription"`
	RiskDetail      struct {
		Faces []struct {
			FaceRatio   float64 `json:"face_ratio"`
			ID          string  `json:"id"`
			Location    []int   `json:"location"`
			Name        string  `json:"name"`
			Probability float64 `json:"probability"`
		} `json:"faces"`
		RiskSource int `json:"riskSource"`
	} `json:"riskDetail"`
	AuxInfo struct {
		Segments    int `json:"segments"`
		TypeVersion struct {
			POLITICS string `json:"POLITICS"`
			VIOLENCE string `json:"VIOLENCE"`
			BAN      string `json:"BAN"`
			PORN     string `json:"PORN"`
		} `json:"typeVersion"`
	} `json:"auxInfo"`
	AllLabels []struct {
		RiskLabel1      string  `json:"riskLabel1"`
		RiskLabel2      string  `json:"riskLabel2"`
		RiskLabel3      string  `json:"riskLabel3"`
		RiskLevel       string  `json:"riskLevel"`
		Probability     float64 `json:"probability,omitempty"`
		RiskDescription string  `json:"riskDescription,omitempty"`
		RiskDetail      struct {
			Faces []struct {
				FaceRatio   float64 `json:"face_ratio"`
				ID          string  `json:"id"`
				Location    []int   `json:"location"`
				Name        string  `json:"name"`
				Probability float64 `json:"probability"`
			} `json:"faces"`
			RiskSocrce int `json:"riskSocrce"`
		} `json:"riskDetail,omitempty"`
	} `json:"allLabels"`
	BusinessLabels []struct {
		BusinessDescription string `json:"businessDescription"`
		BusinessDetail      struct {
		} `json:"businessDetail"`
		BusinessLabel1  string  `json:"businessLabel1"`
		BusinessLabel2  string  `json:"businessLabel2"`
		BusinessLabel3  string  `json:"businessLabel3"`
		ConfidenceLevel int     `json:"confidenceLevel"`
		Probability     float64 `json:"probability"`
	} `json:"businessLabels"`
	TokenLabels struct {
		UGCAccountRisk struct {
		} `json:"UGC_account_risk"`
	} `json:"tokenLabels"`
}

// ImageAsyncDetectResp 图片单条异步检测响应
type ImageAsyncDetectResp struct {
	CommonResp
}

// ImagesSyncDetectResp 图片批量同步检测响应
type ImagesSyncDetectResp struct {
	CommonResp
	AuxInfo struct {
	} `json:"auxInfo"`
	Imgs []struct {
		AllLabels []interface{} `json:"allLabels"`
		AuxInfo   struct {
			Segments    int `json:"segments"`
			TypeVersion struct {
				OCR  string `json:"OCR"`
				PORN string `json:"PORN"`
			} `json:"typeVersion"`
		} `json:"auxInfo"`
		BtID            string `json:"btId"`
		Code            int    `json:"code"`
		Message         string `json:"message"`
		RequestID       string `json:"requestId"`
		RiskDescription string `json:"riskDescription"`
		RiskDetail      struct {
			RiskSource int `json:"riskSource"`
		} `json:"riskDetail"`
		RiskLabel1  string `json:"riskLabel1"`
		RiskLabel2  string `json:"riskLabel2"`
		RiskLabel3  string `json:"riskLabel3"`
		RiskLevel   string `json:"riskLevel"`
		TokenLabels struct {
			UGCAccountRisk struct {
			} `json:"UGC_account_risk"`
		} `json:"tokenLabels"`
	} `json:"imgs"`
}

// ImagesAsyncDetectResp 图片批量异步检测响应
type ImagesAsyncDetectResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	RequestIds []struct {
		RequesTID string `json:"reques-tId"`
		BtID      string `json:"btId"`
	} `json:"requestIds"`
}

// ImageQueryResp 图片查询响应
type ImageQueryResp struct {
	Code     int           `json:"code"`
	Message  string        `json:"message"`
	Contents []interface{} `json:"contents"`
}

// AudioSyncDetectResp 音频同步检测响应
type AudioSyncDetectResp struct {
	CommonResp
	BtID   string `json:"btId"`
	Detail struct {
		AudioDetail []struct {
			RequestID       string `json:"requestId"`
			AudioStarttime  int    `json:"audioStarttime"`
			AudioEndtime    int    `json:"audioEndtime"`
			AudioURL        string `json:"audioUrl"`
			RiskLevel       string `json:"riskLevel"`
			RiskLabel1      string `json:"riskLabel1"`
			RiskLabel2      string `json:"riskLabel2"`
			RiskLabel3      string `json:"riskLabel3"`
			RiskDescription string `json:"riskDescription"`
			RiskDetail      struct {
				AudioText string `json:"audioText"`
			} `json:"riskDetail"`
		} `json:"audioDetail"`
		AudioTags struct {
			Gender struct {
				Label       string `json:"label"`
				Probability int    `json:"probability"`
			} `json:"gender"`
			Language []struct {
				Confidence int `json:"confidence"`
				Label      int `json:"label"`
			} `json:"language"`
			Song   int `json:"song"`
			Timbre []struct {
				Label       string `json:"label"`
				Probability int    `json:"probability"`
			} `json:"timbre"`
		} `json:"audioTags"`
		AudioText     string `json:"audioText"`
		AudioTime     int    `json:"audioTime"`
		Code          int    `json:"code"`
		RequestParams struct {
			Channel       string `json:"channel"`
			Lang          string `json:"lang"`
			ReturnAllText int    `json:"returnAllText"`
			TokenID       string `json:"tokenId"`
		} `json:"requestParams"`
		RiskLevel string `json:"riskLevel"`
	} `json:"detail"`
}

// AudioAsyncDetectResp 音频异步检测响应
type AudioAsyncDetectResp struct {
	CommonResp
}

// AudioQueryResp 音频查询响应
type AudioQueryResp struct {
	CommonResp
	BtID          string        `json:"btId"`
	RiskLevel     string        `json:"riskLevel"`
	AudioText     string        `json:"audioText"`
	AudioTime     int           `json:"audioTime"`
	AudioDetail   []interface{} `json:"audioDetail"`
	AudioTags     interface{}   `json:"audioTags"`
	RequestParams interface{}   `json:"requestParams"`
	AuxInfo       interface{}   `json:"auxInfo"`
}
