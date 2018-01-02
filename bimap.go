package bimap

import "sync"

type biMap struct {
	s         sync.RWMutex
	immutable bool
	forward   map[interface{}]interface{}
	inverse   map[interface{}]interface{}
}

func NewBiMap() *biMap {
	return &biMap{forward: make(map[interface{}]interface{}), inverse: make(map[interface{}]interface{}), immutable: false}
}

func (b *biMap) Insert(k interface{}, v interface{}) {
	b.s.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.s.RUnlock()

	b.s.Lock()
	defer b.s.Unlock()
	b.forward[k] = v
	b.inverse[v] = k
}

func (b *biMap) Exists(k interface{}) bool {
	b.s.RLock()
	defer b.s.RUnlock()
	_, ok := b.forward[k]
	return ok
}

func (b *biMap) ExistsInverse(k interface{}) bool {
	b.s.RLock()
	defer b.s.RUnlock()

	_, ok := b.inverse[k]
	return ok
}

func (b *biMap) Get(k interface{}) (interface{}, bool) {
	if !b.Exists(k) {
		return "", false
	}
	b.s.RLock()
	defer b.s.RUnlock()
	return b.forward[k], true

}

func (b *biMap) GetInverse(v interface{}) (interface{}, bool) {
	if !b.ExistsInverse(v) {
		return "", false
	}
	b.s.RLock()
	defer b.s.RUnlock()
	return b.inverse[v], true

}

func (b *biMap) Delete(k interface{}) {
	b.s.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.s.RUnlock()

	if !b.Exists(k) {
		return
	}
	val, _ := b.Get(k)
	b.s.Lock()
	defer b.s.Unlock()
	delete(b.forward, k)
	delete(b.inverse, val)
}

func (b *biMap) DeleteInverse(v interface{}) {
	b.s.RLock()
	if b.immutable {
		panic("Cannot modify immutable map")
	}
	b.s.RUnlock()

	if !b.ExistsInverse(v) {
		return
	}

	key, _ := b.GetInverse(v)
	b.s.Lock()
	defer b.s.Unlock()
	delete(b.inverse, v)
	delete(b.forward, key)

}

func (b *biMap) Size() int {
	b.s.RLock()
	defer b.s.RUnlock()
	return len(b.forward)
}

func (b *biMap) MakeImmutable() {
	b.s.Lock()
	defer b.s.Unlock()
	b.immutable = true
}

func (b *biMap) GetInverseMap() map[interface{}]interface{} {
	return b.inverse
}

func (b *biMap) GetForwardMap() map[interface{}]interface{} {
	return b.forward
}

func (b *biMap) Lock() {
	b.s.Lock()
}

func (b *biMap) Unlock() {
	b.s.Unlock()
}
