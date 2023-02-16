package cryptor

var _ = New[string]

func New[T typ](from T) *builder[T] {
	return newBuilder[T](from)
}
