package main

import (
	"collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"resource"
	"util"
)
var Conf util.ExportConf
var Logger  log.Logger

func init(){
	Conf = util.ReadConf()
	Logger = util.CreateLogger()
}

func main() {
	MustRest()
	http.Handle("/metrics", promhttp.Handler())

	//log.Fatalln(http.ListenAndServe(Conf.HostIp + ":" + Conf.ServicePort, nil))
	log.Fatalln(http.ListenAndServe(":" + Conf.ServicePort, nil))
}


func MustRest() {
	Logger.Debugln("initialize metrics of cpu/mem start...")
	registM(Conf)

	Logger.Debugln("initialize metrics of services start...")

	num := len(Conf.Metrics)
	for i := 0; i < num; i++ {

		switch Conf.Metrics[i].ServiceType {
		case "aicp_apigw", "aicp_apigw_tls":
			Logger.Debugf("Register aicp_apigw %s",Conf.Metrics[i].Name)
			prometheus.MustRegister(collector.NewAicpApiGwCollector(Conf.HostIp, Conf.Metrics[i].MonitorPort, Conf.Metrics[i].Name))

		case "aicp_asr_ft":
			Logger.Debugf("Register aicp_asr_ft %s",Conf.Metrics[i].Name)
			prometheus.MustRegister(collector.NewAsrFtCollector(Conf.HostIp, Conf.Metrics[i].MonitorPort, Conf.Metrics[i].Name))

		case "aicp_asr_trans":
			Logger.Debugf("Register aicp_asr_trans %s",Conf.Metrics[i].Name)
			prometheus.MustRegister(collector.NewAsrTransCollector(Conf.HostIp, Conf.Metrics[i].MonitorPort, Conf.Metrics[i].Name))

		case "aicp_license":
			Logger.Debugf("Register aicp_license %s",Conf.Metrics[i].Name)
			prometheus.MustRegister(collector.NewAicpLicenseCollector(Conf.HostIp, Conf.Metrics[i].MonitorPort, Conf.Metrics[i].Name))

		case "aicp_mt":
			Logger.Debugf("Register aicp_mt %s",Conf.Metrics[i].Name)
			prometheus.MustRegister(collector.NewAicpMtCollector(Conf.HostIp, Conf.Metrics[i].MonitorPort, Conf.Metrics[i].Name))

		case "aicp_trans_http":
			Logger.Debugf("Register aicp_trans_http %s",Conf.Metrics[i].Name)
			prometheus.MustRegister(collector.NewTransHttpCollector(Conf.HostIp, Conf.Metrics[i].MonitorPort, Conf.Metrics[i].Name))

		case "aicp_tts":
			Logger.Debugf("Register aicp_tts %s",Conf.Metrics[i].Name)
			go prometheus.MustRegister(collector.NewAicpTtsCollector(Conf.HostIp, Conf.Metrics[i].MonitorPort, Conf.Metrics[i].Name))

		case "consul":
			return
		default:
			Logger.Warnf("%s Parameter configuration error, plz check configfile.", Conf.Metrics[i].ServiceType)

		}

	}
	Logger.Debugln("initialize metrics end ...")

}

func registM (Conf util.ExportConf){
	prometheus.MustRegister(resource.NewStatusOfServiceCollector(Conf))
}