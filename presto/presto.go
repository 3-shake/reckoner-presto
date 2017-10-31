package presto

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var (
	// state
	FAILED = "FAILED"
)

type Presto struct {
	URL           string
	DefaultHeader http.Header
}

func New(config *Config) *Presto {
	presto := &Presto{
		URL:           config.Host,
		DefaultHeader: make(http.Header, 4),
	}
	presto.SetHeader(config)
	return presto
}

func (this *Presto) SetHeader(config *Config) {
	val := reflect.ValueOf(config).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		key := fmt.Sprintf("X-Presto-%s", convertCamelToHyphen(typeField.Name))
		value := valueField.String()
		if value == "" {
			value = tag.Get("default")
		}
		this.DefaultHeader.Set(key, value)
	}
}

func (this *Presto) Query(query string) (*QueryResult, error) {
	endpoint := fmt.Sprintf("%s/v1/statement", this.URL)
	req, err := this.Request("POST", endpoint, strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	resp, err := Do(req)
	if err != nil {
		return nil, err
	}

	var queryResult QueryResult
	err = json.NewDecoder(resp.Body).Decode(&queryResult)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if queryResult.Stats.State == FAILED {
		return &QueryResult{}, queryResult.Error
	}

	queryResult.Presto = this
	return &queryResult, nil
}

func (this *Presto) Request(method, endpoint string, body io.Reader) (*http.Request, error) {
	parsedUrl, _ := url.Parse(endpoint)
	req, err := http.NewRequest(method, parsedUrl.String(), body)
	if err != nil {
		return nil, err
	}
	for idx, _ := range this.DefaultHeader {
		req.Header[idx] = this.DefaultHeader[idx]
	}
	return req, nil
}

func Do(req *http.Request) (*http.Response, error) {
	for {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		switch resp.StatusCode {
		case 200:
			return resp, nil
		case 503:
			time.Sleep(1 * time.Millisecond)
			continue
		default:
			resp.Body.Close()
			return nil, errors.New("Query Error")
		}
	}
}

func convertCamelToHyphen(s string) string {
	camel := regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "-"))
}
