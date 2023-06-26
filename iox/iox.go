package iox

import "io"

type readerWithCloser struct {
	io.Reader
	close func() error
}

func (r readerWithCloser) Close() error {
	return r.close()
}

type writerWithCloser struct {
	io.Writer
	close func() error
}

func (w writerWithCloser) Close() error {
	return w.close()
}

// ReaderWithCloser wraps a [io.Reader] with an arbitrary Close() function,
// producing a [io.ReadCloser]. This can be used to implement custom cleanup
// logic in places that expect a [io.ReadCloser] interface.
func ReaderWithCloser(r io.Reader, close func() error) io.ReadCloser {
	return readerWithCloser{
		Reader: r,
		close:  close,
	}
}

// WriterWithCloser wraps a [io.Writer] with an arbitrary Close() function,
// producing a [io.WriteCloser]. This can be used to implement custom cleanup
// logic in places that expect a [io.WriteCloser] interface.
func WriterWithCloser(w io.Writer, close func() error) io.WriteCloser {
	return writerWithCloser{
		Writer: w,
		close:  close,
	}
}
