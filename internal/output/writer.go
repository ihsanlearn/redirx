package output

import (
	"os"
	"sync"

	"github.com/ihsanlearn/redirx/pkg/logger"
)

type Writer struct {
	file *os.File
	mu   sync.Mutex
}

func NewWriter(filename string) (*Writer, error) {
	if filename == "" {
		return &Writer{}, nil
	}

	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Writer{file: file}, nil
}

func (w *Writer) Write(data string) {
	if w.file == nil {
		return
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	_, err := w.file.WriteString(data + "\n")
	if err != nil {
		logger.Error("%s", err)
	}
}

func (w *Writer) Close() {
	if w.file != nil {
		w.file.Close()
	}
}