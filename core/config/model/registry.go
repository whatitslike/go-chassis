package model

//ServiceStruct SC注册中心地址信息结构体
type ServiceStruct struct {
	Registry RegistryStruct `yaml:"registry"`
}

//RegistryStruct SC注册中心地址信息
type RegistryStruct struct {
	Disable           bool                     `yaml:"disabled"`
	Type              string                   `yaml:"type"`
	Scope             string                   `yaml:"scope"`
	AutoDiscovery     bool                     `yaml:"autodiscovery"`
	AutoIPIndex       bool                     `yaml:"autoIPIndex"`
	Address           string                   `yaml:"address"`
	RefreshInterval   string                   `yaml:"refreshInterval"`
	Watch             bool                     `yaml:"watch"`
	Tenant            string                   `yaml:"tenant"`
	AutoRegister      string                   `yaml:"register"`
	APIVersion        RegistryAPIVersionStruct `yaml:"api"`
	Registrator       RegistratorStruct        `yaml:"registrator"`
	ServiceDiscovery  ServiceDiscoveryStruct   `yaml:"serviceDiscovery"`
	ContractDiscovery ContractDiscoveryStruct  `yaml:"contractDiscovery"`
	HealthCheck       bool                     `yaml:"healthCheck"`
}

//RegistratorStruct service registry config struct
type RegistratorStruct struct {
	Disable         bool                     `yaml:"disabled"`
	Type            string                   `yaml:"type"`
	Scope           string                   `yaml:"scope"`
	Address         string                   `yaml:"address"`
	RefreshInterval string                   `yaml:"refreshInterval"`
	Tenant          string                   `yaml:"tenant"`
	AutoRegister    string                   `yaml:"register"`
	APIVersion      RegistryAPIVersionStruct `yaml:"api"`
}

//ServiceDiscoveryStruct service discovery config struct
type ServiceDiscoveryStruct struct {
	Disable         bool                     `yaml:"disabled"`
	Type            string                   `yaml:"type"`
	AutoDiscovery   bool                     `yaml:"autodiscovery"`
	AutoIPIndex     bool                     `yaml:"autoIPIndex"`
	Address         string                   `yaml:"address"`
	RefreshInterval string                   `yaml:"refreshInterval"`
	Watch           bool                     `yaml:"watch"`
	Tenant          string                   `yaml:"tenant"`
	APIVersion      RegistryAPIVersionStruct `yaml:"api"`
	HealthCheck     bool                     `yaml:"healthCheck"`
}

//ContractDiscoveryStruct contract discovery config struct
type ContractDiscoveryStruct struct {
	Disable         bool                     `yaml:"disabled"`
	Type            string                   `yaml:"type"`
	Address         string                   `yaml:"address"`
	RefreshInterval string                   `yaml:"refreshInterval"`
	Tenant          string                   `yaml:"tenant"`
	APIVersion      RegistryAPIVersionStruct `yaml:"api"`
}

// RegistryAPIVersionStruct registry api version structure
type RegistryAPIVersionStruct struct {
	Version string `yaml:"version"`
}
