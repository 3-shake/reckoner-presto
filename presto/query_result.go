package presto

import (
	"encoding/json"
	"errors"
)

type QueryResult struct {
	ID          string         `json:"id"`
	InfoUri     string         `json:"infoUri"`
	NextUri     string         `json:"nextUri"`
	Columns     Columns        `json:"columns"`
	Data        Data           `json:"data"`
	Stats       StatementStats `json:"stats"`
	Error       QueryError     `json:"error"`
	UpdateType  string         `json:"updateType"`
	UpdateCount int            `json:"updateCount"`
	Presto      *Presto        `json:"-"`
}

func (this *QueryResult) Next() (*QueryResult, error) {
	queryResult := &QueryResult{NextUri: this.NextUri}
	var data [][]interface{}
	for queryResult.NextUri != "" {
		req, err := this.Presto.Request("GET", queryResult.NextUri, nil)
		if err != nil {
			return nil, errors.New("invalid request")
		}
		resp, err := Do(req)
		if err != nil {
			return nil, err
		}
		var tmp QueryResult

		d := json.NewDecoder(resp.Body)
		d.UseNumber()
		err = d.Decode(&tmp)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		queryResult = &tmp
		if queryResult.Stats.State == FAILED {
			return &QueryResult{}, queryResult.Error
		}

		if len(queryResult.Data) > 0 {
			data = append(data, queryResult.Data...)
		}
	}
	queryResult.Data = data
	return queryResult, nil
}
