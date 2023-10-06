package main

const hashTableSize = 10

type HashList struct {
	List map[string]string
}

func New() *HashList {
	s := new(HashList)
	return s
}

func hashFunction(key int) int {
	return key % hashTableSize
}

func (h *HashList) Set(key, value string) error {
	if h.List == nil {
		h.List = make(map[string]string)
	}

	index := hashFunction(key)

	return nil
}

func (h *HashList) Get(key string) (string, error) {
	return "", nil
}
