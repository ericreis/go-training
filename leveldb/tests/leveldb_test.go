package tests

import (
	"encoding/binary"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ericreis/go-training/leveldb/models"
	"github.com/syndtr/goleveldb/leveldb"
)

const count = 1000000

var db *leveldb.DB
var err error

func BenchmarkLevelDbSearchForKey(b *testing.B) {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < b.N; i++ {
		randomKey := make([]byte, 8)
		binary.BigEndian.PutUint64(randomKey, uint64(random.Intn(count)))

		data, err := db.Get(randomKey, nil)
		if err != nil {
			panic(err)
		}
		user := models.User{}
		user.Decode(data)
	}
}

func TestMain(m *testing.M) {
	db, err = leveldb.OpenFile("../db", nil)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	os.Exit(m.Run())
}
