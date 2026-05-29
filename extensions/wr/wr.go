package wr

import (
	ical "github.com/arran4/golang-ical"
)

// PropertyXWrTimezone is the X-WR-TIMEZONE property
const PropertyXWrTimezone = "X-WR-TIMEZONE"

// PropertyXWrCalName is the X-WR-CALNAME property
const PropertyXWrCalName = "X-WR-CALNAME"

// PropertyXWrCalDesc is the X-WR-CALDESC property
const PropertyXWrCalDesc = "X-WR-CALDESC"

// PropertyXWrRelCalId is the X-WR-RELCALID property
const PropertyXWrRelCalId = "X-WR-RELCALID"

// SetProperty allows extending the properties easily
func SetProperty(cal *ical.Calendar, property string, value string, params ...ical.PropertyParameter) {
	cal.SetProperty(ical.Property(property), value, params...)
}

// SetTimezone sets the X-WR-TIMEZONE property
func SetTimezone(cal *ical.Calendar, tz string, params ...ical.PropertyParameter) {
	cal.SetXWRTimezone(tz, params...)
}

// SetCalName sets the X-WR-CALNAME property
func SetCalName(cal *ical.Calendar, name string, params ...ical.PropertyParameter) {
	cal.SetXWRCalName(name, params...)
}

// SetCalDesc sets the X-WR-CALDESC property
func SetCalDesc(cal *ical.Calendar, desc string, params ...ical.PropertyParameter) {
	cal.SetXWRCalDesc(desc, params...)
}

// SetRelCalId sets the X-WR-RELCALID property
func SetRelCalId(cal *ical.Calendar, id string, params ...ical.PropertyParameter) {
	cal.SetXWRCalID(id, params...)
}
