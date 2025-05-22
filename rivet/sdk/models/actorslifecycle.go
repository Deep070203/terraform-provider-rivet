package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsLifecycle struct {
    Kill_timeout types.int64 `json:"kill_timeout"`
    Durable types.bool `json:"durable"`
}

func (m *ActorsLifecycle) GetKill_timeout() types.int64 {
    if m == nil {
		return nil
	}
    return m.Kill_timeout
}
func (m *ActorsLifecycle) GetDurable() types.bool {
    if m == nil {
		return nil
	}
    return m.Durable
}