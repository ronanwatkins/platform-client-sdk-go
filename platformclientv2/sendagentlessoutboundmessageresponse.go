package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Sendagentlessoutboundmessageresponse
type Sendagentlessoutboundmessageresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// ConversationId - The identifier of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`


	// FromAddress - The sender of the message.
	FromAddress *string `json:"fromAddress,omitempty"`


	// ToAddress - The recipient of the message.
	ToAddress *string `json:"toAddress,omitempty"`


	// MessengerType - Type of messenger.
	MessengerType *string `json:"messengerType,omitempty"`


	// TextBody - The body of the text message.
	TextBody *string `json:"textBody,omitempty"`


	// MessagingTemplate - The messaging template sent
	MessagingTemplate *Messagingtemplaterequest `json:"messagingTemplate,omitempty"`


	// Timestamp - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// User - Details of the user created the job
	User *Addressableentityref `json:"user,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessageresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
