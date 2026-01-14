package parser

type Parseable[T any] interface {
	Parse(T) error
}

type ParserFn[TInput any, TOutput any] = func(TInput) (TOutput, error)

// Convert generic input to command specific input
func For[TInput any, TOutput Parseable[TInput]](factory func() TOutput) ParserFn[TInput, TOutput] {
	return func(in TInput) (TOutput, error) {
		parseable := factory()
		err := parseable.Parse(in)
		return parseable, err
	}
}

// Don't convert input
func Identity[TInput any](in TInput) (TInput, error) {
	return in, nil
}
