package gocal

import (
	"github.com/teambition/rrule-go"
)

func (gc *Gocal) ExpandRecurringEvent(buf *Event) ([]Event, error) {
	rOption, err := rrule.StrToROptionInLocation(buf.RecurrenceRuleString,
		buf.Start.Location())
	if err != nil {
		return nil, err
	}

	r, err := rrule.NewRRule(*rOption)
	if err != nil {
		return nil, err
	}

	s := rrule.Set{}
	s.RRule(r)
	s.DTStart(*buf.Start)
	s.SetExDates(buf.ExcludeDates)

	endOffset := buf.End.Sub(*buf.Start)

	evs := []Event{}
	for _, occ := range s.Between(*gc.Start, *gc.End, true) {
		start := occ
		end := start.Add(endOffset)

		e := *buf
		e.Start = &start
		e.End = &end
		e.Uid = buf.Uid

		evs = append(evs, e)
	}

	return evs, nil
}
