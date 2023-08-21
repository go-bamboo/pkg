package stdout

type noOutput int

func (*noOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}
