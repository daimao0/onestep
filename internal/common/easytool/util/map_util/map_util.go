package map_util

// PutIfAbsent The function is used to put a value into a map if the key does not exist.
func PutIfAbsent[K comparable, V any](hasTable map[K]V, key K, value V) {
	if _, exists := hasTable[key]; !exists {
		hasTable[key] = value
	}
}
