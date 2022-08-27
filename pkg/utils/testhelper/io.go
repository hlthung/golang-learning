package testhelper

type ErrWriter struct {
	err error
}

func (e *ErrWriter) Write(_ []byte) (n int, err error) {
	return 0, e.err
}

func NewErrWriter(err error) *ErrWriter {
	return &ErrWriter{err: err}
}
