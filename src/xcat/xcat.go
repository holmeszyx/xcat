package xcat

import (
	"io"
//	"encoding/binary"
)

const (
	KEY = byte(0x73)
	posOffset = 2
)

type XReader struct {
	in io.Reader
	pos uint32
}

func NewXReader(in io.Reader) (x *XReader){
	x = &XReader{in, posOffset}
	return
}

func (x *XReader) Read(p []byte) (n int, e error){
	n, e = x.in.Read(p)
	if e != nil{
		return
	}

	for i := 0; i < n; i ++{

		po := (x.pos * x.pos * x.pos) >> 1
//		pe := make([]byte, 4)
//		binary.LittleEndian.PutUint32(pe, po)
//		po = binary.BigEndian.Uint32(pe)
		dk := byte(po)

		f := dk ^ KEY ^ p[i]
		p[i] = f
		x.pos ++
	}

	return
}
