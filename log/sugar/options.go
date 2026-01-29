package sugar

// Options is Log Options.
type Options struct {
	ver  string
	skip int
}

// Option is Log option.
type Option func(*Options)

// WithVersion with filter options
func WithVersion(ver string) Option {
	return func(opts *Options) {
		opts.ver = ver
	}
}

func WithSkip(skip int) Option {
	return func(opts *Options) {
		opts.skip = skip
	}
}
