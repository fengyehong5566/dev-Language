package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type AicpMtMetric struct {
	//AicpMtResultJson *AicpMt
	AicpMtInterfaceM *prometheus.Desc
	Urlmt string
}


//创建一个新的AsrTransMetric实例
func NewAicpMtCollector(ip, port, fqname string ) *AicpMtMetric {
	//fmt.Println("test: ",fqname)
	if bools := strings.Contains(fqname,"-"); bools{
		fqname = strings.ReplaceAll(fqname,"-","_")
	}

	UrlMt := "http://" + ip +  ":" + port + "/v10/status"
	newDesc1 := prometheus.NewDesc(fqname, "trans_http_download server status", []string{"mt","http"}, nil)

	return &AicpMtMetric{
		//AicpMtResultJson: result,
		AicpMtInterfaceM: newDesc1,
		Urlmt: UrlMt,
	}
}

func (c *AicpMtMetric) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.AicpMtInterfaceM
}

//prometheus收集数据函数
func (c *AicpMtMetric) Collect(metrics chan<- prometheus.Metric) {
	//	interfaces := []string{"http_active", "http_requests", "http_success","ws_active", "ws_connects","ws_requests","ws_success"}
	//	mode := []string{"short_audio","continue_stream", "utterance_stream", "short_stream"}

	Logger.Debugln("start update aicp mt ...")
	AicpMtResultJson := getResultMt(c.Urlmt)
	//映射 "http://10.0.20.82:23001/v10/status?field=asr"的结果到metrics
	num := len(AicpMtResultJson.Mt)
	for i := 0; i < num; i++ {
		metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
			AicpMtResultJson.Mt[i].Chars, AicpMtResultJson.Mt[i].Prop + "_chars", "")
		metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
			AicpMtResultJson.Mt[i].Cost, AicpMtResultJson.Mt[i].Prop + "_cost", "")
		metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
			AicpMtResultJson.Mt[i].Current, AicpMtResultJson.Mt[i].Prop + "_current", "")
		metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
			AicpMtResultJson.Mt[i].Total, AicpMtResultJson.Mt[i].Prop + "_total", "")
	}
	metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
		AicpMtResultJson.HTTP.Active,"", "active")
	metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
		AicpMtResultJson.HTTP.Max,"", "max")
	metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
		AicpMtResultJson.HTTP.Success,"", "success")
	metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
		AicpMtResultJson.HTTP.Request,"", "request")
	metrics <- prometheus.MustNewConstMetric(c.AicpMtInterfaceM, prometheus.GaugeValue,
		AicpMtResultJson.HTTP.ResponseTime,"", "response_time")

	Logger.Debugln("update aicp mt end.")

}


func getResultMt(url string) *AicpMt{
	result := GetAicpMtResultMetric(url)
	return result
}