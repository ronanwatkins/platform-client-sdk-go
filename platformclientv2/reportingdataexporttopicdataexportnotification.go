package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingdataexporttopicdataexportnotification
type Reportingdataexporttopicdataexportnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// RunId
	RunId *string `json:"runId,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// ExportFormat
	ExportFormat *string `json:"exportFormat,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// ViewType
	ViewType *string `json:"viewType,omitempty"`


	// ExportErrorMessagesType
	ExportErrorMessagesType *string `json:"exportErrorMessagesType,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// CreatedDateTime
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`


	// ModifiedDateTime
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`


	// PercentageComplete
	PercentageComplete *float32 `json:"percentageComplete,omitempty"`


	// EmailStatuses
	EmailStatuses *map[string]string `json:"emailStatuses,omitempty"`


	// EmailErrorDescription
	EmailErrorDescription *string `json:"emailErrorDescription,omitempty"`


	// ScheduleExpression
	ScheduleExpression *string `json:"scheduleExpression,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportingdataexporttopicdataexportnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
