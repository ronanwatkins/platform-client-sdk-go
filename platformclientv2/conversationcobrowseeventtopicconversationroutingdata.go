package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopicconversationroutingdata
type Conversationcobrowseeventtopicconversationroutingdata struct { 
	// Queue
	Queue *Conversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Language
	Language *Conversationcobrowseeventtopicurireference `json:"language,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// Skills
	Skills *[]Conversationcobrowseeventtopicurireference `json:"skills,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Conversationcobrowseeventtopicscoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicconversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
