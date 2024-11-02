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

func (db Database) Insert(key string, value []byte) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	ptr := unsafe.Pointer(&value[0])
	len := len(value)
	C.insert(db.pointer, ckey, C.struct_Bytes{ptr: (*C.char)(ptr), len: C.uint64_t(len)})

}

func (db Database) Update(key string, value []byte) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	ptr := unsafe.Pointer(&value[0])
	len := len(value)
	C.update(db.pointer, ckey, C.struct_Bytes{ptr: (*C.char)(ptr), len: C.uint64_t(len)})
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

func (db Database) Delete(key string) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	C.delete(db.pointer, ckey)
}
