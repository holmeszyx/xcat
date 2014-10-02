package xcat

import (
	"io"
	"errors"
)

const (
	KEY = byte(0x73)
	posOffset = 2
)

// 异或读流
type XReader struct {
	in io.Reader
	pos uint32
}

// 新建一个异或读流
func NewXReader(in io.Reader) (x *XReader){
	x = &XReader{in, posOffset}
	return
}

// 读
func (x *XReader) Read(p []byte) (n int, e error){
	n, e = x.in.Read(p)
	if e != nil{
		return
	}

	for i := 0; i < n; i ++{

		po := (x.pos * x.pos * x.pos) >> 1
		dk := byte(po)

		f := dk ^ KEY ^ p[i]
		p[i] = f
		x.pos ++
	}

	return
}

// 异或写流
type XWriter struct {
	out io.Writer
	pos uint32
}

// 新建一个异或写流
func NewXWriter(out io.Writer) (x *XWriter){
	x = &XWriter{out, posOffset}
	return
}

// 写
func (x *XWriter)Write(p []byte) (n int, err error){
	pLen := len(p)
	if pLen <= 0{
		return 0, errors.New("write a empty buff")
	}

	for i := 0; i < pLen; i ++{
		data := p[i]

		po := (x.pos * x.pos * x.pos) >> 1
		dk := byte(po)
		f := dk ^ KEY ^ data

		p[i] = f
		x.pos ++
	}

	n, err = x.out.Write(p)
	return
}


