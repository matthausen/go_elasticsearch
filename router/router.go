package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matthausen/go_elastic/model"
	elastic "github.com/olivere/elastic/v7"
)

// GetESClient - connect to ElasticSearch instance
func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ElasticSearch connection initialised...")
	return client, err
}

// FetchDocuments - get all documents in index
func FetchDocuments(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	var documents []model.Document

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", "John"))

	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := esclient.Search().Index("documents").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var document model.Document
		err := json.Unmarshal(hit.Source, &document)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		documents = append(documents, document)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, d := range documents {
			fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", d.Title, d.Author, d.Content)
		}
	}
}

// NewDocument - generate a new document in index
func NewDocument(w http.ResponseWriter, r *http.Request) {
	// initialise connection with ElaasticSearch index
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initialising: ", err)
		panic("Client fail")
	}

	newDocument := model.Document{
		Title:   "My new document",
		Author:  "John Smith",
		Content: "Amazing content",
	}

	dataJSON, err := json.Marshal(newDocument)
	js := string(dataJSON)
	ind, err := esclient.Index().
		Index("documents").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

}

// Router - the app router
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/fetchDocuments", FetchDocuments).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newDocument", NewDocument).Methods("POST", "OPTIONS")
	return router
}
