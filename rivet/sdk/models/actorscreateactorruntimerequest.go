package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsCreateActorRuntimeRequest struct {
    Environment models.environment `json:"environment"`
    Network models.ActorsCreateActorRuntimeNetworkRequest `json:"network"`
}

func (m *ActorsCreateActorRuntimeRequest) GetEnvironment() models.environment {
    if m == nil {
		return nil
	}
    return m.Environment
}
func (m *ActorsCreateActorRuntimeRequest) GetNetwork() models.ActorsCreateActorRuntimeNetworkRequest {
    if m == nil {
		return nil
	}
    return m.Network
}