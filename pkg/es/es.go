package es

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/echoframe/pkg/uber"
	"github.com/olivere/elastic"
)

var EsClient *elastic.Client

// NewEs
func NewEs(url string) {
	var err error
	logger := log.New(os.Stdout, "echo-scaffolding", log.LstdFlags)
	EsClient, err = elastic.NewClient(elastic.SetErrorLog(logger),
		elastic.SetSniff(false), elastic.SetURL(url))
	if err != nil {
		panic(err)
	}
	do, i, err := EsClient.Ping(url).Do(context.Background())
	if err != nil {
		panic(err)
	}
	uber.EchoScaLog.Error(fmt.Sprintf("Es result with code %d and version %v\n", i, do.Version))
}
