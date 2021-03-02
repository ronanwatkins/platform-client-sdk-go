package platformclientv2
import (
	"encoding/json"
)

// Predictionresults
type Predictionresults struct { 
	// Intent - Indicates the media type scope of this estimated wait time
	Intent *string `json:"intent,omitempty"`


	// Formula - Indicates the estimated wait time Formula
	Formula *string `json:"formula,omitempty"`


	// EstimatedWaitTimeSeconds - Estimated wait time in seconds
	EstimatedWaitTimeSeconds *int `json:"estimatedWaitTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Predictionresults) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}