package cache

// #include "rocksdb/c.h"
// #cgo CFLAGS: -I${SRCDIR}/../rocksdb/include
// #cgo LDFLAGS: -L${SRCDIR}/../rocksdb -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"
import "runtime"

type RocksDBCache struct {
	db *C.rocksdb_t

	readOptions *C.rocksdb_readoptions_t

	writeOptions *C.rocksdb_writeoptions_t

	err *C.char
}

func newRocksDBCache() *RocksDBCache {
	options := C.rocksdb_options_create()

	C.rocksdb_options_increase_parallelism(options, C.int(runtime.NumCPU()))
	C.rocksdb_options_set_create_if_missing(options, 1)

	var err *C.char
	db := C.rocksdb_open(options, C.CString("/mnt/rocksdb"), &err)

	if err != nil {
		panic(C.GoString(err))
	}

	C.rocksdb_options_destroy(options)

	return &RocksDBCache{
		db,
		C.rocksdb_readoptions_create(),
		C.rocksdb_writeoptions_create(),
		err,
	}
}
