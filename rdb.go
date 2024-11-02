package rdb

// #cgo CFLAGS: -Iinclude
// #cgo LDFLAGS: -Llib -lrdb
// #include "rdb.h"
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Database struct {
	pointer unsafe.Pointer
}

func New(path string) (Database, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	r := C.create(cpath)
	if r.error != nil {
		return Database{}, errors.New(C.GoString(r.error))
	}
	return Database{pointer: r.database}, nil

}

func (db Database) Insert(key string, value []byte) bool {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	ptr := unsafe.Pointer(&value[0])
	len := len(value)
	valueBytes := C.struct_Bytes{
		ptr: (*C.char)(ptr),
		len: C.uint64_t(len),
	}
	return bool(C.insert(db.pointer, ckey, valueBytes))
}

func (db Database) Update(key string, value []byte) bool {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	ptr := unsafe.Pointer(&value[0])
	len := len(value)
	valueBytes := C.struct_Bytes{
		ptr: (*C.char)(ptr),
		len: C.uint64_t(len),
	}
	return bool(C.update(db.pointer, ckey, valueBytes))
}

func (db Database) Search(key string) []byte {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	result := C.search(db.pointer, ckey)
	if result.ptr == nil {
		return nil
	}
	b := C.GoBytes(unsafe.Pointer(result.ptr), C.int(result.len))

	return b
}

func (db Database) Delete(key string) bool {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	return bool(C.delete(db.pointer, ckey))
}
