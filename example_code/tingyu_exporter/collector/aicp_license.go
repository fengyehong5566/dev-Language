package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type AicpLicenseMetric struct {
	//AicpLicenseResultJson *AicpLicense
	AicpLicenseDownloadM *prometheus.Desc
	Urllicense string
}


//创建一个新的AsrTransMetric实例
func NewAicpLicenseCollector(ip, port, fqname string ) *AicpLicenseMetric {
	//fLicense.Println("test: ",fqname)
	if bools := strings.Contains(fqname,"-"); bools{
		fqname = strings.ReplaceAll(fqname,"-","_")
	}


	UrlLicense := "http://" + ip +  ":" + port + "/v10/status"
	newDesc2 := prometheus.NewDesc(fqname , "aicp_license server status", []string{"label"}, nil)
	return &AicpLicenseMetric{
		//AicpLicenseResultJson: result,
		AicpLicenseDownloadM: newDesc2,
		Urllicense: UrlLicense,
	}
}

func (c *AicpLicenseMetric) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.AicpLicenseDownloadM
}

//prometheus收集数据函数
func (c *AicpLicenseMetric) Collect(metrics chan<- prometheus.Metric) {
	//	interfaces := []string{"http_active", "http_requests", "http_success","ws_active", "ws_connects","ws_requests","ws_success"}
	//	mode := []string{"short_audio","continue_stream", "utterance_stream", "short_stream"}

	Logger.Debugln("start update aicp license ...")
	AicpLicenseResultJson := getResultLicense(c.Urllicense)
	//映射 "http://10.0.20.82:23001/v10/status?field=asr"的结果到metrics
	num := len(AicpLicenseResultJson.Licenses)
	for i := 0; i < num; i++ {
		metrics <- prometheus.MustNewConstMetric(c.AicpLicenseDownloadM, prometheus.GaugeValue,
			AicpLicenseResultJson.Licenses[i].Current, AicpLicenseResultJson.Licenses[i].Cap)

	}
	Logger.Debugln("update aicp license end .")
}


func getResultLicense(url string) *AicpLicense{

	result := GetAicpLicenseResultMetric(url)
	return result
}