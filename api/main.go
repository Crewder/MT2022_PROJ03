package main

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"log"
	_ "net/http"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()

	// envoie sur le port 9200 l'index
	// res, err := es.Index(
	//	"vroum",                                  // Index name
	//	strings.NewReader(`{"title" : "Test"}`), // Document body
	//	es.Index.WithDocumentID("1"),            // Document ID
	//	es.Index.WithRefresh("true"),            // Refresh
	//)

	// GET
	// res, err := es.Get(
	// 	"vroum",
	// 	"1",
	// )
	// Perform the search request.

	var buf bytes.Buffer
	var key = "abstract"
	var search = "VROUM"

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				key: search,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithAnalyzer(search),
	)

	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}

	}(res.Body)
	log.Println(res)

}
