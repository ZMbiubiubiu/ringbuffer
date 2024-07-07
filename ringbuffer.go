package ringbuffer

type RingBuffer struct {
	buf               []byte
	readIdx, writeIdx int
	len               int
}

func New(length int) *RingBuffer {
	return &RingBuffer{
		len: length,
		buf: make([]byte, length),
	}
}

func (r *RingBuffer) Write(bs []byte) (int, error) {
	return 0, nil
}

func (r *RingBuffer) Read(bs []byte) (int, error) {
	return 0, nil
}
