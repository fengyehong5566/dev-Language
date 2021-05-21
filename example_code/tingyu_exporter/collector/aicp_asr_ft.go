package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type AsrFtMetrics struct {
	//resultJson *AicpAsrFT
	AsrFT      *prometheus.Desc
	Urlasrft string
}


//创建一个新的AsrFt实例
func NewAsrFtCollector(ip, port, fqname string) *AsrFtMetrics {
	//fmt.Println("test: ",fqname)
	if bools := strings.Contains(fqname,"-"); bools{
		fqname = strings.ReplaceAll(fqname,"-","_")
	}

	UrlAsrft := "http://" + ip + ":" + port + "/v10/status"
	asrFt := prometheus.NewDesc(fqname, "asr_ft server info", []string{"asr", "interface", "mode"}, nil)
	return &AsrFtMetrics{
		//resultJson: asrFTResult,
		AsrFT:      asrFt,
		Urlasrft:	UrlAsrft,
	}
}

func (c *AsrFtMetrics) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.AsrFT
}

//prometheus收集数据函数
func (c *AsrFtMetrics) Collect(metrics chan<- prometheus.Metric) {
	//	interfaces := []string{"http_active", "http_requests", "http_success","ws_active", "ws_connects","ws_requests","ws_success"}
	//	mode := []string{"short_audio","continue_stream", "utterance_stream", "short_stream"}

	resultJson := getResult(c.Urlasrft)
	Logger.Debugf("start update %s:%s  ...", resultJson.Asr[0].Cap,resultJson.Asr[0].Prop)
	//映射 "http://10.0.20.82:23001/v10/status?field=asr"的结果到metrics
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Asr[0].Cost, "cost", "", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Asr[0].Current, "current", "", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Asr[0].Duration, "duration", "", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Asr[0].Total, "total", "", "")

	//映射 "http://10.0.20.82:23001/v10/status?field=interfaces"的结果到metrics
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Interfaces.HTTPActive, "", "httpcative", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Interfaces.HTTPRequests, "", "httprequests", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Interfaces.HTTPSuccess, "", "httpsuccess", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Interfaces.WsActive, "", "wsactive", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Interfaces.WsConnects, "", "wsconnects", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Interfaces.WsRequests, "", "wsrequests", "")
	metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Interfaces.WsSuccess, "", "wssuccess", "")

	//映射 "http://10.0.20.82:23001/v10/status?field=modes"的结果到metrics
	ModesLen := len(resultJson.Modes)
	for i := 0; i < ModesLen; i++ {
		v := i
		metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Modes[v].Active, "", "", resultJson.Modes[v].Mode+"_active")
		metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Modes[v].Duration, "", "", resultJson.Modes[v].Mode+"_duration")
		metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Modes[v].Requests, "", "", resultJson.Modes[v].Mode+"_requests")
		metrics <- prometheus.MustNewConstMetric(c.AsrFT, prometheus.GaugeValue, resultJson.Modes[v].Success, "", "", resultJson.Modes[v].Mode+"_success")
	}
	Logger.Debugf("update %s:%s  end.", resultJson.Asr[0].Cap,resultJson.Asr[0].Prop)

}


func getResult(url string) *AicpAsrFT{
	asrFTResult := GetAsrFtResultMetric(url)
	return asrFTResult
}