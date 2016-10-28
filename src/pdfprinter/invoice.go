package pdfprinter

import (
    "time"
    "fmt"
    "strings"
)
// all prices are with 0.0001 Euro resolution (so int), currency is assumed euro
type InvoiceLineItem struct {
    netSum      int
    taxRate     int
    taxValue    int
    grossSum    int
    itemPrice   int
    description string
    amount      int
}
func(self *InvoiceLineItem) Set (amount int, itemPrice int, taxRate int, description string) {
    self.amount = amount
    self.itemPrice = itemPrice
    self.taxRate = taxRate
    self.description = description

    self.netSum = self.amount * self.itemPrice
    self.taxValue = self.netSum * self.taxRate / 100
    self.grossSum = self.netSum + self.taxValue
}
func(self *InvoiceLineItem) NetSum() string {
    return strings.Replace(fmt.Sprintf("%.2f", float64(self.netSum) / 10000.0),".", ",", -1)
}
func(self *InvoiceLineItem) TaxRate() string {
    return fmt.Sprintf("%d", self.taxRate)
}
func(self *InvoiceLineItem) TaxValue() string {
    return strings.Replace(fmt.Sprintf("%.2f", float64(self.taxValue) / 10000.0),".", ",", -1)
}
func(self *InvoiceLineItem) GrossSum() string {
    return strings.Replace(fmt.Sprintf("%.2f", float64(self.grossSum) / 10000.0),".", ",", -1)
}
func(self *InvoiceLineItem) ItemPrice() string {
    return strings.Replace(fmt.Sprintf("%.2f", float64(self.itemPrice) / 10000.0),".", ",", -1)
}
func(self *InvoiceLineItem) Description() string {
    return self.description
}
func(self *InvoiceLineItem) Amount() string {
    return fmt.Sprintf("%d", self.amount)
}

type Invoice struct {
    LineItems []InvoiceLineItem
    invoiceDate time.Time
    serviceStartDate time.Time
    serviceEndDate time.Time
    dueDate time.Time
    totalNetSum int
    totalTaxValue int
    totalGrossSum int
}
func(self *Invoice) InvoiceDate() string {
    return fmt.Sprintf("%02d.%02d.%d", self.invoiceDate.Day(), self.invoiceDate.Month(), self.invoiceDate.Year())
}
func(self *Invoice) ServiceEndDate() string {
    return fmt.Sprintf("%02d.%02d.%d", self.serviceEndDate.Day(), self.serviceEndDate.Month(), self.serviceEndDate.Year())
}
func(self *Invoice) ServiceStartDate() string {
    return fmt.Sprintf("%02d.%02d.%d", self.serviceStartDate.Day(), self.serviceStartDate.Month(), self.serviceStartDate.Year())
}
func(self *Invoice) DueDate() string {
    return fmt.Sprintf("%02d.%02d.%d", self.dueDate.Day(), self.dueDate.Month(), self.dueDate.Year())
}
func(self *Invoice) TotalNetSum() string {
    return strings.Replace(fmt.Sprintf("%.02f", float64(self.totalNetSum) / 10000.0),".", ",", -1)
}
func(self *Invoice) TotalTaxValue() string {
    return strings.Replace(fmt.Sprintf("%.02f", float64(self.totalTaxValue) / 10000.0),".", ",", -1)
}
func(self *Invoice) TotalGrossSum() string {
    return strings.Replace(fmt.Sprintf("%.03f", float64(self.totalGrossSum) / 10000.0),".", ",", -1)
}
func(self *Invoice) SetInvoiceDate(invoiceDate time.Time) {
    if invoiceDate.IsZero() {
        self.invoiceDate = time.Now()
    } else {
        self.invoiceDate = invoiceDate
    }
}
func(self *Invoice) SetDueDate(dueDate time.Time) {
    if dueDate.IsZero() {
        // default in one month or one month after invoice time
        if self.invoiceDate.IsZero() {
            self.dueDate = time.Now().AddDate(0,1,0)
        } else {
            self.dueDate = self.invoiceDate.AddDate(0,1,0)
        }
    } else {
        self.dueDate = dueDate
    }
}
func(self *Invoice) SetServiceDateRange(startDate time.Time, endDate time.Time) {
    self.serviceStartDate = startDate
    self.serviceEndDate = endDate
}
func(self *Invoice) AddLineItem(amount int, itemPrice int, taxRate int, description string) {
    li := new(InvoiceLineItem)
    li.Set(amount, itemPrice, taxRate, description)
    self.LineItems = append(self.LineItems, *li)

    // initialize globals in case they have not been initialized
    // and add the line item values to totals
    self.totalNetSum += li.netSum
    self.totalGrossSum += li.grossSum
    self.totalTaxValue += li.taxValue
}
func(self *Invoice) Clear() {
    // clear the line items and the totals:
    self.LineItems = nil
    self.totalNetSum = 0
    self.totalGrossSum = 0
    self.totalTaxValue = 0
}
func(self *Invoice) SetLineItems(items []InvoiceLineItem) {
    self.Clear()
    for _,li := range items {
        self.AddLineItem(li.amount, li.itemPrice, li.taxRate, li.description)
    }
}

// TODO: add other data:
//       * client data
// TODO: that goes in the template:
//       * my data
//       * bank account information
//       * signature and other texts
