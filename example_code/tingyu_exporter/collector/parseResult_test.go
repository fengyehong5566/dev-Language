package collector

import (
	"fmt"
	"testing"
)


func TestGetAsrFtResultMetric(t *testing.T) {
	Url := ""
	resjson := GetAsrFtResultMetric()
	fmt.Println(resjson)
}
