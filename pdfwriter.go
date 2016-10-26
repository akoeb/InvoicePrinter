package main
import (
    "log"
    "strings"
    "github.com/signintech/gopdf"
    "github.com/signintech/gopdf/fontmaker/core"
)

type PdfWriter struct {
    pdf gopdf.GoPdf
    fontSize int
    lineHeight float64
}
func (self *PdfWriter) init() {
    self.pdf = gopdf.GoPdf{}
    self.pdf.Start(gopdf.Config{ PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
    self.pdf.AddPage()

    err := self.pdf.AddTTFFont("default", "Times_New_Roman.ttf")
    if err != nil {
        log.Print(err.Error())
        return
    }

    self.fontSize = 11
    err = self.pdf.SetFont("default", "", self.fontSize)
    if err != nil {
        log.Print(err.Error())
        return
    }

  	//Measure Height
	//get  CapHeight (https://en.wikipedia.org/wiki/Cap_height)
	var parser core.TTFParser
	err = parser.Parse("Times_New_Roman.ttf")
	if err != nil {
	    log.Print(err.Error())
		return
	}
      
    cap := float64(float64(parser.CapHeight()) * 1000.00 / float64(parser.UnitsPerEm()))
	//convert
	self.lineHeight = cap * (float64(self.fontSize) / 1000.0) + 2.25
	  
}
func (self *PdfWriter) Text(text string) {
    twidth := float64(0)
    ttext := ""
    for _,word := range strings.Split(text, " ") {
        width,err := self.pdf.MeasureTextWidth(word + " ")
    	if err != nil {
	        log.Print(err.Error())
		    return
	    }
        if twidth + width < 500 {
            ttext = ttext + " " + word
            twidth += width
        } else {
            // long line, print the one before, print a new line and reset globals:
            self.pdf.Cell(nil, ttext + " " + word)
            self.pdf.Br(self.lineHeight)      
            ttext = ""
            twidth = float64(0)
        }
    }
    self.pdf.Cell(nil, ttext)
}
    //Measure Width
	  //text := "How can I cordinate the text that I want draw?"
	  //pdf.Cell(nil, text)
	  //realWidth, _ := pdf.MeasureTextWidth(text)
	  //fmt.Printf("realWidth = %f", realWidth)


func (self *PdfWriter) Write() {
    self.pdf.WritePdf("hello.pdf")
}

func (self *PdfWriter) Br() {
    self.pdf.Br(self.lineHeight)
}

func main() {


    pdf := new(PdfWriter)
    pdf.init()
    
    pdf.Text("Alexander Köb")
    pdf.Br()
    pdf.Text("Schönhauser Allee 58")
    pdf.Br()
    pdf.Text("10437 Berlin")
    pdf.Br()
    pdf.Text("Ust ID Nr.: DE 2893 54 867")
    pdf.Br()

	pdf.Br()

    pdf.Text( "An:")
	pdf.Br()
    pdf.Text("ITinera projects & experts GmbH & Co.KG")
	pdf.Br()
    pdf.Text("Mergenthalerallee 79-81")
	pdf.Br()
    pdf.Text("65760 Eschborn")
	pdf.Br()

    pdf.Br()
    
    pdf.Text("Berlin, den 06.10.2016") // right aligned
	pdf.Br()
    
    pdf.Text("Rechnung Nr. 2016-10-036") // bold
	pdf.Br()
    pdf.Text("(bei Zahlung bitte angeben)") // smaller
    pdf.Br()

    pdf.Br()
    
    pdf.Text("Sehr geehrte Damen und Herren,")
	pdf.Br()

    pdf.Br()
    
    pdf.Text("hiermit stelle ich Ihnen für meine Tätigkeiten bei der Firma Zalando SE nach Projektauftrag Nr. 1 vom 10.03.2016 für den Leistungszeitraum 01.09.2016 bis 30.09.2016 folgendes in Rechnung:")
	pdf.Br()


    // as table:
    pdf.Text("Pos") // bold, border
    pdf.Text("Bezeichnung") // bold, border
    pdf.Text("Menge") // bold, border
    pdf.Text("Einzelpreis EUR") // bold, border
    pdf.Text("Gesamtpreis EUR") // bold, border

    // missing: line items
    // missing sum, tax, etc
    pdf.Br()
	pdf.Br()

    pdf.Text("Bitte überweisen Sie den Rechnungsbetrag bis zum 06.11.2016 auf folgendes Konto:")
    
    // br

    // table
    pdf.Text("Kontoinhaber") // left
    pdf.Text("Alexander Köb") // right

    pdf.Text("Bank")
    pdf.Text("GLS Bank")
    
    pdf.Text("IBAN")
    pdf.Text("DE21430609671151010200")
    
    pdf.Text("BIC")
    pdf.Text("GENODEM1GLS")
    
	pdf.Br()
    // br

    pdf.Text("Mit freundlichen Grüßen")

    pdf.Text("Alexander Köb")
    pdf.Write()

  }