package Database

import (
	elastic "github.com/elastic/go-elasticsearch/v8"
	"log"
)

var esClient *elastic.Client

func Init() {
	esClient, _ = elastic.NewDefaultClient()
	log.Println(esClient.Info())
}

func GetESClient() *elastic.Client {
	return esClient
}
