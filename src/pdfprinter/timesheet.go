package pdfprinter

import (
    "os"
    "io"
    "time"
    "log"
    "bufio"
    "encoding/csv"
    "strconv"
    "strings"
)

type TimesheetLineItem struct {
    TimeStart       time.Time
    TimeEnd         time.Time
    DurationPauses  time.Duration
    Description     string
    Reference       string
}

type Timesheet struct {
    LineItems []TimesheetLineItem
    TotalWorkTime time.Duration
    PeriodStartDate time.Time
    PeriodEndDate time.Time
}

func (self *Timesheet) ParseFile(filename string) {

    // Load the csv file.
    f, err := os.Open(filename)
    if err != nil {
        log.Fatal("Error in open CSV file: %v", err)
    }
    defer f.Close()

    // timezone of the timestamps is assumed to be local time
    location, err := time.LoadLocation("Local")
    if err != nil {
        log.Fatal("ERROR : %s", err)
        return
    }


    // Create a new reader and read the file
    r := csv.NewReader(bufio.NewReader(f))
    line := 0
    // TODO: calculates every line twice??
    for {
        line = line + 1
        record, err := r.Read()
        // end of file means stop
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal("Error in reading CSV file on line %d: %v", line, err)
        }

        // skip header:
        if line == 1 {
            continue
        }

        // create lineitem and fill it
        li := new(TimesheetLineItem)
        li.Reference = strings.Join(record, ",")

        // parse date, start and end
        day_a  := strings.Split(record[0], ".")
        start  := strings.Split(record[1], ":")
        end    := strings.Split(record[2], ":")
        pauses := strings.Split(record[3], ":")

        //log.Printf("record: %v, day %v, year %v", record[0], day_a, day_a[2] )

        // convert to integer:
        year, err  := strconv.Atoi(day_a[2])
        month, err := strconv.Atoi(day_a[1])
        day, err   := strconv.Atoi(day_a[0])
        startHour, err := strconv.Atoi(start[0])
        startMinute, err := strconv.Atoi(start[1])
        endHour, err := strconv.Atoi(end[0])
        endMinute, err := strconv.Atoi(end[1])

        li.TimeStart = time.Date(year, time.Month(month), day, startHour, startMinute, 0, 0, location)
        li.TimeEnd = time.Date(year, time.Month(month), day, endHour, endMinute, 0, 0, location)
        li.DurationPauses, err = time.ParseDuration(pauses[0] + "h" + pauses[1] + "m")
        li.Description = record[5]

        self.LineItems = append(self.LineItems, *li)

        // keep globals in sync
        self.TotalWorkTime += li.TimeEnd.Sub(li.TimeStart) - li.DurationPauses
        if self.PeriodStartDate.IsZero() || self.PeriodStartDate.After(li.TimeStart) {
            self.PeriodStartDate = li.TimeStart
        }
        if self.PeriodEndDate.IsZero() || self.PeriodEndDate.Before(li.TimeEnd) {
            self.PeriodEndDate = li.TimeEnd
        }

    }
}

// returns days, hours, minutes of work
func (self *Timesheet) GetNormalizedWorktime() (int, int, int) {
    var hours int = 0
    var days int = 0

    minutes := int(self.TotalWorkTime.Minutes())
    if minutes >= 60 {
        hours = minutes / 60
        minutes = minutes % 60
    }
    if hours >= 8 {
        days = hours / 8
        hours = hours % 8
    }
    return days, hours, minutes
}
