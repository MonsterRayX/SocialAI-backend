package backend

import (
	"around/util"
	"context"

	"cloud.google.com/go/storage"
	"github.com/olivere/elastic/v7"
)

var (
    ESBackend *ElasticsearchBackend
)

type ElasticsearchBackend struct {
    client *elastic.Client
}

func InitGCSBackend(config *util.GCSInfo) {
    client, err := storage.NewClient(context.Background())
    if err != nil {
        panic(err)
    }

    GCSBackend = &GoogleCloudStorageBackend{
        client: client,
        bucket: config.Bucket,
    }
}



func (backend *ElasticsearchBackend) ReadFromES(query elastic.Query, index string) (*elastic.SearchResult, error){
	searchResult, err := backend.client.Search().Index(index).Query(query).Pretty(true).Do(context.Background())
	if err != nil{
		return nil, err
	}
	return searchResult, nil
}

func (backend *ElasticsearchBackend) SaveToES(i interface{}, index string, id string) error{
    _, err := backend.client.Index().Index(index).Id(id).BodyJson(i).Do(context.Background())
    return err
}

func (backend *ElasticsearchBackend) DeleteFromES(query elastic.Query, index string) error {
    _, err := backend.client.DeleteByQuery().
        Index(index).
        Query(query).
        Pretty(true).
        Do(context.Background())
        return err
    }