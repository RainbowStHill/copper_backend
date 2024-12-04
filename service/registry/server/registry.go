// Registry server holds infomation about other services
package registry_server

import "sync"

// service -> weight
type registry struct {
	services map[Service]int
	mtx      sync.RWMutex
}

func (r *registry) Discover(st ServiceType, lb LoadBalancer) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	for serv, w := range r.services {
		if serv.Type != st {
			continue
		}
		for ; w > 0; w-- {
			lb.Add(serv)
		}
	}
}

func (r *registry) Register(service Service, weight int) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.services[service] = weight
}

func (r *registry) Deregister(service Service) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.services[service]; !ok {
		return
	}
	delete(r.services, service)
}

var reg = registry{
	services: make(map[Service]int),
	mtx:      sync.RWMutex{},
}

func GetRegistry() *registry {
	return &reg
}
