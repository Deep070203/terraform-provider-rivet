package shared

import (
	"encoding/json"
	"fmt"
)



type ErrorBody struct {
    Code types.string `json:"code"`
    Message types.string `json:"message"`
    Ray_id types.string `json:"ray_id"`
    Documentation *types.string `json:"documentation"`
    Metadata *models.ErrorMetadata `json:"metadata"`
}

func (m *ErrorBody) GetCode() types.string {
    if m == nil {
		return nil
	}
    return m.Code
}
func (m *ErrorBody) GetMessage() types.string {
    if m == nil {
		return nil
	}
    return m.Message
}
func (m *ErrorBody) GetRay_id() types.string {
    if m == nil {
		return nil
	}
    return m.Ray_id
}
func (m *ErrorBody) GetDocumentation() *types.string {
    if m == nil {
		return nil
	}
    return m.Documentation
}
func (m *ErrorBody) GetMetadata() *models.ErrorMetadata {
    if m == nil {
		return nil
	}
    return m.Metadata
}