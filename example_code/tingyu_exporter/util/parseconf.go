package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//读取配置文件，并返回ExprotConf结构体，
func ReadConf() ExportConf {
	confByte, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		panic("open conf/conf.yaml fail. err=" + err.Error())
	}
	exportconf := ExportConf{}
	err = yaml.Unmarshal(confByte, &exportconf)
	if err != nil {
		panic("ReadConf.yaml.Unmarshal() fail. err:"  + err.Error())
	}
	return exportconf
}

//map yaml配置文件
type ExportConf struct {
	PidPath string   `yaml:"pidpath"`
	HostIp string `yaml:"hostip"`
	ServicePort string `yaml:"servirce_port"`
	LogLevel string `yaml:"log_level"`
	Metrics    []metric `yaml:"metrics"`
}


type metric struct {
	Name string `yaml:"metric"`
	ServiceType string `yaml:"service_type"`
	MonitorPort string `yaml:"monitor_port"`
}
