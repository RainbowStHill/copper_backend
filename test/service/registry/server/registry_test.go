package registry_server_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	registry_server "github.com/rainbowsthill/copper_backend/service/registry/server"
)

func TestRegister(t *testing.T) {
	testCases := map[registry_server.Service]int{
		{
			Type:     registry_server.IDService,
			Addr:     "192.168.10.20",
			Port:     1789,
			ExpireAt: 0,
		}: 3,
		{
			Type:     registry_server.IDService,
			Addr:     "192.168.10.30",
			Port:     1790,
			ExpireAt: 0,
		}: 1,
		{
			Type:     registry_server.IDService,
			Addr:     "192.168.10.40",
			Port:     1791,
			ExpireAt: 0,
		}: 1,
		{
			Type:     registry_server.RegistryService,
			Addr:     "192.168.10.50",
			Port:     1789,
			ExpireAt: 0,
		}: 3,
	}
	for serv, w := range testCases {
		registry_server.GetRegistry().Register(serv, w)
	}
}

func TestDiscover(t *testing.T) {
	testCases := map[registry_server.Service]int{
		{
			Type:     registry_server.IDService,
			Addr:     "192.168.10.20",
			Port:     1789,
			ExpireAt: uint64(time.Now().Add(30 * time.Second).UnixNano()),
		}: 3,
		{
			Type:     registry_server.IDService,
			Addr:     "192.168.10.30",
			Port:     1790,
			ExpireAt: uint64(time.Now().Add(30 * time.Second).UnixNano()),
		}: 1,
		{
			Type:     registry_server.IDService,
			Addr:     "192.168.10.40",
			Port:     1791,
			ExpireAt: uint64(time.Now().Add(30 * time.Second).UnixNano()),
		}: 1,
		{
			Type:     registry_server.RegistryService,
			Addr:     "192.168.10.50",
			Port:     1789,
			ExpireAt: uint64(time.Now().Add(30 * time.Second).UnixNano()),
		}: 3,
	}
	for serv, w := range testCases {
		registry_server.GetRegistry().Register(serv, w)
	}
	loadBalancer := registry_server.GetLoadBalancer()
	var wg sync.WaitGroup
	errs := make(chan error, 100)
	testCasesLock := new(sync.Mutex)
	for c, w := range testCases {
		for ; w > 0; w-- {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if serv, err := loadBalancer.Take(c.Type); err != nil {
					errs <- fmt.Errorf("not enough records of service type %d in load balancer: %v", c.Type, err)
					return
				} else {
					testCasesLock.Lock()
					defer testCasesLock.Unlock()
					if x, ok := testCases[serv]; !ok {
						errs <- fmt.Errorf("Unexpected record %v in load balancer", serv)
						return
					} else if x <= 0 {
						errs <- fmt.Errorf("Too many record %v in load balancer", serv)
						return
					}
					testCases[serv] -= 1
				}
			}()
		}
	}
	wg.Wait()
	if len(errs) > 0 {
		var i = 0
		for err := range errs {
			t.Logf("%dth error: %v", i, err)
			i++
		}
		t.Fatalf("%d errors occurred", i)
	}

	for c, w := range testCases {
		if w != 0 {
			t.Fatalf("Not enough %vs in load balancer", c)
		}
	}
}
