package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunitsettingsresponse
type Managementunitsettingsresponse struct { 
	// Adherence - Adherence settings for this management unit
	Adherence *Adherencesettings `json:"adherence,omitempty"`


	// ShortTermForecasting - Short term forecasting settings for this management unit
	ShortTermForecasting *Shorttermforecastingsettings `json:"shortTermForecasting,omitempty"`


	// TimeOff - Time off request settings for this management unit
	TimeOff *Timeoffrequestsettings `json:"timeOff,omitempty"`


	// Scheduling - Scheduling settings for this management unit. These settings are only available if you have the permission wfm:managementUnit:view
	Scheduling *Schedulingsettingsresponse `json:"scheduling,omitempty"`


	// ShiftTrading - Shift trade settings for this management unit
	ShiftTrading *Shifttradesettings `json:"shiftTrading,omitempty"`


	// Metadata - Version info metadata for the associated management unit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Managementunitsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
