package udp

import (
	"fmt"
	"net"
	"time"
)

//随着防火墙和NAT的广泛部署，应用程序也开始加以适应，演变出了多种对付IP地址不唯一，单向通信的问题。目前，应用的NAT穿越已经不再是一个重要议题了。
//
//第一类方法，对NAT做某种方式的静态配置，执行端口的打开/转发等。
//
//第二类方法，间接地使用一些外部服务，来协助建立透过NAT的连接，如STUN和TURN。
//
//第三类方法，与NAT直接配合，动态打开端口，如UPnP-IGD(UPnP网关设备)和NAT-PMP(NAT端口映射协议)。
//
//第四类方法，当使用两级NAT或DS-lite的CGN模式下，破坏了NAT与运行客户端程序的主机位于同一链路上的，UPnP或NAT-PMP的这一基本假设。就要靠IETF正在做的PCP(端口控制协议)协议了。

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
