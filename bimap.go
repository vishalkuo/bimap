package bimap

type biMap struct {
	forward map[string]string
	inverse map[string]string
}

func NewBiMap() *biMap {
	return &biMap{forward: make(map[string]string), inverse: make(map[string]string)}
}


func (b *biMap) Insert(k string, v string) {
	b.forward[k] = v
	b.inverse[v] = k
}

func (b *biMap) Exists(k string) bool {
	_, ok := b.forward[k]
	return ok
}

func (b *biMap) InverseExists(k string) bool {
	_, ok := b.inverse[k]
	return ok
}

func (b *biMap) Get(k string) (string, bool) {
	if b.Exists(k) {
		return b.forward[k], true
	}
	return "", false
}

func (b *biMap) InverseGet(v string) (string, bool) {
	if b.InverseExists(v) {
		return b.inverse[v], true
	}
	return "", false
}

func (b *biMap) Delete(k string) {
	if b.Exists(k) {
		val, _ := b.Get(k)
		delete(b.forward, k)
		delete(b.inverse, val)
	}
}

func (b *biMap) InverseDelete(v string) {
	if b.InverseExists(v) {
		key, _ := b.InverseGet(v)
		delete(b.inverse, v)
		delete(b.forward, key)
	}
}

func (b*biMap) Size() int {
	return len(b.forward)
}