package registry_server

import (
	"fmt"
	"sync"
	"time"
)

// LoadBalancer is a producer-consumer model. Producers add services' infomation to it
// and consumers take the information.
type LoadBalancer interface {
	Add(services ...Service)
	Take(serviceType ServiceType) (Service, error)
}

// WeightedLRULoadBalancer uses weighted least resent used algorithm.
type WeightedLRULoadBalancer struct {
	reg       *registry
	publisher map[ServiceType]chan Service
}

func (lb WeightedLRULoadBalancer) Take(serviceType ServiceType) (Service, error) {
	if len(lb.publisher[serviceType]) <= 0 {
		lb.reg.Discover(serviceType, &lb)
	}
	if len(lb.publisher[serviceType]) <= 0 {
		return Service{}, ErrServiceInfoNotExist
	} else if serv, ok := <-lb.publisher[serviceType]; !ok {
		return Service{}, ErrServiceInfoNotExist
	} else {
		return serv, nil
	}
}

func (lb WeightedLRULoadBalancer) Add(services ...Service) {
	for _, serv := range services {
		now := uint64(time.Now().UnixNano())
		if serv.ExpireAt < now {
			continue
		}
		if len(lb.publisher[serv.Type]) >= cap(lb.publisher[serv.Type]) {
			break
		}
		lb.publisher[serv.Type] <- serv
	}
}

func newWeiWeightedLRULoadBalancer(reg *registry) *WeightedLRULoadBalancer {
	const buffSize = 256

	lb := new(WeightedLRULoadBalancer)
	lb.reg = reg
	lb.publisher = make(map[ServiceType]chan Service)

	for _, servType := range allServiceTypes {
		lb.publisher[servType] = make(chan Service, buffSize)
	}

	return lb
}

var (
	ErrServiceInfoNotExist = fmt.Errorf("specified service not exist")

	loadBalancer LoadBalancer
	once         = sync.Once{}
)

// GetLoadBalancer returns single-instance of load balancer.
func GetLoadBalancer() LoadBalancer {
	once.Do(func() {
		loadBalancer = newWeiWeightedLRULoadBalancer(&reg)
	})
	return loadBalancer
}
