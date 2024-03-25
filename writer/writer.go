package writer

type Writer interface {
	Open(string) error
	Write(namespace, pod string, log []byte) error
	Close() error
}
