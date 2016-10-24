package main

import (
    "time"
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

type Invoice struct {
    lineItems []InvoiceLineItem
    invoiceTime time.Time
    serviceStartTime time.Time
    serviceEndTime time.Time
    dueDate time.Time
    totalNetSum int
    totalTaxValue int
    totalGrossSum int
}
func(self *Invoice) setInvoiceDate(invoiceTime time.Time) {
    if invoiceTime.IsZero() {
        self.invoiceTime = time.Now()
    } else {
        self.invoiceTime = invoiceTime
    }
}
func(self *Invoice) setDueDate(dueDate time.Time) {
    if dueDate.IsZero() {
        // default in one month
        self.dueDate = time.Now().AddDate(0,1,0)
    } else {
        self.dueDate = dueDate
    }
}
func(self *Invoice) setServiceTime(startDate time.Time, endDate time.Time) {
    self.serviceStartTime = startDate
    self.serviceEndTime = endDate
}
func(self *Invoice) addLineItem(amount int, itemPrice int, taxRate int, description string) {
    li := new(InvoiceLineItem)
    li.amount = amount
    li.itemPrice = itemPrice
    li.taxRate = taxRate
    li.description = description

    li.netSum = li.amount * li.itemPrice
    li.taxValue = li.netSum * li.taxRate / 100
    li.grossSum = li.netSum + li.taxValue
    self.lineItems = append(self.lineItems, *li)

    // initialize globals in case they have not been initialized
    // and add the line item values to totals
    self.totalNetSum += li.netSum
    self.totalGrossSum += li.grossSum
    self.totalTaxValue += li.taxValue
}

// TODO: add other data:
//       * client data
// TODO: that goes in the template:
//       * my data
//       * bank account information
//       * signature and other texts
