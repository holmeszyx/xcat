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
	if len(args) < 2{
		log.Println("file args no found. xcat file ")
		os.Exit(1)
	}

	filePath := args[1]
	file, e := os.Open(filePath)
	if e != nil{

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
