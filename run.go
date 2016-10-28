package main

import (
    //"os"
    //"fmt"
    "time"
    "pdfprinter"
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
    ts := new(pdfprinter.Timesheet)
    ts.ParseFile("arbeitszeit.csv")
    days, hours, _ := ts.GetNormalizedWorktime()

    iv := new(pdfprinter.Invoice)
    iv.AddLineItem(days, 540, 19, "day rate")
    iv.AddLineItem(hours, 67.50, 19, "hour rate")
    iv.SetInvoiceDate(time.Now())
    iv.SetDueDate(time.Now().AddDate(0,1,0))
    iv.SetServiceDateRange(ts.PeriodStartDate, ts.PeriodEndDate)

    //fmt.Printf("Normalized Worktime: days %v, hours: %v, minutes: %v\n", days, hours, minutes)
    //fmt.Printf("net: %v, gross: %v, tax: %v\n", float64(iv.TotalNetSum) / 10000, float64(iv.TotalGrossSum) / 10000, float64(iv.TotalTaxValue) / 10000)
    //fmt.Printf("INvoice: %v\n", iv)
    // fmt.Printf("Test1 %.02f\n", 1.425)
    // fmt.Printf("Test2 %.02f\n", 1.525)
    // fmt.Printf("Test3 %.02f\n", 1.625)

    pdf := new(pdfprinter.PdfWriter)
    pdf.Init()
    pdf.WriteSender("Alexander Köb\nSchönhauser Allee 58\n10437 Berlin\nUst ID Nr.: DE 2893 54 867")
    pdf.WriteRecipient("An:\nITinera projects & experts GmbH & Co.KG\nMergenthalerallee 79-81\n65760 Eschborn")
    pdf.WriteDate("Berlin, den 06.10.2016")
    pdf.WriteSubject("Rechnung Nr. 2016-10-036", "(bei Zahlung bitte angeben)")
    pdf.WriteInvoice(*iv)
}
