package udp

import (
	"testing"
	"sync"
	"time"
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
