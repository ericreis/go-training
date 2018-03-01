package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/ericreis/go-training/leveldb/models"
	"github.com/syndtr/goleveldb/leveldb"
)

const count = 1000000

func main() {
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		panic(err)
	}

	batch := new(leveldb.Batch)
	for i := uint64(0); i < count; i++ {
		user := models.User{
			ID:       i,
			Name:     "Eric",
			Email:    "eric_reisfig@poli.ufrj.br",
			Password: "pass",
		}

		bytes, err := user.Encode()
		if err != nil {
			panic(err)
		}

		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, user.ID)
		batch.Put(key, bytes)
	}

	err = db.Write(batch, nil)
	if err != nil {
		panic(err)
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	randomKey := make([]byte, 8)
	binary.BigEndian.PutUint64(randomKey, uint64(random.Intn(count)))

	data, err := db.Get(randomKey, nil)
	if err != nil {
		panic(err)
	}
	user := models.User{}
	user.Decode(data)
	fmt.Println(user)

	defer db.Close()
}
