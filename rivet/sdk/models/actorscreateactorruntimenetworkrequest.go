package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsCreateActorRuntimeNetworkRequest struct {
    Endpoint_type models.ActorsEndpointType `json:"endpoint_type"`
}

func (m *ActorsCreateActorRuntimeNetworkRequest) GetEndpoint_type() models.ActorsEndpointType {
    if m == nil {
		return nil
	}
    return m.Endpoint_type
}