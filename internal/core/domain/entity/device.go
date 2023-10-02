package entity

type Device struct {
	DeviceReference string `json:"device_reference" bson:"device_reference"`
	Imei            string `json:"imei" bson:"imei"`
	Type            int    `json:"type" bson:"type"`
	Brand           string `json:"brand" bson:"brand"`
	Model           string `json:"model" bson:"model"`
}
