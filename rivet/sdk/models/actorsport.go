package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsPort struct {
    Protocol models.ActorsPortProtocol `json:"protocol"`
    Internal_port *types.int64 `json:"internal_port"`
    Hostname *types.string `json:"hostname"`
    Port *types.int64 `json:"port"`
    Path *types.string `json:"path"`
    Url *types.string `json:"url"`
    Routing models.ActorsPortRouting `json:"routing"`
}

func (m *ActorsPort) GetProtocol() models.ActorsPortProtocol {
    if m == nil {
		return nil
	}
    return m.Protocol
}
func (m *ActorsPort) GetInternal_port() *types.int64 {
    if m == nil {
		return nil
	}
    return m.Internal_port
}
func (m *ActorsPort) GetHostname() *types.string {
    if m == nil {
		return nil
	}
    return m.Hostname
}
func (m *ActorsPort) GetPort() *types.int64 {
    if m == nil {
		return nil
	}
    return m.Port
}
func (m *ActorsPort) GetPath() *types.string {
    if m == nil {
		return nil
	}
    return m.Path
}
func (m *ActorsPort) GetUrl() *types.string {
    if m == nil {
		return nil
	}
    return m.Url
}
func (m *ActorsPort) GetRouting() models.ActorsPortRouting {
    if m == nil {
		return nil
	}
    return m.Routing
}