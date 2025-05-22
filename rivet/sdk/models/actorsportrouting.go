package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsPortRouting struct {
    Guard models.ActorsGuardRouting `json:"guard"`
    Host models.ActorsHostRouting `json:"host"`
}

func (m *ActorsPortRouting) GetGuard() models.ActorsGuardRouting {
    if m == nil {
		return nil
	}
    return m.Guard
}
func (m *ActorsPortRouting) GetHost() models.ActorsHostRouting {
    if m == nil {
		return nil
	}
    return m.Host
}