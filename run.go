package main

import (
    //"os"
    "fmt"
    "flag"
    "time"
    "log"
    "path/filepath"
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


func file_exists(filename *string) bool {
    if _, err := os.Stat(*filename); os.IsNotExist(err) {
        return false
    }
    return true
}
func getInvoiceFromDir(dirname *string) *string {
    files, _ := filepath.Glob(dirname + "/invoice_*_signed.pdf")
    for _, file := range files {
        // TODO: find highest invoice id
    }
}

func main() {


    // flag type (timesheet, invoice, all (default))
    docType := flag.String("type", "both", "which document type to print [invoice, timesheet, both]. Defaults to both")
    // flag input file
    inFile := flag.String("in", "", "The Input file, a CSV file with the collection of actually worked days for the specified time period.")
    // flag output dir invoices (obsolete if --type timesheet)
    invoiceDir := flag.String("invoiceDir", "", "directory to place the written invoice (and look for invoice ID)")
    // flag output dir timesheet (obsolete if --type invoice)
    timesheetDir := flag.String("timesheetDir", "", "directory to place the written timesheet")
    // flag invoice ID (obsolete if --type timesheet, per default it is read from --invoice-output-dir)
    invoiceId := flag.String("invoiceId", "", "invoice ID to put on the invoice document")

    flag.Parse()

    // validate flags

    // infile must be given and exist
    if *inFile == "" {
        // no infile defined
        log.Fatal("No input file argument is defined.")
    } else if ! file_exists (*inFile){
        // infile does not exist
        log.Fatal("the specified input file does not exist.")
    }

    switc *docType  {
    case "invoice":
        if *timesheetDir != "" {
            // obsolete timesheet dir in invoice
            log.Fatal("The invoice document type does not need a timesheet dir.")
        }
        if *invoiceDir == "" {
            // error: document type invoice needs invoice dir
            log.Fatal("The invoice document type needs an invoice dir.")
        } else if ! file_exists (*invoiceDir){
            // error: invoice dir does  not exist
            log.Fatal("The specified invoice dir does not exist.")
        } else {
        }
        if *invoiceId == "" {
            invoiceId = getInvoiceIdFromDir(invoiceDir)
        }
        if *invoiceId == "" {
            // error: no ID fround from looking at output directory
            log.Fatal("No invoice ID specified, and we could not determine one ourselves.")
        }

        break;
    case "timesheet":
        if *invoiceDir != "" {
            // obsolete invoice dir in invoice
            log.Fatal("The timesheet document type does not need an invoice dir.")
        }
        if *timesheetDir == "" {
            // error: document type timesheet needs timesheet dir
            log.Fatal("The timesheet document type needs a timesheet dir.")

        } else if ! file_exists (*timesheetDir){
            // error: timesheet dir does  not exist
            log.Fatal("The specified timesheet dir does not exist.")
        }

    case "both":

        // no error condition here

    default:
        // TODO error: unknown doc type
        log.Fatal("Unknown document type found")
    }




    ts := new(pdfprinter.Timesheet)
    ts.ParseFile(*inFile)
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
    pdf.WriteSubject(*invoiceId, "(bei Zahlung bitte angeben)")
    pdf.WriteInvoice(*iv)
}
