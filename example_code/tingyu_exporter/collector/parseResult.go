package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"util"
)

var Logger  =  util.CreateLogger()



//请求数据并解析成 AicpLicense 结构体
func GetAicpTtsResultMetric(url string) *AicpTts{
	data, err := getUrl(url)
	if err != nil{
		Logger.Warnf("get aicp_tts data fail. err=%s\n",err.Error())
	}
	resultJson := &AicpTts{}
	err = json.Unmarshal(data,resultJson)
	if err !=nil {
		Logger.Debugln("GetAicpTtsResultMetric.json.Unmarshal fail. err=",err)
	}
	return resultJson
}

//请求数据并解析成 AicpLicense 结构体
func GetAicpLicenseResultMetric(url string) *AicpLicense{
	data,err := getUrl(url)
	if err != nil{
		Logger.Warnf("get aicp_license data fail. err=%s\n",err.Error())
	}

	resultJson := &AicpLicense{}
	err = json.Unmarshal(data,resultJson)

	if err !=nil {
		Logger.Debugln("GetAicpLicenseResultMetric.json.Unmarshal fail. err=",err)
	}
	return resultJson
}

//请求数据并解析成 AicpMt 结构体
func GetAicpMtResultMetric(url string) *AicpMt{
	data,err := getUrl(url)
	if err != nil{
		Logger.Warnf("get aicp_mt data fail. err=%s\n",err.Error())
	}
	resultJson := &AicpMt{}
	err = json.Unmarshal(data,resultJson)
	if err !=nil {
		Logger.Debugln("GetAicpMtResultMetric.json.Unmarshal fail. err=",err)
	}
	return resultJson
}


//请求数据并解析成 AicpApiGw 结构体
func GetTransHttpResultMetric(url string) *AicpTransHttp{
	data,err := getUrl(url)
	if err != nil{
		Logger.Warnf("get trans_http data fail. err=%s\n",err.Error())
	}
	resultJson := &AicpTransHttp{}
	err = json.Unmarshal(data,resultJson)
	if err !=nil {
		Logger.Debugln("GetTransHttpResultMetric.json.Unmarshal fail. err=",err)
	}
	return resultJson
}

//请求数据并解析成 AicpApiGw 结构体
func GetAicpApiGwResultMetric(url string) *AicpApiGw{
	data,err  := getUrl(url)
	if err != nil{
		Logger.Warnf("get aicp_apigw data fail. err=%s\n",err.Error())
	}
	resultJson := &AicpApiGw{}
	err = json.Unmarshal(data,resultJson)
	if err !=nil {
		Logger.Debugln("GetAicpApiGwResultMetric.json.Unmarshal fail. err=",err)
	}
	return resultJson
}

//请求数据并解析成 AsrTrans 结构体
func GetAsrTransResultMetric(url string) *AicpAsrTrans{
	data,err  := getUrl(url)
	if err != nil{
		Logger.Warnf("get asr_trans data fail. err=%s\n",err.Error())
	}
	resultJson := &AicpAsrTrans{}
	err = json.Unmarshal(data,resultJson)
	if err !=nil {
		Logger.Debugln("GetAsrTransResultMetric.json.Unmarshal fail. err=",err)
	}
	return resultJson
}


//请求数据并解析成 ResultJson 结构体
func GetAsrFtResultMetric(url string) *AicpAsrFT{
	data,err := getUrl(url)
	if err != nil{
		Logger.Warnf("get asr_ft data fail. err=%s\n",err.Error())
	}
	resultjson := &AicpAsrFT{}
	err = json.Unmarshal(data,resultjson)
	if err !=nil {
		Logger.Debugln("GetAsrFtResultMetric.json.Unmarshal fail. err=",err)
	}
	return resultjson
}


func getUrl(url string) (data []byte, err error) {
	result,err := http.Get(url)
	if err != nil{
		Logger.Warnf("http.Get(%s) fail. err=%s\n",url,err.Error())
		return []byte{}, err
	}
	defer result.Body.Close()

	data, err = ioutil.ReadAll(result.Body)
	if err !=nil {
		Logger.Infoln("getUrl.ioutil.ReadAll() fail. err=",err)
		return []byte{}, err
	}
	return
}