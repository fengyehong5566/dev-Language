package resource

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"strconv"
	"util"
)

var Logger log.Logger

type CpuMetrics struct {

	CpuPercentM *prometheus.Desc
	MemPercentM *prometheus.Desc
}

var Conf util.ExportConf

func NewStatusOfServiceCollector(conf util.ExportConf) *CpuMetrics{
	Conf = conf
	Logger = util.CreateLogger()
	Logger.Debugln("NewStatusOfServiceCollector CpuMetrics init.")
	cpuPercentM := prometheus.NewDesc("CpuPercent","serivce used cpu info",[]string{"cpu"},nil)
	memPercentM := prometheus.NewDesc("MemPercent","serivce used mem info",[]string{"mem"},nil)
	Logger.Debugln("NewStatusOfServiceCollector CpuMetrics end.")
	return &CpuMetrics{
		CpuPercentM: cpuPercentM,
		MemPercentM: memPercentM,
	}
}

func (c *CpuMetrics) Describe(descs chan <- *prometheus.Desc){
	descs <- c.CpuPercentM
	descs <- c.MemPercentM

}

func (c *CpuMetrics) Collect(metrics chan<- prometheus.Metric) {
	Logger.Infoln("start update cpu and mem  ...")
	num := len(Conf.Metrics)
	for i := 0; i < num; i++ {
		metrics <- prometheus.MustNewConstMetric(c.CpuPercentM, prometheus.GaugeValue, getCpu(Conf.PidPath+"/"+Conf.Metrics[i].Name+".pid"), Conf.Metrics[i].Name)
		metrics <- prometheus.MustNewConstMetric(c.MemPercentM, prometheus.GaugeValue, getMem(Conf.PidPath+"/"+Conf.Metrics[i].Name+".pid"), Conf.Metrics[i].Name)
	}
	Logger.Infoln("update cpu and mem  end.")
}

var pidProc *process.Process


func getMem(filename string) float64{
	var memorypercentT float64
	var memory uint64
	initProcess(filename)
	memoryInfo,err  := pidProc.MemoryInfo()
	if err != nil {
		memorypercentT = 0
	} else {
		memory = memoryInfo.RSS
	}
	memorypercentT = float64(memory)
	if err != nil {
		Logger.Warnln("getMem.pidProc.MemoryPercent() fail, err=",err)
		return 0
	}
	Logger.Infof("getMem.pidProc.MEMPercent(%v), value=%v\n",pidProc.Pid, memorypercentT)
	memorypercent := float64(memorypercentT)
	// memorypercent / 1024 / 1024 /1024.0
	return memorypercent/1073741824.0
}


func getCpu(filename string) float64{
	initProcess(filename)
	cpupercent,err  := pidProc.CPUPercent()

	if err != nil {
		Logger.Warnf("getCpu.pidProc.CPUPercent(%v) fail, err=%s\n",pidProc.Pid,err)
		return 0
	}
	Logger.Infof("getCpu.pidProc.CPUPercent(%v), value=%v\n",pidProc.Pid,cpupercent)
	return cpupercent
}

func initProcess(servicename string){
	pid,err := ioutil.ReadFile(servicename)
	if err != nil {
		Logger.Errorln("Get Pid fail from pid file.")
		pid =[]byte{}
	}
	pidN,_ := strconv.Atoi(string(pid))


	pidProc,err = process.NewProcess(int32(pidN))
	if err != nil {
		fmt.Println("initProcess.process.Newprocess() fail, err=",err)
	}
}

