package main

import (
	"os"
	"log"
	"./xcat"
	"fmt"
	"flag"
)

func main(){
	log.SetOutput(os.Stdout)
	var isWriteX bool
	flag.BoolVar(&isWriteX, "x", false, "encode by x stream")
	flag.Parse()
	argCount := flag.NArg()

	var filePath string
	if argCount > 0{
		filePath = flag.Arg(0)
	}else{
		log.Fatalln("file args no found. xcat file ")
		os.Exit(1)
	}

	if isWriteX{
		writeX(filePath)
	}else{
		readX(filePath)
	}
}

// 读异或流
func readX(filePath string){
	file, e := os.Open(filePath)
	if e != nil{
		log.Fatalln("file open error")
	}
	defer file.Close()
	xReader := xcat.NewXReader(file)
	buff := make([]byte, 4096)
	for {
		n, err := xReader.Read(buff)
		if err != nil{
			break;
		}

		fmt.Println(string(buff[:n]))
	}
}

func writeX(filePath string){
	file, e := os.Open(filePath)
	if e != nil{
		log.Fatalln("file open error")
	}
	defer file.Close()

	xWriter := xcat.NewXWriter(os.Stdout)

	buff := make([]byte, 4096)
	for{
		n, err := file.Read(buff)
		if err != nil{
			break;
		}
		xWriter.Write(buff[:n])
	}
}
