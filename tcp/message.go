package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

///Message Custom protocol msg type
type Message struct {
	ReqId   uint64
	Len     uint64
	PayLoad []byte
}

type Protocol struct {
	reqId uint64
}

func (p *Protocol) DeCode(data []byte) (*Message, error) {
	if len(data) < 8*2 {
		return nil, errors.New("length too small !")
	}

	b := binary.BigEndian
	_d := data[0:8]
	reqId := b.Uint64(_d)
	length := b.Uint64(data[8:16])
	l := uint64(len(data[16:]))

	if l < length {
		return nil, errors.New("length too small!")

	}
	//else if l > length {
	//	return nil, errors.New("length too large!")
	//}

	//TODO BUG 字节是否还需要转化成大小端模式呢？
	//reader := bytes.NewReader(data[16:])
	//binary.Read(reader,b,data[16:])

	return &Message{
		ReqId:   reqId,
		Len:     l,
		PayLoad: data[16:],
	}, nil
}

func (p *Protocol) EnCode(msg *Message) ([]byte, error) {
	var (
		b    = binary.BigEndian
		data = make([]byte, 16)
	)

	b.PutUint64(data, msg.ReqId)
	b.PutUint64(data[8:], msg.Len)

	data = append(data, msg.PayLoad...)
	return data, nil
}
