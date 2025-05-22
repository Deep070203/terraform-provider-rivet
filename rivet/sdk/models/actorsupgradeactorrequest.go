package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsUpgradeActorRequest struct {
    Build types.string `json:"build"`
    Build_tags types.interface{} `json:"build_tags"`
}

func (m *ActorsUpgradeActorRequest) GetBuild() types.string {
    if m == nil {
		return nil
	}
    return m.Build
}
func (m *ActorsUpgradeActorRequest) GetBuild_tags() types.interface{} {
    if m == nil {
		return nil
	}
    return m.Build_tags
}