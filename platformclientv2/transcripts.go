package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Transcripts
type Transcripts struct { 
	// ExactMatch - List of transcript contents which needs to satisfy exact match criteria
	ExactMatch *[]string `json:"exactMatch,omitempty"`


	// Contains - List of transcript contents which needs to satisfy contains criteria
	Contains *[]string `json:"contains,omitempty"`


	// DoesNotContain - List of transcript contents which needs to satisfy does not contain criteria
	DoesNotContain *[]string `json:"doesNotContain,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcripts) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
