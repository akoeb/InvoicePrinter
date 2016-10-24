package main

import (
    //"os"
    "fmt"
    "time"
)

// TODO
// * create invoice data from time sheet
// * calculate earning, VAT, total from time sheet
// * a class for creating PDFS, one invoice, one time sheet type
// * the invoice generator gets the invice or timesheet structs passed in
// * output invoice or time sheet pdf
// * tests for time sheet
// * tests for invoices
// * add REST API
// * test for REST


func main() {

    ts := new(Timesheet)
    ts.parseFile("arbeitszeit.csv")
    days, hours, minutes := ts.getNormalizedWorktime()

    iv := new(Invoice)
    iv.addLineItem(days, 5400000, 19, "day rate")
    iv.addLineItem(hours, 675000, 19, "hour rate")
    iv.setInvoiceDate(time.Now())
    iv.setDueDate(time.Now().AddDate(0,1,0))
    iv.setServiceTime(ts.periodStartDate, ts.periodEndDate)

    fmt.Printf("Normalized Worktime: days %v, hours: %v, minutes: %v\n", days, hours, minutes)
    fmt.Printf("net: %v, gross: %v, tax: %v\n", float64(iv.totalNetSum) / 10000, float64(iv.totalGrossSum) / 10000, float64(iv.totalTaxValue) / 10000)
    fmt.Printf("INvoice: %v\n", iv)
}
