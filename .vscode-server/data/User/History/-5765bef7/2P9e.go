package backend

import(
	"context"
	"fmt"

	"aroundconstants"

	"github.com/olivere/elastic/v7"
)

var (
	ESBackend *ElasticsearchBackend
)

type ElasticsearchBackend struct{
	client *elas
}