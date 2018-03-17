package goroutine

import (
	"fmt"
	"github.com/jackpal/go-nat-pmp"
)

func goroutine() {
	//quit := make(chan int)
	//quit <- 2   无缓存的通道是一直阻塞的。
	fmt.Println("end")
	//<-quit
	gatewayIP, err := gateway.DiscoverGateway()
	if err != nil {
		return
	}

	client := natpmp.NewClient(gatewayIP)
	response, err := client.GetExternalAddress()
	if err != nil {
		return
	}
	fmt.Println("External IP address:", response.ExternalIPAddress)
}
