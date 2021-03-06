package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequenceconfigchangecampaignsequence
type Dialersequenceconfigchangecampaignsequence struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// Campaigns
	Campaigns *[]Dialersequenceconfigchangeurireference `json:"campaigns,omitempty"`


	// CurrentCampaign
	CurrentCampaign *int `json:"currentCampaign,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// StopMessage
	StopMessage *string `json:"stopMessage,omitempty"`


	// Repeat
	Repeat *bool `json:"repeat,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialersequenceconfigchangecampaignsequence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
