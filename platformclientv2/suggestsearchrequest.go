package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestsearchrequest
type Suggestsearchrequest struct { 
	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - Suggest query
	Query *[]Suggestsearchcriteria `json:"query,omitempty"`

}

// String returns a JSON representation of the model
func (o *Suggestsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
