package e2e

import "sync"

type mockWriter struct {
	received string
	mut      sync.Mutex
}

func NewMockWriter() *mockWriter {
	return &mockWriter{}
}

func (w *mockWriter) Write(p []byte) (int, error) {
	w.mut.Lock()
	defer w.mut.Unlock()
	w.received += string(p)
	return len(p), nil
}

func (w *mockWriter) Received() string {
	w.mut.Lock()
	defer w.mut.Unlock()
	return w.received
}
