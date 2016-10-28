package pdfprinter

import (
    "testing"
    "time"
)

func TestParseFile(t *testing.T) {
    ts := new(Timesheet)
    ts.ParseFile("arbeitszeit.csv")

    expectedDuration, err := time.ParseDuration("137h10m0s")
    if err != nil {
        t.Error("Error while parsing expected value")
    }
    if ts.totalWorkTime != expectedDuration {
        t.Error("TotalWorkTime: Expected %v, got %v", expectedDuration, ts.totalWorkTime)
    }

    if len(ts.lineItems) != 17 {
        t.Error("lineitems: Expected 17, got ", len(ts.lineItems))
    }

    if ts.periodStartDate.IsZero() {
        t.Error("start time: Expected something, got Zero")
    }
    if ts.periodEndDate.IsZero() {
        t.Error("end time: Expected something, got Zero")
    }
    hour, min, sec := ts.periodStartDate.Clock()
    year, month, day := ts.periodStartDate.Date()
    if hour != 10 {
        t.Error("start time: Expected hour 10, got", hour)
    }
    if min != 0 {
        t.Error("start time: Expected min 0, got", min)
    }
    if sec != 0 {
        t.Error("start time: Expected sec 0, got", sec)
    }
    if year != 2016 {
        t.Error("start date: Expected year 2016, got", year)
    }
    if month != time.September {
        t.Error("start date: Expected september, got", month)
    }
    if day != 1 {
        t.Error("start date: Expected day 1, got", day)
    }

    hour, min, sec = ts.periodEndDate.Clock()
    year, month, day = ts.periodEndDate.Date()
    if hour != 18 {
        t.Error("end time: Expected hour 18, got", hour)
    }
    if min != 0 {
        t.Error("end time: Expected hour 0, got", min)
    }
    if sec != 0 {
        t.Error("end time: Expected sec 0, got", sec)
    }
    if year != 2016 {
        t.Error("end date: Expected year 2016, got", year)
    }
    if month != time.September {
        t.Error("end date: Expected september, got", month)
    }
    if day != 29 {
        t.Error("end date: Expected day 29, got", day)
    }

}



func TestGetNormalizedWorktime(t *testing.T) {
    ts := new(Timesheet)
    ts.ParseFile("arbeitszeit.csv")

    days, hours, minutes := ts.GetNormalizedWorktime()
    if days != 17 {
        t.Error("Days: Expected 17, got ", days)
    }

    if hours != 1 {
        t.Error("Hours: Expected 1, got ", hours)
    }
    if minutes != 10 {
        t.Error("Minutes Expected 10, got ", minutes)
    }
}
