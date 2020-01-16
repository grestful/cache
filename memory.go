package cache

import "unsafe"

type MemoryCache struct {

}
//
//func runtimeMemory() {
//	runtime.ReadMemStats()
//}


func (mem *MemoryCache) GetSize() {
	return unsafe.Sizeof(mem)
}