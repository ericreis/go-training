package tests

import (
	"encoding/binary"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/ericreis/go-training/leveldb/models"
)

const count = 1000000
const bucketName = "demo"

var db *bolt.DB
var err error

func BenchmarkBoltDbSearchForKey(b *testing.B) {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < b.N; i++ {
		randomKey := make([]byte, 8)
		binary.BigEndian.PutUint64(randomKey, uint64(random.Intn(count)))

		user := models.User{}
		var data []byte

		err = db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(bucketName))
			data = bucket.Get(randomKey)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}

		err = user.Decode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestMain(m *testing.M) {
	db, err = bolt.Open("../db/demo.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	os.Exit(m.Run())
}
