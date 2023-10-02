package entity

import "walls-notification-service/internal/core/domain/shared"



type Notification struct {
	Reference       string                    `json:"reference" bson:"reference"`
	UserReference   string                    `json:"user_reference" bson:"user_reference"`
	DeviceReference string                    `json:"device_reference" bson:"device_reference"`
	From            string                    `json:"from" bson:"from"`
	Contact         string                    `json:"contact" bson:"contact"`
	Channel         shared.Channel           `json:"channel" bson:"channel"`
	Type            string  `json:"notification_type" bson:"notification_type"`
	Subject         string                    `json:"subject" bson:"subject"`
	MessageBody     string                    `json:"message_body" bson:"message_body"`
	NotifiedBy      string                    `json:"notified_by" bson:"notified_by"`
	NotifyOn        string                    `json:"notify_on" bson:"notify_on"`
	NotifiedOn      string                    `json:"notified_on" bson:"notified_on"`
	Status          string `json:"status" bson:"status"`
}
