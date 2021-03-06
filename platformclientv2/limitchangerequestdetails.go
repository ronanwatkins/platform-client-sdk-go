package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Limitchangerequestdetails
type Limitchangerequestdetails struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Key - Limit key to be overridden (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
	Key *string `json:"key,omitempty"`


	// Namespace - Namespace the key belongs to (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
	Namespace *string `json:"namespace,omitempty"`


	// RequestedValue - Requested limit value for a given key
	RequestedValue *float64 `json:"requestedValue,omitempty"`


	// Description - Description of the need for the limit change request
	Description *string `json:"description,omitempty"`


	// SupportCaseUrl - The support case url created by Care
	SupportCaseUrl *string `json:"supportCaseUrl,omitempty"`


	// CreatedBy - The user who created the change request
	CreatedBy *string `json:"createdBy,omitempty"`


	// Status - Current status of the limit change request
	Status *string `json:"status,omitempty"`


	// CurrentValue - Current limit value for a given key
	CurrentValue *float64 `json:"currentValue,omitempty"`


	// DateCreated - The date of the limit change request creation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// StatusHistory - List of statuses that a limit change request has gone through
	StatusHistory *[]Statuschange `json:"statusHistory,omitempty"`


	// DateCompleted - The date of the limit change request completion (ChangeImplemented, Rejected, or RollbackImplemented. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// LastChangedBy - The user who last updated the status of the limit change request
	LastChangedBy *string `json:"lastChangedBy,omitempty"`


	// RejectReason - The reason for rejecting the limit override request
	RejectReason *string `json:"rejectReason,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Limitchangerequestdetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
