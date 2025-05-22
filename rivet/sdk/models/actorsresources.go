package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsResources struct {
    Cpu types.int64 `json:"cpu"`
    Memory types.int64 `json:"memory"`
}

func (m *ActorsResources) GetCpu() types.int64 {
    if m == nil {
		return nil
	}
    return m.Cpu
}
func (m *ActorsResources) GetMemory() types.int64 {
    if m == nil {
		return nil
	}
    return m.Memory
}