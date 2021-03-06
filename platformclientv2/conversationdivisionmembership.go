package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdivisionmembership
type Conversationdivisionmembership struct { 
	// Division - A division the conversation belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// Entities - The entities on the conversation within the division. These are the users, queues, work flows, etc. that can be on conversations and and be assigned to different divisions.
	Entities *[]Domainentityref `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdivisionmembership) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
