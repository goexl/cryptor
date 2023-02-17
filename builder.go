package cryptor

type builder[T typ] struct {
	from T
}

func newBuilder[T typ](from T) *builder[T] {
	return &builder[T]{
		from: from,
	}
}

func (b *builder[T]) Sha() *shaBuilder {
	return newShaBuilder(b.bytes())
}

func (b *builder[T]) Hmac(key T) *hmacBuilder[T] {
	return newHmacBuilder[T](key, b.bytes())
}

func (b *builder[T]) Md5() *_md5 {
	return newMd5(b.bytes())
}

func (b *builder[T]) bytes() (bytes []byte) {
	switch from := any(b.from).(type) {
	case string:
		bytes = []byte(from)
	case []byte:
		bytes = from
	}

	return
}
