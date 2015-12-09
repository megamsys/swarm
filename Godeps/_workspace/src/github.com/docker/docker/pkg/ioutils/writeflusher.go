package ioutils

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> origin/master
	"io"
	"net/http"
	"sync"
)

<<<<<<< HEAD
// WriteFlusher wraps the Write and Flush operation ensuring that every write
// is a flush. In addition, the Close method can be called to intercept
// Read/Write calls if the targets lifecycle has already ended.
type WriteFlusher struct {
	mu      sync.Mutex
	w       io.Writer
	flusher http.Flusher
	flushed bool
	closed  error

	// TODO(stevvooe): Use channel for closed instead, remove mutex. Using a
	// channel will allow one to properly order the operations.
}

var errWriteFlusherClosed = errors.New("writeflusher: closed")

func (wf *WriteFlusher) Write(b []byte) (n int, err error) {
	wf.mu.Lock()
	defer wf.mu.Unlock()
	if wf.closed != nil {
		return 0, wf.closed
	}

	n, err = wf.w.Write(b)
	wf.flush() // every write is a flush.
=======
// WriteFlusher wraps the Write and Flush operation.
type WriteFlusher struct {
	sync.Mutex
	w       io.Writer
	flusher http.Flusher
	flushed bool
}

func (wf *WriteFlusher) Write(b []byte) (n int, err error) {
	wf.Lock()
	defer wf.Unlock()
	n, err = wf.w.Write(b)
	wf.flushed = true
	wf.flusher.Flush()
>>>>>>> origin/master
	return n, err
}

// Flush the stream immediately.
func (wf *WriteFlusher) Flush() {
<<<<<<< HEAD
	wf.mu.Lock()
	defer wf.mu.Unlock()

	wf.flush()
}

// flush the stream immediately without taking a lock. Used internally.
func (wf *WriteFlusher) flush() {
	if wf.closed != nil {
		return
	}

=======
	wf.Lock()
	defer wf.Unlock()
>>>>>>> origin/master
	wf.flushed = true
	wf.flusher.Flush()
}

// Flushed returns the state of flushed.
// If it's flushed, return true, or else it return false.
func (wf *WriteFlusher) Flushed() bool {
<<<<<<< HEAD
	// BUG(stevvooe): Remove this method. Its use is inherently racy. Seems to
	// be used to detect whether or a response code has been issued or not.
	// Another hook should be used instead.
	wf.mu.Lock()
	defer wf.mu.Unlock()

	return wf.flushed
}

// Close closes the write flusher, disallowing any further writes to the
// target. After the flusher is closed, all calls to write or flush will
// result in an error.
func (wf *WriteFlusher) Close() error {
	wf.mu.Lock()
	defer wf.mu.Unlock()

	if wf.closed != nil {
		return wf.closed
	}

	wf.closed = errWriteFlusherClosed
	return nil
}

=======
	wf.Lock()
	defer wf.Unlock()
	return wf.flushed
}

>>>>>>> origin/master
// NewWriteFlusher returns a new WriteFlusher.
func NewWriteFlusher(w io.Writer) *WriteFlusher {
	var flusher http.Flusher
	if f, ok := w.(http.Flusher); ok {
		flusher = f
	} else {
		flusher = &NopFlusher{}
	}
	return &WriteFlusher{w: w, flusher: flusher}
}
