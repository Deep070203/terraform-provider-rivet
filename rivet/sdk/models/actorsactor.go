package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsActor struct {
    Id types.string `json:"id"`
    Region types.string `json:"region"`
    Tags types.interface{} `json:"tags"`
    Runtime models.ActorsRuntime `json:"runtime"`
    Network models.ActorsNetwork `json:"network"`
    Resources models.ActorsResources `json:"resources"`
    Lifecycle models.ActorsLifecycle `json:"lifecycle"`
    Created_at models.Timestamp `json:"created_at"`
    Started_at *models.Timestamp `json:"started_at"`
    Destroyed_at *models.Timestamp `json:"destroyed_at"`
}

func (m *ActorsActor) GetId() types.string {
    if m == nil {
		return nil
	}
    return m.Id
}
func (m *ActorsActor) GetRegion() types.string {
    if m == nil {
		return nil
	}
    return m.Region
}
func (m *ActorsActor) GetTags() types.interface{} {
    if m == nil {
		return nil
	}
    return m.Tags
}
func (m *ActorsActor) GetRuntime() models.ActorsRuntime {
    if m == nil {
		return nil
	}
    return m.Runtime
}
func (m *ActorsActor) GetNetwork() models.ActorsNetwork {
    if m == nil {
		return nil
	}
    return m.Network
}
func (m *ActorsActor) GetResources() models.ActorsResources {
    if m == nil {
		return nil
	}
    return m.Resources
}
func (m *ActorsActor) GetLifecycle() models.ActorsLifecycle {
    if m == nil {
		return nil
	}
    return m.Lifecycle
}
func (m *ActorsActor) GetCreated_at() models.Timestamp {
    if m == nil {
		return nil
	}
    return m.Created_at
}
func (m *ActorsActor) GetStarted_at() *models.Timestamp {
    if m == nil {
		return nil
	}
    return m.Started_at
}
func (m *ActorsActor) GetDestroyed_at() *models.Timestamp {
    if m == nil {
		return nil
	}
    return m.Destroyed_at
}