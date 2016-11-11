package bimap

type biMap struct {

	forward map[interface{}]interface{}
	inverse map[interface{}]interface{}
}

func NewBiMap() *biMap {
	return &biMap{forward: make(map[interface{}]interface{}), inverse: make(map[interface{}]interface{})}
}


func (b *biMap) Insert(k interface{}, v interface{}) {
	b.forward[k] = v
	b.inverse[v] = k
}

func (b *biMap) Exists(k interface{}) bool {
	_, ok := b.forward[k]
	return ok
}

func (b *biMap) InverseExists(k interface{}) bool {
	_, ok := b.inverse[k]
	return ok
}

func (b *biMap) Get(k interface{}) (interface{}, bool) {
	if b.Exists(k) {
		return b.forward[k], true
	}
	return "", false
}

func (b *biMap) InverseGet(v interface{}) (interface{}, bool) {
	if b.InverseExists(v) {
		return b.inverse[v], true
	}
	return "", false
}

func (b *biMap) Delete(k interface{}) {
	if b.Exists(k) {
		val, _ := b.Get(k)
		delete(b.forward, k)
		delete(b.inverse, val)
	}
}

func (b *biMap) InverseDelete(v interface{}) {
	if b.InverseExists(v) {
		key, _ := b.InverseGet(v)
		delete(b.inverse, v)
		delete(b.forward, key)
	}
}

func (b*biMap) Size() int {
	return len(b.forward)
}