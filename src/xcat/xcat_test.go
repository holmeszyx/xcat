package xcat

import (
	"testing"
	"os"
)

func TestReader(t *testing.T){
	xFile, err := os.Open("/home/holmes/xout.test")
	if err != nil{
		t.Error("file no found")
	}
	defer xFile.Close()

	t.Log("start")

	xReader := NewXReader(xFile)
	buff := make([]byte, 4096)
	for {
		n, e := xReader.Read(buff)
		if e != nil{
			break
		}
		t.Log(string(buff[:n]))
	}

}
