package registry_server

// service's infomation
type Service struct {
	Type     ServiceType
	Addr     string
	Port     int
	ExpireAt uint64
}

// Identity of a type of service.
type ServiceType uint8

const (
	RegistryService = ServiceType(0)
	IDService       = ServiceType(1) // Service supplying distributed id.
)

var allServiceTypes = []ServiceType{
	RegistryService,
	IDService,
}
