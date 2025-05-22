package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsGetActorResponse struct {
    Actor models.ActorsActor `json:"actor"`
}

func (m *ActorsGetActorResponse) GetActor() models.ActorsActor {
    if m == nil {
		return nil
	}
    return m.Actor
}