package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsRuntime struct {
    Build types.string `json:"build"`
    Arguments *models.[]string `json:"arguments"`
    Environment *models.environment `json:"environment"`
}

func (m *ActorsRuntime) GetBuild() types.string {
    if m == nil {
		return nil
	}
    return m.Build
}
func (m *ActorsRuntime) GetArguments() *models.[]string {
    if m == nil {
		return nil
	}
    return m.Arguments
}
func (m *ActorsRuntime) GetEnvironment() *models.environment {
    if m == nil {
		return nil
	}
    return m.Environment
}