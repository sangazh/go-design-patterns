package adapter

import (
	"io"
	"os"
	"testing"
)

func TestCounter(t *testing.T) {
	c := new(Counter)
	c.Count(3)
}

func TestCounterWriter(t *testing.T) {
	counter := CounterWriter{Writer: os.Stdout}
	counter.Count(5)
}

func TestCounterWriterPipe(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()

	counter := CounterWriter{Writer: pipeWriter}

	f, _ := os.Create("pipe")
	tee := io.TeeReader(pipeReader, f)

	go func() {
		io.Copy(os.Stdout, tee)
	}()

	counter.Count(5)
}
