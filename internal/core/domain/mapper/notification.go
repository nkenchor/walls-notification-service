package mapper

import (
	"time"
	"walls-notification-service/internal/core/domain/dto"
	"walls-notification-service/internal/core/domain/entity"
	configuration "walls-notification-service/internal/core/helper/configuration-helper"

	"github.com/google/uuid"
)

func MapCreateDto(notificationDto dto.CreateNotification) entity.Notification {

	notificationDto.MessageBody = "From: WallsPay" + "\r\n" +
		"To: " + notificationDto.Contact + "\r\n" +
		"Subject: " + notificationDto.Subject + "\r\n" +
		"\r\n" +
		notificationDto.MessageBody

	notificationMap := entity.Notification{
		Reference:       uuid.New().String(),
		UserReference:   notificationDto.UserReference,
		DeviceReference: notificationDto.DeviceReference,
		From:            "WallsPay",
		Contact:         notificationDto.Contact,
		Channel:         notificationDto.Channel,
		Type:            notificationDto.Type,
		Subject:         notificationDto.Subject,
		MessageBody:     notificationDto.MessageBody,
		NotifiedBy:      configuration.ServiceConfiguration.ServiceName,
		NotifyOn:        time.Now().Format(time.RFC3339),
		NotifiedOn:      time.Now().Format(time.RFC3339),
	}
	return notificationMap
}
