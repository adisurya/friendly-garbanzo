package responses

import "github.com/adisurya/friendly-garbanzo/database/events"

type EventList struct {
	Events []events.EventList `json:"events"`
}
