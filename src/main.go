package main

import (
	"os"
	"log"
	"./xcat"
	"fmt"
)

func main(){
	log.SetOutput(os.Stdout)
	args := os.Args
	argsLen := len(args)
	if argsLen < 2{
		log.Fatalln("file args no found. xcat file ")
		os.Exit(1)
	}

	arg := args[1]
	var filePath string
	if argsLen >= 3 && arg == "-x"{
		filePath = args[2]
		writeX(filePath)
	}else{
		filePath = args[1]
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
