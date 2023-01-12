package hashMaps

import "log"

type customHashMap struct {
	Data [][][2]any
}

func NewCustomHashMap(size int) *customHashMap {
	return &customHashMap{Data: make([][][2]any, size)}
}

func (m customHashMap) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % len(m.Data)
	}

	return hash
}

func (m customHashMap) GetKeys() []string {
	var keys []string
	for i := 0; i < len(m.Data); i++ {
		if len(m.Data[i]) != 0 {
			for j := 0; j < len(m.Data[i]); j++ {
				key, ok := m.Data[i][j][0].(string)
				if !ok {
					log.Panic("key is not a string")
				}
				keys = append(keys, key)
			}
		}
	}

	return keys
}

func (m *customHashMap) Set(key string, value int) {
	hash := m.hash(key)

	m.Data[hash] = append(m.Data[hash], [2]any{key, value})
}

func (m *customHashMap) Get(key string) any {
	hash := m.hash(key)
	currentBucket := m.Data[hash]

	if len(currentBucket) != 0 {
		for i := 0; i < len(currentBucket); i++ {
			if currentBucket[i][0] == key {
				return currentBucket[i][1]
			}
		}
	}
	return 0
}
