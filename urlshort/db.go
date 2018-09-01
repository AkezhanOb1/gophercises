package urlshort

import (
	"log"
	"github.com/boltdb/bolt"
	"fmt"
	"net/http"
)

var Db *bolt.DB

func init() {
	var err error
	Db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func BoltDbWriter(pathsToUrls map[string]string) {
	if err := Db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		b, err := tx.CreateBucketIfNotExists([]byte("paths"))
		if err != nil {
			return err
		}

		for key, value := range pathsToUrls {
			fmt.Println("Key ", key, " VAlue ", value)
			if err := b.Put([]byte(key), []byte(value)); err != nil {
				return err
			}
		}
		// Set the value "bar" for the key "foo".

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Close database to release file lock.

}


func BoltDbReader (pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	var value []byte
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if err := Db.View(func(tx *bolt.Tx) error {
			value = tx.Bucket([]byte("paths")).Get([]byte(path))
			fmt.Println("Akezhan Path ", path)
			return nil
		}); err != nil {
			panic(err)
		}

		if string(value) != "" {
			http.Redirect(w, r, string(value), http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	})

}