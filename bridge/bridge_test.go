package bridge

import (
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Error(err)
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}
	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrMsg := "you need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Errorf("Error message was not correct. \n got: %s\n expected: %s \n", err.Error(), expectedErrMsg)
		}
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: new(PrinterImpl1),
	}

	err := normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	testWriter := new(TestWriter)
	normal = NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl2{Writer: testWriter},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("expect:%s got:%s", expectedMessage, testWriter.Msg)
	}
}

func TestPacketPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packet: Hello io.Writer"

	packt := PacketPrinter{
		Msg:     passedMessage,
		Printer: new(PrinterImpl1),
	}

	err := packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	testWriter := new(TestWriter)
	packt = PacketPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl2{Writer: testWriter},
	}

	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("expect:%s got:%s", expectedMessage, testWriter.Msg)
	}
}
