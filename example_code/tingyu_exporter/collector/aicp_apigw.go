package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type AicpApiGwMetric struct {
	//AicpApiGwResultJson *AicpApiGw
	AicpApiGwM *prometheus.Desc
	Urlapigw string
}



//创建一个新的AsrTransMetric实例
func NewAicpApiGwCollector(ip, port, fqname string ) *AicpApiGwMetric {
	//fmt.Println("test: ",fqname)
	if bools := strings.Contains(fqname,"-"); bools{
		fqname = strings.ReplaceAll(fqname,"-","_")
	}


	urlApigw := "http://" + ip +  ":" + port + "/v10/status"
	newDesc := prometheus.NewDesc(fqname, "aicp_apigw server status", []string{"label"}, nil)
	return &AicpApiGwMetric{
		//AicpApiGwResultJson: result,
		AicpApiGwM:          newDesc,
		Urlapigw:	urlApigw,
	}
}

func (c *AicpApiGwMetric) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.AicpApiGwM
}

//prometheus收集数据函数
func (c *AicpApiGwMetric) Collect(metrics chan<- prometheus.Metric) {
	//	interfaces := []string{"http_active", "http_requests", "http_success","ws_active", "ws_connects","ws_requests","ws_success"}
	//	mode := []string{"short_audio","continue_stream", "utterance_stream", "short_stream"}

	Logger.Debugln("start update aicp_apigw ..." )
	AicpApiGwResultJson := getResultapigw(c.Urlapigw)

	//映射 "http://10.0.20.82:23001/v10/status?field=asr"的结果到metrics
	num := len(AicpApiGwResultJson.API)
	for i := 0; i < num; i++ {
		metrics <- prometheus.MustNewConstMetric(c.AicpApiGwM, prometheus.GaugeValue,
			AicpApiGwResultJson.API[i].Active, AicpApiGwResultJson.API[i].Name + "_active")
		metrics <- prometheus.MustNewConstMetric(c.AicpApiGwM, prometheus.GaugeValue,
			AicpApiGwResultJson.API[i].Requests, AicpApiGwResultJson.API[i].Name + "_request")
		metrics <- prometheus.MustNewConstMetric(c.AicpApiGwM, prometheus.GaugeValue,
			AicpApiGwResultJson.API[i].Success, AicpApiGwResultJson.API[i].Name + "_success")
	}
	Logger.Debugln("update aicp_apigw end..." )
}

func getResultapigw(url string) *AicpApiGw{
	result := GetAicpApiGwResultMetric(url)
	return result
}