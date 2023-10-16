package Output

type ByteSlice struct {
	Base
	Bytes []byte
}

func (b *ByteSlice) Result() ([]byte, error) {
	return b.Bytes, b.Error
}

func (b *ByteSlice) Err() error {
	return b.Error
}
