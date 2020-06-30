package main

// LogMediator ...
type LogMediator interface {

	// Populate method populate fields into the event. It should return amount of matched fields
	Populate(event LogEvent) int
}