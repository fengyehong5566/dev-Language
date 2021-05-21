package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type AicpTtsMetric struct {
	//AicpTtsResultJson *AicpTts
	AicpTtsTTSM *prometheus.Desc
	Urltts string
}



//创建一个新的AicpTtsMetric实例
func NewAicpTtsCollector(ip, port, fqname string ) *AicpTtsMetric {
	Logger.Infoln("init AicpTtsMetric.")
	if bools := strings.Contains(fqname,"-"); bools{
		fqname = strings.ReplaceAll(fqname,"-","_")
	}
	newDesc1 := prometheus.NewDesc(fqname, "aicp_tts_http server status", []string{"http", "tts","websocket"}, nil)
	return &AicpTtsMetric{
		//AicpTtsResultJson: result,
		AicpTtsTTSM: newDesc1,
		Urltts:"http://" + ip +  ":" + port + "/v10/status",
	}
}

func (c *AicpTtsMetric) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.AicpTtsTTSM
}

//prometheus收集数据函数
func (c *AicpTtsMetric) Collect(metrics chan<- prometheus.Metric) {
	//映射 "http://10.0.20.82:23001/v10/status?field=asr"的结果到metrics

	Logger.Debugln("start update tts ...")
	AicpTtsResultJson := getResultTts(c.Urltts)
	num := len(AicpTtsResultJson.Tts)
	//fmt.Println("tts num: ",AicpTtsResultJson.Tts)
	for i := 0; i < num; i++ {
		metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
			AicpTtsResultJson.Tts[i].Chars, "", AicpTtsResultJson.Tts[i].Prop + "_chars", "")
		metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
			AicpTtsResultJson.Tts[i].Cost, "", AicpTtsResultJson.Tts[i].Prop + "_cost", "")
		metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
			AicpTtsResultJson.Tts[i].Current, "", AicpTtsResultJson.Tts[i].Prop + "_current", "")
		metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
			AicpTtsResultJson.Tts[i].Total, "", AicpTtsResultJson.Tts[i].Prop + "_total", "")
	}
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Active,"active", "", "")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Max,"max", "", "")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Success,"success", "", "")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Request,"request", "", "")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.ResponseTime,"response_time", "", "")

	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Active,"", "", "active")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Max,"", "", "max")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Success,"", "", "success")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.Request,"", "", "request")
	metrics <- prometheus.MustNewConstMetric(c.AicpTtsTTSM, prometheus.GaugeValue,
		AicpTtsResultJson.TtsHttp.ResponseTime,"", "", "response_time")

	Logger.Debugln("update tts end.")
}



func getResultTts(url string) *AicpTts{
	result := GetAicpTtsResultMetric(url)
	fmt.Println("tts result: ",result)
	return result
}