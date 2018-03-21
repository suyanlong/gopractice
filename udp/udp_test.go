package udp

import (
	"sync"
	"testing"
	"time"
	"fmt"
	"net"
)

func TestSendUdpData(t *testing.T) {
	t.Log("listener")
	
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		sendUdpData()
		t.Log("listener1212")
		wg.Done()
	}()
	
	time.Sleep(1 * time.Second)
	LisenterUdpData()
	wg.Wait()
}

func TestResolveUdpaddr(t *testing.T) {
	udpcConn, err := net.ResolveUDPAddr("udp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println(udpcConn.String())
}
