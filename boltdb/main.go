package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/boltdb/bolt"

	"github.com/ericreis/go-training/boltdb/models"
)

const count = 1000000
const bucketName = "demo"

func main() {
	db, err := bolt.Open("db/demo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		for i := uint64(0); i < count; i++ {
			user := models.User{
				ID:       i,
				Name:     "Eric",
				Email:    "eric_reisfig@poli.ufrj.br",
				Password: "pass",
			}

			bytes, err := user.Encode()
			if err != nil {
				return err
			}

			key := make([]byte, 8)
			binary.BigEndian.PutUint64(key, user.ID)
			err = bucket.Put(key, bytes)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
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
	fmt.Println(user)

	defer db.Close()
}
