package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangeoutboundsettings
type Dialeroutboundsettingsconfigchangeoutboundsettings struct { 
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


	// MaxCallsPerAgent
	MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`


	// MaxLineUtilization
	MaxLineUtilization *float32 `json:"maxLineUtilization,omitempty"`


	// AbandonSeconds
	AbandonSeconds *float32 `json:"abandonSeconds,omitempty"`


	// ComplianceAbandonRateDenominator
	ComplianceAbandonRateDenominator *string `json:"complianceAbandonRateDenominator,omitempty"`


	// AutomaticTimeZoneMapping
	AutomaticTimeZoneMapping *Dialeroutboundsettingsconfigchangeautomatictimezonemappingsettings `json:"automaticTimeZoneMapping,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeoutboundsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
