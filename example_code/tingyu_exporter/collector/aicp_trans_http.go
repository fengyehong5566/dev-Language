package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type TransHttpMetric struct {
	//TransHttpResultJson *AicpTransHttp
	TransHttpInterfaceM *prometheus.Desc
	Urltranshttp string

}


//创建一个新的AsrTransMetric实例
func NewTransHttpCollector(ip, port, fqname string ) *TransHttpMetric {
	//fmt.Println("test: ",fqname)
	if bools := strings.Contains(fqname,"-"); bools{
		fqname = strings.ReplaceAll(fqname,"-","_")
	}

	UrlTranshttp := "http://" + ip +  ":" + port + "/v10/status"
	newDesc1 := prometheus.NewDesc(fqname + "_download", "trans_http_download server status", []string{"interface","download"}, nil)

	return &TransHttpMetric{
		//TransHttpResultJson: result,
		TransHttpInterfaceM: newDesc1,
		Urltranshttp: UrlTranshttp,
	}
}



func (c *TransHttpMetric) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.TransHttpInterfaceM
}

//prometheus收集数据函数
func (c *TransHttpMetric) Collect(metrics chan<- prometheus.Metric) {
	//	interfaces := []string{"http_active", "http_requests", "http_success","ws_active", "ws_connects","ws_requests","ws_success"}
	//	mode := []string{"short_audio","continue_stream", "utterance_stream", "short_stream"}

	Logger.Debugln("start update trans http ...")
	TransHttpResultJson := getResultTransHttp(c.Urltranshttp)
	//映射 "http://10.0.20.82:23001/v10/status?field=asr"的结果到metrics
	num := len(TransHttpResultJson.TransInterfaces)
	for i := 0; i < num; i++ {
		metrics <- prometheus.MustNewConstMetric(c.TransHttpInterfaceM, prometheus.GaugeValue,
			TransHttpResultJson.TransInterfaces[i].CancelFiles, TransHttpResultJson.TransInterfaces[i].Property + "_cancel_files", "")
		metrics <- prometheus.MustNewConstMetric(c.TransHttpInterfaceM, prometheus.GaugeValue,
			TransHttpResultJson.TransInterfaces[i].CancelTasks, TransHttpResultJson.TransInterfaces[i].Property + "_cancel_tasks", "")
		metrics <- prometheus.MustNewConstMetric(c.TransHttpInterfaceM, prometheus.GaugeValue,
			TransHttpResultJson.TransInterfaces[i].SubmitFiles, TransHttpResultJson.TransInterfaces[i].Property + "_submit_files", "")
		metrics <- prometheus.MustNewConstMetric(c.TransHttpInterfaceM, prometheus.GaugeValue,
			TransHttpResultJson.TransInterfaces[i].SubmitTasks, TransHttpResultJson.TransInterfaces[i].Property + "_submit_tasks", "")
	}
	metrics <- prometheus.MustNewConstMetric(c.TransHttpInterfaceM, prometheus.GaugeValue,
		TransHttpResultJson.Download.Downloading,"", "downloading")
	metrics <- prometheus.MustNewConstMetric(c.TransHttpInterfaceM, prometheus.GaugeValue,
		TransHttpResultJson.Download.Failed,"", "failed")
	metrics <- prometheus.MustNewConstMetric(c.TransHttpInterfaceM, prometheus.GaugeValue,
		TransHttpResultJson.Download.Success,"", "success")

	Logger.Debugln("update trans http end.")

}


func getResultTransHttp(url string) *AicpTransHttp{
	result := GetTransHttpResultMetric(url)
	return result
}