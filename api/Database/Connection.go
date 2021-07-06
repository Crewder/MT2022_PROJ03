package Database

import (
	elastic "github.com/elastic/go-elasticsearch/v8"
	"log"
)

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewDefaultClient()
	if err != nil {
		log.Fatal("Error creating the client: %s", err)
	}
	return client, err
}
