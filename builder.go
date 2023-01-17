package errs

import "fmt"

// B returns a new error builder.
// Optionally, you can pass an error instance to the builder, ONLY the first error will be used.
//
// # Possible cases:
//
// 1. B() -> new error
//
// 2. B(err) -> update existing error (because of pointers) and returns it after calling Err()
//
// 3. B(err1, err2, err3) -> updates only err1 and returns it after calling Err()
func B(initial ...error) *Builder {
	var err *Error
	if len(initial) <= 0 || initial[0] == nil {
		err = new(Error)
	} else {
		err = Convert(initial[0]).(*Error)
	}
	return &Builder{err}
}

type Builder struct {
	err *Error
}

func (b *Builder) Code(code Code) *Builder {
	b.err.Code = code
	return b
}

func (b *Builder) Msg(msg ...string) *Builder {
	b.err.Msg = append(b.err.Msg, rmNilStr(msg)...)
	return b
}

func rmNilStr(s []string) []string {
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}

func (b *Builder) Msgf(format string, parameters ...any) *Builder {
	b.err.Msg = append(b.err.Msg, fmt.Sprintf(format, parameters...))
	return b
}

func (b *Builder) Details(details ...any) *Builder {
	b.err.Details = details
	return b
}

func (b *Builder) Err() error {
	return b.err
}
