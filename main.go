package main

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	db, err := leveldb.OpenFile("testdb", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 100; i++ {
		err = db.Put([]byte("hoge/"+strconv.Itoa(i)), []byte(strconv.Itoa(i)), nil)
	}

	for j := 0; j < 100; j++ {
		err = db.Put([]byte("fuga/"+strconv.Itoa(j)), []byte(strconv.Itoa(j)), nil)
	}

	for k := 0; k < 100; k++ {
		err = db.Put([]byte("piyo/"+strconv.Itoa(k)), []byte(strconv.Itoa(k)), nil)
	}

	iter := db.NewIterator(util.BytesPrefix([]byte("hoge-")), nil)
	for iter.Next() {
		fmt.Printf("%s: %s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	err = iter.Error()

	err = db.Delete([]byte("hoge-0"), nil)
	err = db.Delete([]byte("fuga-0"), nil)
	err = db.Delete([]byte("piyo-0"), nil)
}
