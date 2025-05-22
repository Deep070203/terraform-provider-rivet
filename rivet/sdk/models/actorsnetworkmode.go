package shared

import (
	"encoding/json"
	"fmt"
)

type ActorsNetworkMode string

const (

    ActorsNetworkModebridge ActorsNetworkMode = "bridge"

    ActorsNetworkModehost ActorsNetworkMode = "host"

)

func (e ActorsNetworkMode) ToPointer() *ActorsNetworkMode {
    return &e
}

func (e *ActorsNetworkMode) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch v {

        case "bridge":
           *e = ActorsNetworkMode(v)
           return nil

        case "host":
           *e = ActorsNetworkMode(v)
           return nil

        default:
            return fmt.Errorf("invalid value for ActorsNetworkMode: %v", v)
    }
}