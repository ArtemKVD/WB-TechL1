package main

import "fmt"

type SimplePrinter interface {
	Print(s string)
}

type Printer struct{}

func (p *Printer) Print(s string) {
	fmt.Println(s)
}

type DPrinter interface {
	PrintWithLength(s string)
}

type PrinterAdapter struct {
	simplePrinter SimplePrinter
}

func NewPrinterAdapter(printer SimplePrinter) *PrinterAdapter {
	return &PrinterAdapter{simplePrinter: printer}
}

func (a *PrinterAdapter) PrintWithLength(s string) {
	length := len(s)
	formatted := fmt.Sprintf("%s : len %v", s, length)
	a.simplePrinter.Print(formatted)
}

func main() {
	oldPrinter := &Printer{}
	adapter := NewPrinterAdapter(oldPrinter)
	adapter.PrintWithLength("Hello")
	oldPrinter.Print("Hello")
}
