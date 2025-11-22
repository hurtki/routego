package routego

import "regexp"

type RoutegoConfig struct {
	Port string // :port
}

// NewAppConfig creates a new RoutegoConfig entity with validation
// all RoutegoConfig fields and how they should look see in AppConfig structure
func NewRoutegoConfig(port string) RoutegoConfig {
	port_reg_exp := regexp.MustCompile(`^:\d{1,5}$`)
	if !port_reg_exp.MatchString(port) {
		panic("wrong port specified")
	}
	return RoutegoConfig{
		Port: port,
	}
}
