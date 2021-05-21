package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type AsrTransMetric struct {
	//AsrTransResultJson *AicpAsrTrans
	AsrTransM *prometheus.Desc
	Urlasrtrans string
}



//创建一个新的AsrTransMetric实例
func NewAsrTransCollector(ip, port, fqname string ) *AsrTransMetric {
	if bools := strings.Contains(fqname,"-"); bools{
		fqname = strings.ReplaceAll(fqname,"-","_")
	}

	UrlAsrtrans := "http://" + ip +  ":" + port + "/v10/status"
	newDesc := prometheus.NewDesc(fqname, "asr_trans server status", []string{"trans"}, nil)
	return &AsrTransMetric{
		//AsrTransResultJson: result,
		AsrTransM:          newDesc,
		Urlasrtrans: UrlAsrtrans,
	}
}

func (c *AsrTransMetric) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.AsrTransM
}

//prometheus收集数据函数
func (c *AsrTransMetric) Collect(metrics chan<- prometheus.Metric) {
	//	interfaces := []string{"http_active", "http_requests", "http_success","ws_active", "ws_connects","ws_requests","ws_success"}
	//	mode := []string{"short_audio","continue_stream", "utterance_stream", "short_stream"}
	Logger.Debugln("start update asrtrans ...")
	AsrTransResultJson := getResultAsrTrans(c.Urlasrtrans)
	//映射 "http://10.0.20.82:23001/v10/status?field=asr"的结果到metrics
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM, prometheus.GaugeValue, AsrTransResultJson.Trans.Cost, "cost")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM, prometheus.GaugeValue, AsrTransResultJson.Trans.CostSuccess, "cost_success")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM, prometheus.GaugeValue, AsrTransResultJson.Trans.DecodeFailed, "decode_failed")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM, prometheus.GaugeValue, AsrTransResultJson.Trans.DecodeSuccess, "decode_success")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM,prometheus.GaugeValue, AsrTransResultJson.Trans.DecodingCount,"decoding_count")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM,prometheus.GaugeValue, AsrTransResultJson.Trans.Duration,"duration")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM,prometheus.GaugeValue, AsrTransResultJson.Trans.DurationSuccess,"duration_success")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM,prometheus.GaugeValue, AsrTransResultJson.Trans.FileTotal,"file_total")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM,prometheus.GaugeValue, AsrTransResultJson.Trans.ThreadTotal,"thread_total")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM,prometheus.GaugeValue, AsrTransResultJson.Trans.TranscribeFailed,"transcribe_failed")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM,prometheus.GaugeValue, AsrTransResultJson.Trans.TranscribeSuccess,"transcribe_success")
	metrics <- prometheus.MustNewConstMetric(c.AsrTransM, prometheus.GaugeValue, AsrTransResultJson.Trans.TranscribingCount, "transcribing_count")

	Logger.Debugln("update asrtrans end.")
}


func getResultAsrTrans(url string) *AicpAsrTrans{
	result := GetAsrTransResultMetric(url)
	return result
}

