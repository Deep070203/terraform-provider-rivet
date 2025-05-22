package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsCreateActorPortRequest struct {
    Protocol models.ActorsPortProtocol `json:"protocol"`
    Internal_port *types.int64 `json:"internal_port"`
    Routing *models.ActorsPortRouting `json:"routing"`
}

func (m *ActorsCreateActorPortRequest) GetProtocol() models.ActorsPortProtocol {
    if m == nil {
		return nil
	}
    return m.Protocol
}
func (m *ActorsCreateActorPortRequest) GetInternal_port() *types.int64 {
    if m == nil {
		return nil
	}
    return m.Internal_port
}
func (m *ActorsCreateActorPortRequest) GetRouting() *models.ActorsPortRouting {
    if m == nil {
		return nil
	}
    return m.Routing
}