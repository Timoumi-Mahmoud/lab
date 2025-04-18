package main

import "fmt"

type myPrinter struct{}
type secondPrinter struct{}

type Printer interface {
	Print() string
}
type Scanner interface {
	Scan() string
}

type Faxer interface {
	Fax() string
}
type PrinterScanner interface {
	Printer
	Scanner
}
type PrinterFaxer interface {
	Printer
	Faxer
}

type FaxerPrinterScanner interface {
	Faxer
	Printer
	Scanner
}

func (mp myPrinter) Print() string {
	return "PRINTING one page ..."
}

func (mp myPrinter) Scan() string {
	return "Scaned one page ..."
}

func (sp secondPrinter) Print() string {
	return "PRINTING five page ..."

}

func (sp secondPrinter) Scan() string {
	return "Scaned five page ..."
}

func (mp myPrinter) Fax() string {
	return "Faxed one page ..."
}

func (sp secondPrinter) Fax() string {
	return "Faxed five page ..."
}

func process(equipment PrinterScanner) {
	fmt.Println("Running Print ...", equipment.Print())
	fmt.Println("Running Scanner ...", equipment.Scan())
}

func processComplite(equipment FaxerPrinterScanner) {
	fmt.Println("Running Print ...", equipment.Print())
	fmt.Println("Running Scanner ...", equipment.Scan())
	fmt.Println("Sending Fax ...", equipment.Fax())
}

func processPrintFax(equipment PrinterFaxer) {
	fmt.Println("Running Print ...", equipment.Print())
	fmt.Println("Sending Fax ...", equipment.Fax())
}
func main() {
	printer := myPrinter{}
	otherPrinter := secondPrinter{}
	process(printer)
	process(otherPrinter)
	fmt.Println()

	processComplite(printer)
	processComplite(otherPrinter)

	fmt.Println()
	processPrintFax(printer)
	processPrintFax(otherPrinter)

}
