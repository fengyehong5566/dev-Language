package util

import (
	"fmt"
	"testing"
)

func TestReadConf(t *testing.T) {
	exportConf := ReadConf()
	fmt.Println(exportConf)
}