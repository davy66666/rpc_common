package elasticx

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type ElasticConf struct {
	Addr     string
	Username string
	Password string
	Sniff    bool
}

func MustElastic(c ElasticConf) *elastic.Client {
	var options []elastic.ClientOptionFunc
	options = append(options, elastic.SetURL(c.Addr))
	if c.Username != "" && c.Password != "" {
		options = append(options, elastic.SetBasicAuth(c.Username, c.Password))
	}
	options = append(options, elastic.SetSniff(c.Sniff))
	client, err := elastic.NewClient(options...)
	if err != nil {
		panic(err)
	}

	_, _, err = client.Ping(c.Addr).Do(context.Background())
	if err != nil {
		panic(err)
	}

	return client
}
