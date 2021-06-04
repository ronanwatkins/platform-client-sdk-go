package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultnotenote
type Bulkresponseresultnotenote struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Note `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrornote `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultnotenote) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}