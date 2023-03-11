package types

import (
	"net/url"
	"strings"

	e "github.com/pergamenum/go-consensus-standards/ehandler"
)

type Query struct {
	Key      string
	Operator string
	Value    any
}

func (q Query) FromURL(input url.Values) ([]Query, error) {

	if len(input) == 0 {
		return []Query{}, nil
	}

	cause := "query must be q=(key),(operator),(value)"
	err := e.Wrap(cause, e.ErrBadRequest)

	// Ensure that any non-conforming query is reported back as invalid.
	qss, found := input["q"]
	if !found || len(input) > 1 {
		return nil, err
	}

	var queries []Query
	for _, qs := range qss {

		split := strings.Split(qs, ",")
		if len(split) != 3 {
			return nil, err
		}

		query := Query{
			Key:      split[0],
			Operator: split[1],
			Value:    split[2],
		}
		queries = append(queries, query)
	}

	return queries, nil
}
