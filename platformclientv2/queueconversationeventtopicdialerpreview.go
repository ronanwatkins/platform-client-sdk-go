package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicdialerpreview
type Queueconversationeventtopicdialerpreview struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactId
	ContactId *string `json:"contactId,omitempty"`


	// ContactListId
	ContactListId *string `json:"contactListId,omitempty"`


	// CampaignId
	CampaignId *string `json:"campaignId,omitempty"`


	// PhoneNumberColumns
	PhoneNumberColumns *[]Queueconversationeventtopicphonenumbercolumn `json:"phoneNumberColumns,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicdialerpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
