package shared

import (
	"encoding/json"
	"fmt"
)



type Pagination struct {
    Cursor types.string `json:"cursor"`
}

func (m *Pagination) GetCursor() types.string {
    if m == nil {
		return nil
	}
    return m.Cursor
}