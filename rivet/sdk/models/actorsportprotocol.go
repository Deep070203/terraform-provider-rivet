package shared

import (
	"encoding/json"
	"fmt"
)

type ActorsPortProtocol string

const (

    ActorsPortProtocolhttp ActorsPortProtocol = "http"

    ActorsPortProtocolhttps ActorsPortProtocol = "https"

    ActorsPortProtocoltcp ActorsPortProtocol = "tcp"

    ActorsPortProtocolTcpTls ActorsPortProtocol = "tcp_tls"

    ActorsPortProtocoludp ActorsPortProtocol = "udp"

)

func (e ActorsPortProtocol) ToPointer() *ActorsPortProtocol {
    return &e
}

func (e *ActorsPortProtocol) UnmarshalJSON(data []byte) error {
    var v string
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch v {

        case "http":
           *e = ActorsPortProtocol(v)
           return nil

        case "https":
           *e = ActorsPortProtocol(v)
           return nil

        case "tcp":
           *e = ActorsPortProtocol(v)
           return nil

        case "tcp_tls":
           *e = ActorsPortProtocol(v)
           return nil

        case "udp":
           *e = ActorsPortProtocol(v)
           return nil

        default:
            return fmt.Errorf("invalid value for ActorsPortProtocol: %v", v)
    }
}