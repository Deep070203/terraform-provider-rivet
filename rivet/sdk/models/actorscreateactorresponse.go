package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsCreateActorResponse struct {
    Actor models.ActorsActor `json:"actor"`
}

func (m *ActorsCreateActorResponse) GetActor() models.ActorsActor {
    if m == nil {
		return nil
	}
    return m.Actor
}