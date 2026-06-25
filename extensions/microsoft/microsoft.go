package microsoft

import (
	ical "github.com/arran4/golang-ical"
)

// PropertyCalend is the X-CALEND property
const PropertyCalend = "X-CALEND"

// PropertyCalstart is the X-CALSTART property
const PropertyCalstart = "X-CALSTART"

// PropertyClipend is the X-CLIPEND property
const PropertyClipend = "X-CLIPEND"

// PropertyClipstart is the X-CLIPSTART property
const PropertyClipstart = "X-CLIPSTART"

// PropertyMicrosoftCalscale is the X-MICROSOFT-CALSCALE property
const PropertyMicrosoftCalscale = "X-MICROSOFT-CALSCALE"

// PropertyMsOlkForceinspectoropen is the X-MS-OLK-FORCEINSPECTOROPEN property
const PropertyMsOlkForceinspectoropen = "X-MS-OLK-FORCEINSPECTOROPEN"

// PropertyMsWkhrdays is the X-MS-WKHRDAYS property
const PropertyMsWkhrdays = "X-MS-WKHRDAYS"

// PropertyMsWkhrend is the X-MS-WKHREND property
const PropertyMsWkhrend = "X-MS-WKHREND"

// PropertyMsWkhrstart is the X-MS-WKHRSTART property
const PropertyMsWkhrstart = "X-MS-WKHRSTART"

// PropertyOwner is the X-OWNER property
const PropertyOwner = "X-OWNER"

// PropertyPrimaryCalendar is the X-PRIMARY-CALENDAR property
const PropertyPrimaryCalendar = "X-PRIMARY-CALENDAR"

// PropertyPublishedTtl is the X-PUBLISHED-TTL property
const PropertyPublishedTtl = "X-PUBLISHED-TTL"

// PropertyCaldesc is the X-WR-CALDESC property
const PropertyCaldesc = "X-WR-CALDESC"

// PropertyCalname is the X-WR-CALNAME property
const PropertyCalname = "X-WR-CALNAME"

// PropertyRelcalid is the X-WR-RELCALID property
const PropertyRelcalid = "X-WR-RELCALID"

// ComponentPropertyAltDesc is the X-ALT-DESC component property
const ComponentPropertyAltDesc = "X-ALT-DESC"

// ComponentPropertyMicrosoftCdoAlldayevent is the X-MICROSOFT-CDO-ALLDAYEVENT component property
const ComponentPropertyMicrosoftCdoAlldayevent = "X-MICROSOFT-CDO-ALLDAYEVENT"

// ComponentPropertyMicrosoftCdoApptSequence is the X-MICROSOFT-CDO-APPT-SEQUENCE component property
const ComponentPropertyMicrosoftCdoApptSequence = "X-MICROSOFT-CDO-APPT-SEQUENCE"

// ComponentPropertyMicrosoftCdoAttendeeCriticalChange is the X-MICROSOFT-CDO-ATTENDEE-CRITICAL-CHANGE component property
const ComponentPropertyMicrosoftCdoAttendeeCriticalChange = "X-MICROSOFT-CDO-ATTENDEE-CRITICAL-CHANGE"

// ComponentPropertyMicrosoftCdoBusystatus is the X-MICROSOFT-CDO-BUSYSTATUS component property
const ComponentPropertyMicrosoftCdoBusystatus = "X-MICROSOFT-CDO-BUSYSTATUS"

// ComponentPropertyMicrosoftCdoImportance is the X-MICROSOFT-CDO-IMPORTANCE component property
const ComponentPropertyMicrosoftCdoImportance = "X-MICROSOFT-CDO-IMPORTANCE"

// ComponentPropertyMicrosoftCdoInsttype is the X-MICROSOFT-CDO-INSTTYPE component property
const ComponentPropertyMicrosoftCdoInsttype = "X-MICROSOFT-CDO-INSTTYPE"

// ComponentPropertyMicrosoftCdoIntendedstatus is the X-MICROSOFT-CDO-INTENDEDSTATUS component property
const ComponentPropertyMicrosoftCdoIntendedstatus = "X-MICROSOFT-CDO-INTENDEDSTATUS"

// ComponentPropertyMicrosoftCdoOwnerapptid is the X-MICROSOFT-CDO-OWNERAPPTID component property
const ComponentPropertyMicrosoftCdoOwnerapptid = "X-MICROSOFT-CDO-OWNERAPPTID"

// ComponentPropertyMicrosoftCdoOwnerCriticalChange is the X-MICROSOFT-CDO-OWNER-CRITICAL-CHANGE component property
const ComponentPropertyMicrosoftCdoOwnerCriticalChange = "X-MICROSOFT-CDO-OWNER-CRITICAL-CHANGE"

// ComponentPropertyMicrosoftCdoReplytime is the X-MICROSOFT-CDO-REPLYTIME component property
const ComponentPropertyMicrosoftCdoReplytime = "X-MICROSOFT-CDO-REPLYTIME"

// ComponentPropertyMicrosoftDisallowCounter is the X-MICROSOFT-DISALLOW-COUNTER component property
const ComponentPropertyMicrosoftDisallowCounter = "X-MICROSOFT-DISALLOW-COUNTER"

// ComponentPropertyMicrosoftExdate is the X-MICROSOFT-EXDATE component property
const ComponentPropertyMicrosoftExdate = "X-MICROSOFT-EXDATE"

// ComponentPropertyMicrosoftIsdraft is the X-MICROSOFT-ISDRAFT component property
const ComponentPropertyMicrosoftIsdraft = "X-MICROSOFT-ISDRAFT"

// ComponentPropertyMicrosoftMsncalendarAlldayevent is the X-MICROSOFT-MSNCALENDAR-ALLDAYEVENT component property
const ComponentPropertyMicrosoftMsncalendarAlldayevent = "X-MICROSOFT-MSNCALENDAR-ALLDAYEVENT"

// ComponentPropertyMicrosoftMsncalendarBusystatus is the X-MICROSOFT-MSNCALENDAR-BUSYSTATUS component property
const ComponentPropertyMicrosoftMsncalendarBusystatus = "X-MICROSOFT-MSNCALENDAR-BUSYSTATUS"

// ComponentPropertyMicrosoftMsncalendarImportance is the X-MICROSOFT-MSNCALENDAR-IMPORTANCE component property
const ComponentPropertyMicrosoftMsncalendarImportance = "X-MICROSOFT-MSNCALENDAR-IMPORTANCE"

// ComponentPropertyMicrosoftMsncalendarIntendedstatus is the X-MICROSOFT-MSNCALENDAR-INTENDEDSTATUS component property
const ComponentPropertyMicrosoftMsncalendarIntendedstatus = "X-MICROSOFT-MSNCALENDAR-INTENDEDSTATUS"

// ComponentPropertyMicrosoftRrule is the X-MICROSOFT-RRULE component property
const ComponentPropertyMicrosoftRrule = "X-MICROSOFT-RRULE"

// ComponentPropertyMsOlkAllowexterncheck is the X-MS-OLK-ALLOWEXTERNCHECK component property
const ComponentPropertyMsOlkAllowexterncheck = "X-MS-OLK-ALLOWEXTERNCHECK"

// ComponentPropertyMsOlkApptlastsequence is the X-MS-OLK-APPTLASTSEQUENCE component property
const ComponentPropertyMsOlkApptlastsequence = "X-MS-OLK-APPTLASTSEQUENCE"

// ComponentPropertyMsOlkApptseqtime is the X-MS-OLK-APPTSEQTIME component property
const ComponentPropertyMsOlkApptseqtime = "X-MS-OLK-APPTSEQTIME"

// ComponentPropertyMsOlkAutofilllocation is the X-MS-OLK-AUTOFILLLOCATION component property
const ComponentPropertyMsOlkAutofilllocation = "X-MS-OLK-AUTOFILLLOCATION"

// ComponentPropertyMsOlkAutostartcheck is the X-MS-OLK-AUTOSTARTCHECK component property
const ComponentPropertyMsOlkAutostartcheck = "X-MS-OLK-AUTOSTARTCHECK"

// ComponentPropertyMsOlkCollaboratedoc is the X-MS-OLK-COLLABORATEDOC component property
const ComponentPropertyMsOlkCollaboratedoc = "X-MS-OLK-COLLABORATEDOC"

// ComponentPropertyMsOlkConfcheck is the X-MS-OLK-CONFCHECK component property
const ComponentPropertyMsOlkConfcheck = "X-MS-OLK-CONFCHECK"

// ComponentPropertyMsOlkConftype is the X-MS-OLK-CONFTYPE component property
const ComponentPropertyMsOlkConftype = "X-MS-OLK-CONFTYPE"

// ComponentPropertyMsOlkDirectory is the X-MS-OLK-DIRECTORY component property
const ComponentPropertyMsOlkDirectory = "X-MS-OLK-DIRECTORY"

// ComponentPropertyMsOlkMwsurl is the X-MS-OLK-MWSURL component property
const ComponentPropertyMsOlkMwsurl = "X-MS-OLK-MWSURL"

// ComponentPropertyMsOlkNetshowurl is the X-MS-OLK-NETSHOWURL component property
const ComponentPropertyMsOlkNetshowurl = "X-MS-OLK-NETSHOWURL"

// ComponentPropertyMsOlkOnlinepassword is the X-MS-OLK-ONLINEPASSWORD component property
const ComponentPropertyMsOlkOnlinepassword = "X-MS-OLK-ONLINEPASSWORD"

// ComponentPropertyMsOlkOrgalias is the X-MS-OLK-ORGALIAS component property
const ComponentPropertyMsOlkOrgalias = "X-MS-OLK-ORGALIAS"

// ComponentPropertyMsOlkOriginalend is the X-MS-OLK-ORIGINALEND component property
const ComponentPropertyMsOlkOriginalend = "X-MS-OLK-ORIGINALEND"

// ComponentPropertyMsOlkOriginalstart is the X-MS-OLK-ORIGINALSTART component property
const ComponentPropertyMsOlkOriginalstart = "X-MS-OLK-ORIGINALSTART"

// ComponentPropertyMsOlkSender is the X-MS-OLK-SENDER component property
const ComponentPropertyMsOlkSender = "X-MS-OLK-SENDER"

type propertySetter interface {
	SetProperty(property ical.ComponentProperty, value string, props ...ical.PropertyParameter)
}

// SetProperty allows extending the properties easily
func SetProperty(cal *ical.Calendar, property string, value string, params ...ical.PropertyParameter) {
	if cal == nil {
		return
	}
	// Fallback to directly modify CalendarProperties if no public SetProperty exists
	found := false
	for i := range cal.CalendarProperties {
		if cal.CalendarProperties[i].IANAToken == property {
			cal.CalendarProperties[i].Value = value
			cal.CalendarProperties[i].ICalParameters = map[string][]string{}
			for _, p := range params {
				k, v := p.KeyValue()
				cal.CalendarProperties[i].ICalParameters[k] = v
			}
			found = true
			break
		}
	}
	if !found {
		r := ical.CalendarProperty{
			BaseProperty: ical.BaseProperty{
				IANAToken:      property,
				Value:          value,
				ICalParameters: map[string][]string{},
			},
		}
		for _, p := range params {
			k, v := p.KeyValue()
			r.ICalParameters[k] = v
		}
		cal.CalendarProperties = append(cal.CalendarProperties, r)
	}
}

func SetComponentProperty(c ical.Component, property string, value string, params ...ical.PropertyParameter) {
	if c == nil {
		return
	}
	if ps, ok := c.(propertySetter); ok {
		ps.SetProperty(ical.ComponentProperty(property), value, params...)
	}
}

// Calendar properties
// SetCalend sets the X-CALEND property for the calendar
func SetCalend(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyCalend, value, params...)
}

// SetCalstart sets the X-CALSTART property for the calendar
func SetCalstart(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyCalstart, value, params...)
}

// SetClipend sets the X-CLIPEND property for the calendar
func SetClipend(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyClipend, value, params...)
}

// SetClipstart sets the X-CLIPSTART property for the calendar
func SetClipstart(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyClipstart, value, params...)
}

// SetMicrosoftCalscale sets the X-MICROSOFT-CALSCALE property for the calendar
func SetMicrosoftCalscale(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyMicrosoftCalscale, value, params...)
}

// SetMsOlkForceinspectoropen sets the X-MS-OLK-FORCEINSPECTOROPEN property for the calendar
func SetMsOlkForceinspectoropen(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyMsOlkForceinspectoropen, value, params...)
}

// SetMsWkhrdays sets the X-MS-WKHRDAYS property for the calendar
func SetMsWkhrdays(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyMsWkhrdays, value, params...)
}

// SetMsWkhrend sets the X-MS-WKHREND property for the calendar
func SetMsWkhrend(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyMsWkhrend, value, params...)
}

// SetMsWkhrstart sets the X-MS-WKHRSTART property for the calendar
func SetMsWkhrstart(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyMsWkhrstart, value, params...)
}

// SetOwner sets the X-OWNER property for the calendar
func SetOwner(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyOwner, value, params...)
}

// SetPrimaryCalendar sets the X-PRIMARY-CALENDAR property for the calendar
func SetPrimaryCalendar(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyPrimaryCalendar, value, params...)
}

// SetPublishedTtl sets the X-PUBLISHED-TTL property for the calendar
func SetPublishedTtl(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyPublishedTtl, value, params...)
}

// SetCaldesc sets the X-WR-CALDESC property for the calendar
func SetCaldesc(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyCaldesc, value, params...)
}

// SetCalname sets the X-WR-CALNAME property for the calendar
func SetCalname(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyCalname, value, params...)
}

// SetRelcalid sets the X-WR-RELCALID property for the calendar
func SetRelcalid(cal *ical.Calendar, value string, params ...ical.PropertyParameter) {
	SetProperty(cal, PropertyRelcalid, value, params...)
}

// Component properties
// SetAltDesc sets the X-ALT-DESC property for a component
func SetAltDesc(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyAltDesc, value, params...)
}

// SetMicrosoftCdoAlldayevent sets the X-MICROSOFT-CDO-ALLDAYEVENT property for a component
func SetMicrosoftCdoAlldayevent(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoAlldayevent, value, params...)
}

// SetMicrosoftCdoApptSequence sets the X-MICROSOFT-CDO-APPT-SEQUENCE property for a component
func SetMicrosoftCdoApptSequence(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoApptSequence, value, params...)
}

// SetMicrosoftCdoAttendeeCriticalChange sets the X-MICROSOFT-CDO-ATTENDEE-CRITICAL-CHANGE property for a component
func SetMicrosoftCdoAttendeeCriticalChange(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoAttendeeCriticalChange, value, params...)
}

// SetMicrosoftCdoBusystatus sets the X-MICROSOFT-CDO-BUSYSTATUS property for a component
func SetMicrosoftCdoBusystatus(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoBusystatus, value, params...)
}

// SetMicrosoftCdoImportance sets the X-MICROSOFT-CDO-IMPORTANCE property for a component
func SetMicrosoftCdoImportance(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoImportance, value, params...)
}

// SetMicrosoftCdoInsttype sets the X-MICROSOFT-CDO-INSTTYPE property for a component
func SetMicrosoftCdoInsttype(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoInsttype, value, params...)
}

// SetMicrosoftCdoIntendedstatus sets the X-MICROSOFT-CDO-INTENDEDSTATUS property for a component
func SetMicrosoftCdoIntendedstatus(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoIntendedstatus, value, params...)
}

// SetMicrosoftCdoOwnerapptid sets the X-MICROSOFT-CDO-OWNERAPPTID property for a component
func SetMicrosoftCdoOwnerapptid(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoOwnerapptid, value, params...)
}

// SetMicrosoftCdoOwnerCriticalChange sets the X-MICROSOFT-CDO-OWNER-CRITICAL-CHANGE property for a component
func SetMicrosoftCdoOwnerCriticalChange(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoOwnerCriticalChange, value, params...)
}

// SetMicrosoftCdoReplytime sets the X-MICROSOFT-CDO-REPLYTIME property for a component
func SetMicrosoftCdoReplytime(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftCdoReplytime, value, params...)
}

// SetMicrosoftDisallowCounter sets the X-MICROSOFT-DISALLOW-COUNTER property for a component
func SetMicrosoftDisallowCounter(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftDisallowCounter, value, params...)
}

// SetMicrosoftExdate sets the X-MICROSOFT-EXDATE property for a component
func SetMicrosoftExdate(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftExdate, value, params...)
}

// SetMicrosoftIsdraft sets the X-MICROSOFT-ISDRAFT property for a component
func SetMicrosoftIsdraft(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftIsdraft, value, params...)
}

// SetMicrosoftMsncalendarAlldayevent sets the X-MICROSOFT-MSNCALENDAR-ALLDAYEVENT property for a component
func SetMicrosoftMsncalendarAlldayevent(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftMsncalendarAlldayevent, value, params...)
}

// SetMicrosoftMsncalendarBusystatus sets the X-MICROSOFT-MSNCALENDAR-BUSYSTATUS property for a component
func SetMicrosoftMsncalendarBusystatus(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftMsncalendarBusystatus, value, params...)
}

// SetMicrosoftMsncalendarImportance sets the X-MICROSOFT-MSNCALENDAR-IMPORTANCE property for a component
func SetMicrosoftMsncalendarImportance(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftMsncalendarImportance, value, params...)
}

// SetMicrosoftMsncalendarIntendedstatus sets the X-MICROSOFT-MSNCALENDAR-INTENDEDSTATUS property for a component
func SetMicrosoftMsncalendarIntendedstatus(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftMsncalendarIntendedstatus, value, params...)
}

// SetMicrosoftRrule sets the X-MICROSOFT-RRULE property for a component
func SetMicrosoftRrule(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMicrosoftRrule, value, params...)
}

// SetMsOlkAllowexterncheck sets the X-MS-OLK-ALLOWEXTERNCHECK property for a component
func SetMsOlkAllowexterncheck(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkAllowexterncheck, value, params...)
}

// SetMsOlkApptlastsequence sets the X-MS-OLK-APPTLASTSEQUENCE property for a component
func SetMsOlkApptlastsequence(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkApptlastsequence, value, params...)
}

// SetMsOlkApptseqtime sets the X-MS-OLK-APPTSEQTIME property for a component
func SetMsOlkApptseqtime(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkApptseqtime, value, params...)
}

// SetMsOlkAutofilllocation sets the X-MS-OLK-AUTOFILLLOCATION property for a component
func SetMsOlkAutofilllocation(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkAutofilllocation, value, params...)
}

// SetMsOlkAutostartcheck sets the X-MS-OLK-AUTOSTARTCHECK property for a component
func SetMsOlkAutostartcheck(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkAutostartcheck, value, params...)
}

// SetMsOlkCollaboratedoc sets the X-MS-OLK-COLLABORATEDOC property for a component
func SetMsOlkCollaboratedoc(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkCollaboratedoc, value, params...)
}

// SetMsOlkConfcheck sets the X-MS-OLK-CONFCHECK property for a component
func SetMsOlkConfcheck(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkConfcheck, value, params...)
}

// SetMsOlkConftype sets the X-MS-OLK-CONFTYPE property for a component
func SetMsOlkConftype(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkConftype, value, params...)
}

// SetMsOlkDirectory sets the X-MS-OLK-DIRECTORY property for a component
func SetMsOlkDirectory(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkDirectory, value, params...)
}

// SetMsOlkMwsurl sets the X-MS-OLK-MWSURL property for a component
func SetMsOlkMwsurl(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkMwsurl, value, params...)
}

// SetMsOlkNetshowurl sets the X-MS-OLK-NETSHOWURL property for a component
func SetMsOlkNetshowurl(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkNetshowurl, value, params...)
}

// SetMsOlkOnlinepassword sets the X-MS-OLK-ONLINEPASSWORD property for a component
func SetMsOlkOnlinepassword(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkOnlinepassword, value, params...)
}

// SetMsOlkOrgalias sets the X-MS-OLK-ORGALIAS property for a component
func SetMsOlkOrgalias(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkOrgalias, value, params...)
}

// SetMsOlkOriginalend sets the X-MS-OLK-ORIGINALEND property for a component
func SetMsOlkOriginalend(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkOriginalend, value, params...)
}

// SetMsOlkOriginalstart sets the X-MS-OLK-ORIGINALSTART property for a component
func SetMsOlkOriginalstart(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkOriginalstart, value, params...)
}

// SetMsOlkSender sets the X-MS-OLK-SENDER property for a component
func SetMsOlkSender(c ical.Component, value string, params ...ical.PropertyParameter) {
	SetComponentProperty(c, ComponentPropertyMsOlkSender, value, params...)
}
