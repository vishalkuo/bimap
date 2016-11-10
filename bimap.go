package bimap

type BiMap struct {
	forward map[string]string
	inverse map[string]string
}

func NewBiMap() *BiMap {
	return &BiMap{forward: make(map[string]string), inverse: make(map[string]string)}
}

func (b *BiMap) Insert(k string, v string) {
	b.forward[k] = v
	b.inverse[v] = k
}

func (b *BiMap) Exists(k string) bool {
	_, ok := b.forward[k]
	return ok
}

func (b *BiMap) InverseExists(k string) bool {
	_, ok := b.inverse[k]
	return ok
}

func (b *BiMap) Get(k string) (string, bool) {
	if b.Exists(k) {
		return b.forward[k], true
	}
	return "", false
}

func (b *BiMap) GetInverse(k string) (string, bool) {
	if b.InverseExists(k) {
		return b.inverse[k], true
	}
	return "", false
}

func (b*BiMap) Size() int {
	return len(b.forward)
}