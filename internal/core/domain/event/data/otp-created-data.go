package event

import (
	"walls-notification-service/internal/core/domain/entity"
	"walls-notification-service/internal/core/domain/shared"
)


type OtpCreatedEventData struct {
	Reference     string               `json:"reference"`
	UserReference string			   `json:"user_reference"`				
	Otp           string			   `json:"otp"`
	Contact       string			   `json:"contact"`
	Channel       shared.Channel      `json:"channel"`
	Device        entity.Device		   `json:"device"`
	Expired       bool			       `json:"expired"`
	ExpiresAt     string			   `json:"expires_at"`
	CreatedOn     string               `json:"created_on"`
	UpdatedOn     string               `json:"updated_on"`
}
