package repository

import (
	"context"
	"walls-notification-service/internal/core/domain/entity"

	logger "walls-notification-service/internal/core/helper/log-helper"
	ports "walls-notification-service/internal/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationInfra struct {
	Collection *mongo.Collection
}

func NewNotification(Collection *mongo.Collection) *NotificationInfra {
	return &NotificationInfra{Collection}
}

// UserRepo implements the repository.UserRepository interface
var _ ports.NotificationRepository = &NotificationInfra{}

func (r *NotificationInfra) CreateNotification(notification entity.Notification) (interface{}, error) {
	logger.LogEvent("INFO", "Persisting notification with reference: "+notification.Reference)

	_, err := r.Collection.InsertOne(context.TODO(), notification)
	if err != nil {
		return nil, err
	}

	logger.LogEvent("INFO", "Persisting notification with reference: "+notification.Reference+" completed successfully...")
	return notification.Reference, nil
}
func (r *NotificationInfra) GetNotificationByDeviceReference(device_reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving last notification with device reference: "+device_reference)
	notification := entity.Notification{}
	filter := bson.M{"device_reference": device_reference}

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&notification)
	if err != nil || notification == (entity.Notification{}) {
		return nil, err
	}

	logger.LogEvent("INFO", "Retrieving notifications with device reference: "+device_reference+" completed successfully. ")
	return notification, nil

}

func (r *NotificationInfra) GetNotificationByUserReference(user_reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving last notification with user reference: "+user_reference)
	notification := entity.Notification{}
	filter := bson.M{"user_reference": user_reference}

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&notification)
	if err != nil || notification == (entity.Notification{}) {
		return nil, err
	}

	logger.LogEvent("INFO", "Retrieving notifications with notification reference: "+user_reference+" completed successfully. ")
	return notification, nil

}

func (r *NotificationInfra) GetNotificationByReference(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving last notification with notification reference: "+reference)
	notification := entity.Notification{}

	filter := bson.M{"reference": reference}

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&notification)
	if err != nil || notification == (entity.Notification{}) {
		return nil, err
	}
	logger.LogEvent("INFO", "Retrieving notification with reference: "+reference+" completed successfully. ")
	return notification, nil

}
