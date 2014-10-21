package main

import (
	"./xcat"
	"flag"
	"io"
	"log"
	"os"
)

var outToFile *os.File

func main() {
	log.SetOutput(os.Stdout)
	var isWriteX bool
	var outFilePath string
	flag.BoolVar(&isWriteX, "x", false, "encode by x stream")
	flag.StringVar(&outFilePath, "o", "", "Out put file")
	flag.Parse()
	argCount := flag.NArg()

	var filePath string
	if argCount > 0 {
		filePath = flag.Arg(0)
	} else {
		log.Fatalln("file args no found. xcat file ")
		os.Exit(1)
	}

	if outFilePath != "" {
		f, err := os.Create(outFilePath)
		if err != nil {
			log.Fatalln("Can not create out put file ", outFilePath)
			os.Exit(2)
		}
		outToFile = f
		if outToFile != nil {
			defer outToFile.Close()
		}
	}

	if isWriteX {
		writeX(filePath)
	} else {
		readX(filePath)
	}
}

// 获取当前的输出流
// 文件获取标输出
func getOutWriter() (w io.Writer) {
	if outToFile != nil {
		w = outToFile
	} else {
		w = os.Stdout
	}
	return
}

// 读异或流
func readX(filePath string) {
	file, e := os.Open(filePath)
	if e != nil {
		log.Fatalln("file open error")
	}
	defer file.Close()
	xReader := xcat.NewXReader(file)
	buff := make([]byte, 4096)

	var currentWrite io.Writer = getOutWriter()

	for {
		n, err := xReader.Read(buff)
		if err != nil {
			break
		}

		currentWrite.Write(buff[:n])
	}
}

func writeX(filePath string) {
	file, e := os.Open(filePath)
	if e != nil {
		log.Fatalln("file open error")
	}
	defer file.Close()

	var currentWrite io.Writer = getOutWriter()
	xWriter := xcat.NewXWriter(currentWrite)

	buff := make([]byte, 4096)
	for {
		n, err := file.Read(buff)
		if err != nil {
			break
		}
		xWriter.Write(buff[:n])
	}
}
