package tcp

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestMessage(t *testing.T) {
	listen, _ := net.ListenTCP("tcp", &net.TCPAddr{Port: 8909, IP: net.ParseIP("127.0.0.1")})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {

		defer wg.Done()
		for {
			con, err := listen.AcceptTCP()
			if err != nil {
				panic(err)
			}
			go func() {
				wg.Add(1)
				defer wg.Done()

				data := make([]byte, 100)
				fmt.Println("remote addr :", con.RemoteAddr().String())
				l, err := con.Read(data)
				p := Protocol{}
				msg, err := p.DeCode(data)
				fmt.Printf("len %v , reqId %v, data %v \n", msg.Len, msg.ReqId, string(msg.PayLoad))

				fmt.Println("len ", l, "err", err)
				fmt.Println("data", string(data))
			}()
		}

	}()

	wg.Add(1)
	go func() {
		wg.Add(1)
		defer wg.Done()
		con, _ := net.DialTCP("tcp", nil, &net.TCPAddr{Port: 8909, IP: net.ParseIP("127.0.0.1")})

		data := []byte("123")
		msg := Message{
			120000,
			uint64(len(data)),
			data,
		}
		p := Protocol{}
		data, err := p.EnCode(&msg)
		if err != nil {
			panic(err)
		}

		l, err := con.Write(data)
		fmt.Println("len : ", l, "error", err)
	}()

	wg.Wait()
}
