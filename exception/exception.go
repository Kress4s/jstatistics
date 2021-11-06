package exception

type Type interface {
	Code() int
}

type Exception interface {
	error
	Type() Type
}

type exception struct {
	t Type
	s string
}

func (e exception) Error() string {
	return e.s
}

func (e exception) Type() Type {
	return e.t
}

func New(t Type, text string) Exception {
	return &exception{t, text}
}

func Wrap(t Type, err error) Exception {
	if err == nil {
		return nil
	}
	return &exception{t, err.Error()}
}
