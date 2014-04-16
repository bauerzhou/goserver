package server

import (
		"errors"
)

type DataHeader struct{
	flow uint32
	cmd  uint32
	ip   [4]byte
}

func (this *DataHeader)WriteToBuf(b []byte) error{
	if len(b) < 12 {
		return errors.New("buf len error")
	}
	
	PutUint32(b, this.flow)
	PutUint32(b[4:], this.cmd)
	copy(b[8:], this.ip[:])
	return nil
}

func (this *DataHeader) ReadFromBuf(b []byte) error{
	if len(b) < 12 {
		return errors.New("buf len error")
	}
	this.flow = Uint32(b[0:4])
	this.cmd  = Uint32(b[4:8])
	copy(this.ip[:], b[8:12])
	return nil
}

func PutUint32(b []byte, v uint32) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func Uint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func Uint16(b []byte) uint16 { return uint16(b[0]) | uint16(b[1])<<8 }

func PutUint16(b []byte, v uint16) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
}


