package udp

import (
	"fmt"
	"net"
	"time"
)

func sendUdpData() {
	/// laddr 本地地址， raadr远程地址。
	udpConn, err := net.DialUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 6786}, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8686})
	fmt.Println(udpConn.LocalAddr().String())
	if err != nil {
		fmt.Printf("udp dial err %v", err)
	}
	for {
		udpConn.Write([]byte("12312"))
		fmt.Println("send data")
		time.Sleep(1 * time.Second)
	}

	defer udpConn.Close()

}

func LisenterUdpData() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8686})

	fmt.Println("listener udp data ")
	fmt.Println(udpConn.LocalAddr().String())
	if err != nil {
		fmt.Printf("udp dial err %v", err)
	}

	data := make([]byte, 1024)
	for {
		udpConn.Read(data)
		fmt.Println(string(data))
	}

	fmt.Println("LisenterUdpData end! ")
	defer udpConn.Close()
}
