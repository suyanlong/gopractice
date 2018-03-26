package p2p

import (
	"fmt"
	"github.com/jackpal/gateway"
	"github.com/jackpal/go-nat-pmp"
	"net"
	"os/exec"
)

func nat() {

	ip, err := net.LookupIP("")
	fmt.Println("ip = ", ip)

	if out, err := exec.Command("ifconfig").CombinedOutput(); err == nil {
		fmt.Println(string(out))
	}

	/// Get gateway address!
	gatewayIP, err := gateway.DiscoverGateway()
	if err != nil {
		return
	}
	fmt.Println("gatewayIP:", gatewayIP.String())

	client := natpmp.NewClient(gatewayIP)
	response, err := client.GetExternalAddress()
	if err != nil {
		return
	}

	fmt.Printf("External IP address:%v, SecondsSinceStartOfEpoc = %v ",
		response.ExternalIPAddress, response.SecondsSinceStartOfEpoc)

}
