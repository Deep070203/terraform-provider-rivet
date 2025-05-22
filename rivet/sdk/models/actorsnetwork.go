package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsNetwork struct {
    Mode models.ActorsNetworkMode `json:"mode"`
    Ports models.ports `json:"ports"`
}

func (m *ActorsNetwork) GetMode() models.ActorsNetworkMode {
    if m == nil {
		return nil
	}
    return m.Mode
}
func (m *ActorsNetwork) GetPorts() models.ports {
    if m == nil {
		return nil
	}
    return m.Ports
}