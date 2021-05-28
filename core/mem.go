package core

type StoredObject struct {
	key int
}

type ObjectStore struct {
	objs map[int]interface{}
	size int
}

var globalStore *ObjectStore

func init() {
	globalStore = &ObjectStore{
		size: 0,
		objs: make(map[int]interface{}),
	}
}

func AppendObject(obj interface{}) *StoredObject {
	globalStore.objs[globalStore.size] = obj
	so := &StoredObject{
		key: globalStore.size,
	}
	globalStore.size++
	return so
}

func (o StoredObject) Get() interface{} {
	return globalStore.objs[o.key]
}
