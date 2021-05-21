package collector

//各种服务监控端口的json数据 映射 到 struct

//aicp_tts 监控端口数据映射
type AicpTts struct {
	TtsHttp TTSHTTP `json:"http"`
	Tts  []TTS `json:"tts"`
	WebSockets Websocket `json:"websocket"`
}

type TTSHTTP struct {
	Active float64 `json:"active"`
	Max float64 `json:"max"`
	Request float64 `json:"request"`
	ResponseTime float64 `json:"response_time"`
	Success float64 `json:"success"`
}

type TTS struct {
	Cap string `json:"cap"`
	Chars float64 `json:"chars"`
	Cost float64 `json:"const"`
	Current float64 `json:"current"`
	Prop string `json:"prop"`
	Total float64 `json:"total"`
}

type Websocket struct {
	Active float64 `json:"active"`
	Max float64 `json:"max"`
	Request float64 `json:"request"`
	ResponseTime float64 `json:"response_time"`
	Success float64 `json:"success"`
}


//aicp_license 监控端口数据映射
type AicpLicense struct {
	Licenses []Listen `json:"license"`
}

type Listen struct {
	Cap string `json:"cap"`
	Current float64 `json:"current"`
	//Props string `json:"props"`
	//Total float64 `json:"total"`
}

// aicp_mt 监控端口数据映射
type AicpMt struct {
	HTTP HTTP `json:"http"`
	Mt []Mt `json:"mt"`
}
type HTTP struct {
	Active float64 `json:"active"`
	Max float64 `json:"max"`
	Request float64 `json:"request"`
	ResponseTime float64 `json:"response_time"`
	Success float64 `json:"success"`
}
type Mt struct {
	Cap string `json:"cap"`
	Chars float64 `json:"chars"`
	Cost float64 `json:"cost"`
	Current float64 `json:"current"`
	Prop string `json:"prop"`
	Total float64 `json:"total"`
}

// AicpTransHttp  监控端口数据映射
type AicpTransHttp struct {
	Download Download `json:"download"`
	TransInterfaces []TransInterface `json:"interface"`
}
type Download struct {
	Downloading float64 `json:"downloading"`
	Failed float64 `json:"failed"`
	Success float64 `json:"success"`
}
type TransInterface struct {
	CancelFiles float64 `json:"cancel_files"`
	CancelTasks float64 `json:"cancel_tasks"`
	Property string `json:"property"`
	SubmitFiles float64 `json:"submit_files"`
	SubmitTasks float64 `json:"submit_tasks"`
}


// apigw  监控端口数据映射
type AicpApiGw struct {
	API []API `json:"api"`
}
type API struct {
	Active float64 `json:"active"`
	Name string `json:"name"`
	Requests float64 `json:"requests"`
	Success float64 `json:"success"`
}

//asr_ft 监控端口数据映射
type AicpAsrFT struct {
	Interfaces Interface `json:"interface"`
	Modes []Modes `json:"modes"`
	Asr []Asr `json:"asr"`
}
type Interface struct {
	HTTPRequests float64 `json:"http_requests"`
	HTTPSuccess float64 `json:"http_success"`
	HTTPActive float64 `json:"http_active"`
	WsRequests float64 `json:"ws_requests"`
	WsSuccess float64 `json:"ws_success"`
	WsActive float64 `json:"ws_active"`
	WsConnects float64 `json:"ws_connects"`
}
type Modes struct {
	Mode string `json:"mode"`
	Requests float64 `json:"requests"`
	Success float64 `json:"success"`
	Active float64 `json:"active"`
	Duration float64 `json:"duration"`
}
type Asr struct {
	Cap string `json:"cap"`
	Prop string `json:"prop"`
	Total float64 `json:"total"`
	Current float64 `json:"current"`
	Duration float64 `json:"duration"`
	Cost float64 `json:"cost"`
}




//asr_trans 监控端口数据映射
type AicpAsrTrans struct {
	Trans Trans `json:"trans"`
}
type Trans struct {
	Cap string `json:"cap"`
	Cost float64 `json:"cost"`
	CostSuccess float64 `json:"cost_success"`
	DecodeFailed float64 `json:"decode_failed"`
	DecodeSuccess float64 `json:"decode_success"`
	DecodingCount float64 `json:"decoding_count"`
	Duration float64 `json:"duration"`
	DurationSuccess float64 `json:"duration_success"`
	FileTotal float64 `json:"file_total"`
	Property string `json:"property"`
	ThreadTotal float64 `json:"thread_total"`
	TranscribeFailed float64 `json:"transcribe_failed"`
	TranscribeSuccess float64 `json:"transcribe_success"`
	TranscribingCount float64 `json:"transcribing_count"`
}