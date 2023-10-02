package shared

type Channel string 


const (
	Sms Channel = "sms"
	Email Channel = "email"
	InApp Channel = "in_app"
)