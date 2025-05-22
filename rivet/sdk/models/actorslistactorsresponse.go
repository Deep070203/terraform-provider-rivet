package shared

import (
	"encoding/json"
	"fmt"
)



type ActorsListActorsResponse struct {
    Actors models.[]ActorsActor `json:"actors"`
    Pagination models.Pagination `json:"pagination"`
}

func (m *ActorsListActorsResponse) GetActors() models.[]ActorsActor {
    if m == nil {
		return nil
	}
    return m.Actors
}
func (m *ActorsListActorsResponse) GetPagination() models.Pagination {
    if m == nil {
		return nil
	}
    return m.Pagination
}