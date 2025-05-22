package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsCreateActorNetworkRequest struct {
    Mode models.ActorsNetworkMode `json:"mode"`
    Ports models.ports `json:"ports"`
}

func (m *ActorsCreateActorNetworkRequest) GetMode() models.ActorsNetworkMode {
    if m == nil {
		return nil
	}
    return m.Mode
}
func (m *ActorsCreateActorNetworkRequest) GetPorts() models.ports {
    if m == nil {
		return nil
	}
    return m.Ports
}