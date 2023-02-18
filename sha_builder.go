package cryptor

type shaBuilder struct {
	from   []byte
	length int32
}

func newShaBuilder(from []byte) *shaBuilder {
	return &shaBuilder{
		from:   from,
		length: 256,
	}
}

func (sb *shaBuilder) Bit224() *shaBuilder {
	sb.length = 224

	return sb
}

func (sb *shaBuilder) Bit384() *shaBuilder {
	sb.length = 384

	return sb
}

func (sb *shaBuilder) Bit512() *shaBuilder {
	sb.length = 512

	return sb
}

func (sb *shaBuilder) Build() *sha {
	return newSha(sb.from, sb.length)
}
