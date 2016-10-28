package pdfprinter

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
)

// TODO: validation: https://husobee.github.io/golang/validation/2016/01/08/input-validation.html

type HttpInvoiceLineItem struct {
    taxRate     int
    itemPrice   int
    description string
    amount      int
}

type HttpInvoice struct {
    lineItems []HttpInvoiceLineItem
    invoiceDate time.Time
    serviceStartDate time.Time
    serviceEndDate time.Time
    dueDate time.Time
}
type HttpTimesheet struct {
}

func invoice_handler(w http.ResponseWriter, r *http.Request) {
    o := new(HttpInvoice)
    if r.Body == nil {
        http.Error(w, "Request body is empty", 400)
        return
    }
    if r.Method != "POST" {
        http.Error(w, "Method Not Allowed", 405)
        return
    }

    err := json.NewDecoder(r.Body).Decode(&o)
    if err != nil {
        http.Error(w, fmt.Sprintf("Validation Error: %s", err.Error()), 400)
        return
    }

    // TODO: validation

    iv := new(Invoice)
    iv.SetInvoiceDate(o.invoiceDate)
    iv.SetServiceDateRange(o.serviceStartDate, o.serviceEndDate)
    iv.SetDueDate(o.dueDate)
    // TODO: first map HttpInvoiceLineItem to InvoiceLineItem
    //iv.setLineItems(o.lineItems)

    fmt.Fprintf(w, "result: %v", iv)
    fmt.Fprintf(w, "\nInvoice %s!\n", r.URL.Path[1:])
}

func timesheet_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Timesheet %s!", r.URL.Path[1:])
}

//func main() {
//
//    http.HandleFunc("/invoice", invoice_handler)
//    http.HandleFunc("/timesheet", timesheet_handler)
//    http.ListenAndServe(":8080", nil)
//}
