package lsd

const (
	StatsEventPublishedCmd 			= "published_cmd"
	StatsEventClientSent 			= "client_sent"
)

type StatsEvent struct {
	MessageUID 	string `json:"message_uid"`
	Event 		string `json:"event"`
	Timestamp   int64  `json:"ts"`
}
