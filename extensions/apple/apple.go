//go:build !noapple
// +build !noapple

package apple

import (
	ical "github.com/arran4/golang-ical"
)

// PropertyXAppleCalendarColor is the X-APPLE-CALENDAR-COLOR property
const PropertyXAppleCalendarColor = "X-APPLE-CALENDAR-COLOR"

// PropertyXAppleRegion is the X-APPLE-REGION property
const PropertyXAppleRegion = "X-APPLE-REGION"

// ComponentPropertyXAppleStructuredLocation is the X-APPLE-STRUCTURED-LOCATION component property
const ComponentPropertyXAppleStructuredLocation = "X-APPLE-STRUCTURED-LOCATION"

// ComponentPropertyXAppleTravelDuration is the X-APPLE-TRAVEL-DURATION component property
const ComponentPropertyXAppleTravelDuration = "X-APPLE-TRAVEL-DURATION"

type propertySetter interface {
	SetProperty(property ical.ComponentProperty, value string, props ...ical.PropertyParameter)
}

// SetProperty allows extending the properties easily
func SetProperty(cal *ical.Calendar, property string, value string, params ...ical.PropertyParameter) {
	cal.SetProperty(ical.Property(property), value, params...)
}

func SetComponentProperty(c ical.Component, property string, value string, params ...ical.PropertyParameter) {
	if ps, ok := c.(propertySetter); ok {
		ps.SetProperty(ical.ComponentProperty(property), value, params...)
	}
}

// Calendar properties

// SetCalendarColor sets the X-APPLE-CALENDAR-COLOR property for the calendar
func SetCalendarColor(cal *ical.Calendar, color string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyXAppleCalendarColor, color, params...)
}

// SetRegion sets the X-APPLE-REGION property for the calendar
func SetRegion(cal *ical.Calendar, region string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyXAppleRegion, region, params...)
}

// Component properties

// SetStructuredLocation sets the X-APPLE-STRUCTURED-LOCATION property for a component
func SetStructuredLocation(c ical.Component, location string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyXAppleStructuredLocation, location, params...)
}

// SetTravelDuration sets the X-APPLE-TRAVEL-DURATION property for a component
func SetTravelDuration(c ical.Component, duration string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyXAppleTravelDuration, duration, params...)
}
