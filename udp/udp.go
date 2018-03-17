package udp

import (
	"fmt"
	"net"
	"time"
)

func sendUdpData() {
	udpConn, err := net.DialUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0}, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8686})
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
