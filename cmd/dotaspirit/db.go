package main

import (
	"log"
	"strconv"
	"time"

	"github.com/dgraph-io/badger"
)

func markSent(matchID int64, postID int) {
	err := db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte("vk_posted"+strconv.Itoa(int(matchID))), []byte(strconv.Itoa(postID))).WithTTL(time.Hour * 24)
		err := txn.SetEntry(e)
		return err
	})
	if err != nil {
		log.Fatal("Cant mark as sent", err)
	}
}

func markDone(matchID int64) {
	err := db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte("vk_posted"+strconv.Itoa(int(matchID))), []byte("done")).WithTTL(time.Hour * 24)
		err := txn.SetEntry(e)
		return err
	})
	if err != nil {
		log.Fatal("Cant mark as sent")
	}
}

func getData(matchID int64) string {
	var valCopy []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("vk_posted" + strconv.Itoa(int(matchID))))

		if err != nil {
			log.Println(err)
			return err
		}

		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(valCopy[:])
}
