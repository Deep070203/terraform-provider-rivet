package shared

import (
	"encoding/json"
	"fmt"
)

type ActorsEndpointType string

const (

    ActorsEndpointTypehostname ActorsEndpointType = "hostname"

    ActorsEndpointTypepath ActorsEndpointType = "path"

)

func (e ActorsEndpointType) ToPointer() *ActorsEndpointType {
    return &e
}

func (e *ActorsEndpointType) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch v {

        case "hostname":
           *e = ActorsEndpointType(v)
           return nil

        case "path":
           *e = ActorsEndpointType(v)
           return nil

        default:
            return fmt.Errorf("invalid value for ActorsEndpointType: %v", v)
    }
}