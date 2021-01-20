package geographer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//go:generate go run internal/generator.go

type attributes struct {
	Region      string `json:"aws:region"`
	ServiceName string `json:"aws:serviceName"`
	ServiceURL  string `json:"aws:serviceUrl"`
}

type metadata struct {
	Copyright     string `json:"copyright"`
	Disclaimer    string `json:"disclaimer"`
	FormatVersion string `json:"format:version"`
	SourceVersion string `json:"source:version"`
}

type price struct {
	Attributes *attributes `json:"attributes"`
	ID         string      `json:"id"`
}

type services struct {
	Metadata metadata `json:"metadata"`
	Prices   []price  `json:"prices"`
}

// Services returns a map of service codes to regions (sorted in alphabetical order)
func (r *services) Services() map[string]Regions {

	regions := make(map[string]Regions, len(r.Prices))

	for _, price := range r.Prices {

		var (
			key   = strings.Split(price.ID, ":")[0]
			value = price.Attributes.Region
		)

		regions[key] = append(regions[key], value)

	}

	for _, v := range regions {
		sort.Strings(v)
	}

	return regions

}

// GetServices fetches services metadata from AWS.
func GetServices(ctx context.Context) (*services, error) {

	url := "https://api.regional-table.region-services.aws.a2z.com/index.json"

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to create http client")
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to request region table")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to read region table")
	}

	result := new(services)

	if err := json.Unmarshal(body, result); err != nil {
		return nil, errors.Wrapf(err, "failed to parse region table")
	}

	return result, nil

}
