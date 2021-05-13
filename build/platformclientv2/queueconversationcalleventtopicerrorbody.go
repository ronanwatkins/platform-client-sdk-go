package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcalleventtopicerrorbody
type Queueconversationcalleventtopicerrorbody struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// MessageWithParams
	MessageWithParams *string `json:"messageWithParams,omitempty"`


	// MessageParams
	MessageParams *map[string]string `json:"messageParams,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// Details
	Details *[]Queueconversationcalleventtopicdetail `json:"details,omitempty"`


	// Errors
	Errors *[]Queueconversationcalleventtopicerrorbody `json:"errors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicerrorbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}