package models

type OpeningHours struct {
	Monday    Events `json:"monday"`
	Tuesday   Events `json:"tuesday"`
	Wednesday Events `json:"wednesday"`
	Thursday  Events `json:"thursday"`
	Friday    Events `json:"friday"`
	Saturday  Events `json:"saturday"`
	Sunday    Events `json:"sunday"`
}

type Events []Event

type Event struct {
	Type  EventType `json:"type"`
	Value int       `json:"value"`
}

type EventType string
