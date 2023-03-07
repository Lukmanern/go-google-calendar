package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	// Set up the Google Calendar API client
	ctx := context.Background()
	service, err := calendar.NewService(ctx, option.WithCredentialsFile("./credentials.json"))
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
	}

	// Create a new event
	event := &calendar.Event{
		Summary:     "My Event",
		Description: "This is an example event",
		Start: &calendar.EventDateTime{
			DateTime: time.Now().Format(time.RFC3339),
			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			DateTime: time.Now().Add(time.Hour).Format(time.RFC3339),
			TimeZone: "Asia/Jakarta",
		},
	}

	// Add the event to the calendar
	calendarID := "primary"
	event, err = service.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event: %v", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
}
