package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsCreateActorRequest struct {
    Region *types.string `json:"region"`
    Tags types.interface{} `json:"tags"`
    Build *types.string `json:"build"`
    Build_tags *types.interface{} `json:"build_tags"`
    Runtime *models.ActorsCreateActorRuntimeRequest `json:"runtime"`
    Network *models.ActorsCreateActorNetworkRequest `json:"network"`
    Resources *models.ActorsResources `json:"resources"`
    Lifecycle *models.ActorsLifecycle `json:"lifecycle"`
}

func (m *ActorsCreateActorRequest) GetRegion() *types.string {
    if m == nil {
		return nil
	}
    return m.Region
}
func (m *ActorsCreateActorRequest) GetTags() types.interface{} {
    if m == nil {
		return nil
	}
    return m.Tags
}
func (m *ActorsCreateActorRequest) GetBuild() *types.string {
    if m == nil {
		return nil
	}
    return m.Build
}
func (m *ActorsCreateActorRequest) GetBuild_tags() *types.interface{} {
    if m == nil {
		return nil
	}
    return m.Build_tags
}
func (m *ActorsCreateActorRequest) GetRuntime() *models.ActorsCreateActorRuntimeRequest {
    if m == nil {
		return nil
	}
    return m.Runtime
}
func (m *ActorsCreateActorRequest) GetNetwork() *models.ActorsCreateActorNetworkRequest {
    if m == nil {
		return nil
	}
    return m.Network
}
func (m *ActorsCreateActorRequest) GetResources() *models.ActorsResources {
    if m == nil {
		return nil
	}
    return m.Resources
}
func (m *ActorsCreateActorRequest) GetLifecycle() *models.ActorsLifecycle {
    if m == nil {
		return nil
	}
    return m.Lifecycle
}