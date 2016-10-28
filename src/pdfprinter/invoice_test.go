package pdfprinter

import (
    "testing"
    "time"
)
// global times for use in tests
var (
    t1 =  time.Date(2016, time.July, 1, 10, 0, 0, 0, time.UTC)
    t2 =  time.Date(2016, time.July, 31, 18, 10, 0, 0, time.UTC)

    days = 15
    dayRate = 1600000
    tax = 19
    dayDesc = "day rate"
    hours = 5
    hourRate = dayRate / 8
    hourDesc = "hour rate"

)

func TestSetInvoiceDate(t *testing.T) {
    iv := new(Invoice)
    iv.SetInvoiceDate(t1)
    if iv.invoiceDate != t1 {
        t.Error("Invoice time: Expected %v, got %v", t1, iv.invoiceDate)
    }
    // null value returns Now()
    before := time.Now()
    iv.SetInvoiceDate(time.Time{})
    after := time.Now()
    if iv.invoiceDate.Before(before) {
        t.Error("Invoice time should not come before %v but is %v", before, iv.invoiceDate)
    }
    if iv.invoiceDate.After(after) {
        t.Error("Invoice time should not come after %v but is %v", after, iv.invoiceDate)
    }
}
func TestSetDueDate(t *testing.T) {
    iv := new(Invoice)
    iv.SetDueDate(t1)
    if iv.dueDate != t1 {
        t.Error("Due Date: Expected %v but got %v", t1, iv.dueDate)
    }
    // null value returns Now() + 1 month
    before := time.Now().AddDate(0,1,0)
    iv.SetDueDate(time.Time{})
    after := time.Now().AddDate(0,1,0)
    if iv.dueDate.Before(before) {
        t.Error("due date should not come before %v but is %v", before, iv.dueDate)
    }
    if iv.dueDate.After(after) {
        t.Error("due date should not come after %v but is %v", after, iv.dueDate)
    }

    // check for one month after invoice time:
    iv.SetInvoiceDate(t1)
    iv.SetDueDate(time.Time{})
    expected := t1.AddDate(0,1,0)
    if iv.dueDate != expected {
        t.Error("Due Date: Expected %v but is %v", expected, iv.dueDate)
    }
}
func TestSetServiceTime(t *testing.T) {
    iv := new(Invoice)
    iv.SetServiceDateRange(t1, t2)
    if iv.serviceStartDate != t1 {
        t.Error("Service Date Start: Expected %v but is %v", t1, iv.serviceStartDate)
    }
    if iv.serviceEndDate != t2 {
        t.Error("Service Date End: Expected %v but is %v", t2, iv.serviceEndDate)
    }
}

func TestAddLineItem(t *testing.T) {
    iv := new(Invoice)
    // params to set:
    iv.AddLineItem(days, dayRate, tax, dayDesc)
    iv.AddLineItem(hours, hourRate, tax, hourDesc)

    _testInvoiceLineItems(t, *iv)
}

func TestSetLineItems(t *testing.T) {
    // set initial state with one line item:
    iv := new(Invoice)
    // params to set:
    beforeAmount := 15
    beforeItemPrice := 1600000
    beforeTax := 19
    beforeDesc := "day rate"
    iv.AddLineItem(beforeAmount, beforeItemPrice, beforeTax, beforeDesc)

    // create line items and call setLineItems
    li1 := new(InvoiceLineItem)
    li1.Set(days, dayRate, tax, dayDesc)
    li2 := new(InvoiceLineItem)
    li2.Set(hours, hourRate, tax, hourDesc)
    items :=  []InvoiceLineItem{ *li1, *li2}
    iv.SetLineItems(items)

    _testInvoiceLineItems(t, *iv)
}

func TestLineItemSet(t *testing.T) {
    // set initial state with one line item:
    li := new(InvoiceLineItem)
    li.Set(days, dayRate, tax, dayDesc)
    _testLineItem(t, li, days, dayRate, tax, dayDesc)
}


func _testInvoiceLineItems(t *testing.T, iv Invoice) {

    t.Log("Entering _testInvoiceLineItems")

    // two line items should be set
    if len(iv.lineItems) != 2 {
        t.Error("AddLineItem: Expected array with %v elements, but is %v", 2, len(iv.lineItems))
    }

    // check the totals:
    expectedTotalNetSum := (days * dayRate) + (hours * hourRate)
    expectedTotalTaxValue := expectedTotalNetSum * tax / 100
    expectedTotalGrossSum := expectedTotalNetSum + expectedTotalTaxValue

    if iv.totalNetSum != expectedTotalNetSum {
        t.Error("AddLineItem: Expected total net sum of %v, but is %v", expectedTotalNetSum, iv.totalNetSum)
    }
    if iv.totalTaxValue != expectedTotalTaxValue {
        t.Error("AddLineItem: Expected total tax value of %v, but is %v", expectedTotalTaxValue, iv.totalTaxValue)
    }
    if iv.totalGrossSum != expectedTotalGrossSum {
        t.Error("AddLineItem: Expected total gross sum of %v, but is %v", expectedTotalGrossSum, iv.totalGrossSum)
    }

    // check the line items themselves
    // days
    _testLineItem(t, &iv.lineItems[0], days, dayRate, tax, dayDesc)
    // hours
    _testLineItem(t, &iv.lineItems[1], hours, hourRate, tax, hourDesc)
}

func _testLineItem(t *testing.T, li *InvoiceLineItem, expectedAmount int, expectedItemPrice int, expectedTaxRate int,  expectedDescription string) {
    t.Log("Entering _testLineItem")

    expectedNetSum := expectedAmount * expectedItemPrice
    expectedTaxValue := expectedNetSum * expectedTaxRate / 100
    expectedGrossSum := expectedNetSum + expectedTaxValue

    if li.netSum != expectedNetSum {
        t.Error("AddLineItem: Item 2: Expected net sum of %v, but is %v", expectedNetSum, li.netSum)
    }
    if li.taxRate != expectedTaxRate {
        t.Error("AddLineItem: Item 2: Expected tax %v, but is %v", expectedTaxRate, li.taxRate)
    }
    if li.taxValue != expectedTaxValue {
        t.Error("AddLineItem: Item 2: Expected tax value of %v, but is %v", expectedTaxValue, li.taxValue)
    }
    if li.grossSum != expectedGrossSum {
        t.Error("AddLineItem: Item 2: Expected gross sum %v, but is %v", expectedGrossSum, li.grossSum)
    }
    if li.itemPrice != expectedItemPrice {
        t.Error("AddLineItem: Item 2: Expected item price of %v, but is %v", expectedItemPrice, li.itemPrice)
    }
    if li.description != expectedDescription {
        t.Error("AddLineItem: Item 2: Expected description %v, but is %v", expectedDescription, li.description)
    }
    if li.amount != expectedAmount {
        t.Error("AddLineItem: Item 2: Expected amount %v, but is %v", expectedAmount, li.amount)
    }
}
