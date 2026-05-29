package wg

import (
	ical "github.com/arran4/golang-ical"
)

// SetProperty allows extending the properties easily
func SetProperty(cal *ical.Calendar, property string, value string, params ...ical.PropertyParameter) {
	cal.SetProperty(ical.Property(property), value, params...)
}

// SetTimezone sets the X-WG-TIMEZONE property
func SetTimezone(cal *ical.Calendar, tz string, params ...ical.PropertyParameter) {
	SetProperty(cal, "X-WG-TIMEZONE", tz, params...)
}

// SetCalName sets the X-WG-CALNAME property
func SetCalName(cal *ical.Calendar, name string, params ...ical.PropertyParameter) {
	SetProperty(cal, "X-WG-CALNAME", name, params...)
}

// SetCalDesc sets the X-WG-CALDESC property
func SetCalDesc(cal *ical.Calendar, desc string, params ...ical.PropertyParameter) {
	SetProperty(cal, "X-WG-CALDESC", desc, params...)
}

// SetRelCalId sets the X-WG-RELCALID property
func SetRelCalId(cal *ical.Calendar, id string, params ...ical.PropertyParameter) {
	SetProperty(cal, "X-WG-RELCALID", id, params...)
}
