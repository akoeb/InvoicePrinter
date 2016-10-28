package pdfprinter
import (
    "log"
    "fmt"
    "github.com/jung-kurt/gofpdf"
)

type PdfWriter struct {
    pdf *gofpdf.Fpdf
    tr func(string) string
}

func (self *PdfWriter) Init() {
    self.pdf = gofpdf.New("P", "mm", "A4", "")
    self.tr = self.pdf.UnicodeTranslatorFromDescriptor("")
    self.pdf.SetTextColor(0, 0, 0)
    self.pdf.SetTitle("Invoice", true)
    self.pdf.SetAuthor("Alexander Köb", true)
    self.pdf.AddPage()
    self.pdf.SetFont("times", "", 11)
    self.pdf.Ln(3)
}

func(self *PdfWriter) WriteSender(sender string) {
    self.pdf.MultiCell(0, 4.5, self.tr(sender), "", "L", false)
    self.pdf.Ln(8)
}

func (self *PdfWriter) WriteRecipient(recipient string) {
    self.pdf.MultiCell(0, 4.5, self.tr(recipient), "", "L", false)
	self.pdf.Ln(8)
}

func (self *PdfWriter) WriteDate(date string) {
    self.pdf.MultiCell(0, 4.5, self.tr(date), "", "R", false) // right aligned
    self.pdf.Ln(8)
}

func (self *PdfWriter) WriteSubject(subject string, second string) {
    self.pdf.SetFont("times", "B", 11)
    self.pdf.MultiCell(0, 5, self.tr(subject), "", "L", false) // bold
    if second != "" {
        self.pdf.SetFont("times", "", 10)
        self.pdf.MultiCell(0, 5, self.tr(second), "", "LT", false) // smaller
    }
    self.pdf.SetFont("times", "", 11)
    self.pdf.Ln(8)
}


func (self *PdfWriter) WriteInvoice(iv Invoice) {


    initText := "Sehr geehrte Damen und Herren,\n\n"
    initText += "    hiermit stelle ich Ihnen für meine Tätigkeiten bei der Firma Zalando SE nach "
    initText += "Projektauftrag Nr. 1 vom 10.03.2016 für den Leistungszeitraum "
    initText += fmt.Sprintf("%s bis %s ", iv.ServiceStartDate(), iv.ServiceEndDate())
    initText += "folgendes in Rechnung:"
    self.pdf.MultiCell(0, 5, self.tr(initText), "", "L", false)
	self.pdf.Ln(8)

    // as table:
    self.pdf.SetFont("times", "B", 11)
    self.pdf.SetX(20)
    self.pdf.CellFormat(10, 5, "Pos", "1", 0, "C", false, 0, "")
    self.pdf.CellFormat(65, 5, "Bezeichnung", "1", 0, "C", false, 0, "")
    self.pdf.CellFormat(15, 5, "Menge", "1", 0, "C", false, 0, "")
    self.pdf.CellFormat(40, 5, "Einzelpreis" , "1", 0, "C", false, 0, "")
    self.pdf.CellFormat(40, 5, "Gesamtpreis", "1", 0, "C", false, 0, "")
    self.pdf.Ln(-1)
    self.pdf.SetFont("times", "", 11)
    for idx, item := range iv.LineItems {
        self.pdf.SetX(20)
        self.pdf.CellFormat(10, 5, fmt.Sprintf("%d",idx + 1), "1", 0, "C", false, 0, "")
        self.pdf.CellFormat(65, 5, item.Description(), "1", 0, "L", false, 0, "")
        self.pdf.CellFormat(15, 5, item.Amount(), "1", 0, "C", false, 0, "")
        self.pdf.CellFormat(40, 5, item.ItemPrice() + " EUR ", "1", 0, "R", false, 0, "")
        self.pdf.CellFormat(40, 5, item.NetSum() + " EUR ", "1", 0, "R", false, 0, "")
        self.pdf.Ln(-1)
    }

    self.pdf.SetX(110)
    self.pdf.CellFormat(40, 5, "Summe" , "1", 0, "L", false, 0, "")
    self.pdf.CellFormat(40, 5, iv.TotalNetSum() + " EUR ", "1", 0, "R", false, 0, "")
    self.pdf.Ln(-1)

    self.pdf.SetX(110)
    self.pdf.CellFormat(40, 5, "+19% Umsatzsteuer" , "1", 0, "L", false, 0, "")
    self.pdf.CellFormat(40, 5, iv.TotalTaxValue() + " EUR ", "1", 0, "R", false, 0, "")
    self.pdf.Ln(-1)

    self.pdf.SetFont("times", "B", 11)
    self.pdf.SetX(110)
    self.pdf.CellFormat(40, 5, "Rechnungsbetrag" , "1", 0, "L", false, 0, "")
    self.pdf.CellFormat(40, 5, iv.TotalGrossSum() + " EUR ", "1", 0, "R", false, 0, "")
    self.pdf.Ln(-1)
    self.pdf.SetFont("times", "", 11)

    self.pdf.SetY(-100)

    self.pdf.MultiCell(0, 5, self.tr("Bitte überweisen Sie den Rechnungsbetrag bis zum " + iv.DueDate() + " auf folgendes Konto:"), "", "L", false)
    self.pdf.Ln(5)

    self.pdf.SetX(60)
    self.pdf.CellFormat(35, 5, "Kontoinhaber:" , "1", 0, "L", false, 0, "")
    self.pdf.CellFormat(55, 5, self.tr("Alexander Köb"), "1", 0, "R", false, 0, "")
    self.pdf.Ln(-1)

    self.pdf.SetX(60)
    self.pdf.CellFormat(35, 5, "Bank:" , "1", 0, "L", false, 0, "")
    self.pdf.CellFormat(55, 5, self.tr("GLS Bank"), "1", 0, "R", false, 0, "")
    self.pdf.Ln(-1)

    self.pdf.SetX(60)
    self.pdf.CellFormat(35, 5, "IBAN:" , "1", 0, "L", false, 0, "")
    self.pdf.CellFormat(55, 5, "DE21430609671151010200", "1", 0, "R", false, 0, "")
    self.pdf.Ln(-1)

    self.pdf.SetX(60)
    self.pdf.CellFormat(35, 5, "BIC:" , "1", 0, "L", false, 0, "")
    self.pdf.CellFormat(55, 5, "GENODEM1GLS", "1", 0, "R", false, 0, "")
    
    self.pdf.Ln(10)
    
    self.pdf.MultiCell(0, 5, self.tr("Mit freundlichen Grüßen"), "", "L", false)
    self.pdf.Ln(30)
    self.pdf.MultiCell(0, 5, self.tr("Alexander Köb"), "", "L", false)

    // error check: only one page!
    if self.pdf.PageNo() != 1 {
        log.Println("Error: incorrect page number: " + string(self.pdf.PageNo()))
    }
    err := self.pdf.OutputFileAndClose("hello.pdf")
    if err != nil {
        log.Println("Error in writing pdf: ", err)
    }
  }
