package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicdisconnectreason
type Queueconversationsocialexpressioneventtopicdisconnectreason struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Code
	Code *int `json:"code,omitempty"`


	// Phrase
	Phrase *string `json:"phrase,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicdisconnectreason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}