package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationnotificationtemplateparameter - Template parameters for placeholders in template.
type Conversationnotificationtemplateparameter struct { 
	// Name - Parameter name.
	Name *string `json:"name,omitempty"`


	// Text - Parameter text value.
	Text *string `json:"text,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
