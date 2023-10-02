package handlers

import (
	"encoding/json"
	"fmt"
	"walls-notification-service/internal/core/domain/dto"
	events "walls-notification-service/internal/core/domain/event/data"
	eto "walls-notification-service/internal/core/helper/event-helper/eto"
	helper "walls-notification-service/internal/core/helper/message-helper"
	"walls-notification-service/internal/core/services"
)

// Event handler function
func OtpCreatedEventHandler(event interface{}) {

	var otpCreatedEvent eto.Event
	err := convertEvent(event, &otpCreatedEvent)
	if err != nil {
		fmt.Println("Error converting event to otpCreatedEvent:", err)
		return
	}

	var otpCreatedEventData events.OtpCreatedEventData
	err = convertEvent(otpCreatedEvent.EventData, &otpCreatedEventData)
	if err != nil {
		fmt.Println("Error converting event data to otpCreatedData:", err)
		return
	}

	createNotificationDto := dto.CreateNotification{
		UserReference:   otpCreatedEventData.UserReference,
		DeviceReference: otpCreatedEventData.Device.DeviceReference,
		Contact:         otpCreatedEventData.Contact,
		Channel:         otpCreatedEventData.Channel,
		Subject:         "Walls OTP",
		MessageBody:     helper.OtpMessage + otpCreatedEventData.Otp,
	}

	// Create an instance of the NotificationService

	services.NotificationService.CreateNotification(createNotificationDto)

}

func convertEvent(event interface{}, outputEvent interface{}) error {
	// Convert interface{} to byte array
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Deserialize JSON to outputEvent
	err = json.Unmarshal(jsonBytes, outputEvent)
	if err != nil {
		return err
	}

	return nil
}
