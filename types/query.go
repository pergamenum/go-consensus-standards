package types

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	c "github.com/pergamenum/go-consensus-standards/constants"
	e "github.com/pergamenum/go-consensus-standards/ehandler"
)

type Query struct {
	Key      string
	Operator string
	Value    any
}

func (q *Query) Validate(ttt map[string]string, otb map[string]bool) error {

	if q == nil {
		return fmt.Errorf("query was nil")
	}

	var sb strings.Builder
	// Validate Key and Value.
	if t, found := ttt[q.Key]; !found {
		var vks []string
		for vk := range ttt {
			vks = append(vks, vk)
		}
		sb.WriteString(fmt.Sprintf("key(%v): invalid ", q.Key))
		sb.WriteString(fmt.Sprintf("- valid keys are: (%v) ", strings.Join(vks, " ")))
	} else {
		actuallyString := t == "string"
		_, foundString := q.Value.(string)
		// When the 'Value any' field holds a string representation of another type.
		if foundString && !actuallyString {
			err := q.setValueFromString(t)
			if err != nil {
				sb.WriteString(err.Error())
			}
		} else {
			err := AssertAny(&q.Value, t)
			if err != nil {
				sb.WriteString(err.Error())
			}
		}
	}
	// Validate Operator.
	if valid, found := otb[q.Operator]; !found || !valid {
		var vos []string
		for vo, v := range otb {
			if v {
				vos = append(vos, vo)
			}
		}
		sb.WriteString(fmt.Sprintf("operator(%v): invalid ", q.Operator))
		sb.WriteString(fmt.Sprintf("- valid operators are: (%v) ", strings.Join(vos, " ")))
	}

	if len(sb.String()) > 0 {
		cause := fmt.Sprint("invalid query: ", strings.TrimSpace(sb.String()))
		return fmt.Errorf(cause)
	}

	return nil
}

func (q *Query) setValueFromString(t string) error {

	var sv string
	if s, ok := q.Value.(string); !ok {
		return fmt.Errorf("value is not of type string")
	} else {
		sv = s
	}
	if sv == "" {
		return fmt.Errorf("value is empty")
	}
	if t == "" {
		return fmt.Errorf("type parameter is empty")
	}

	f := func(result any, err error) error {
		if err != nil {
			return fmt.Errorf("value(%v) is not a valid %s, info: (%s)", q.Value, t, err.Error())
		}
		q.Value = result
		return nil
	}

	switch strings.ToLower(t) {

	case "bool":
		result, err := strconv.ParseBool(sv)
		return f(result, err)

	case "int":
		result, err := strconv.ParseInt(sv, 10, 0)
		convert := int(result)
		return f(convert, err)

	case "int8":
		result, err := strconv.ParseInt(sv, 10, 8)
		convert := int8(result)
		return f(convert, err)

	case "int16":
		result, err := strconv.ParseInt(sv, 10, 16)
		convert := int16(result)
		return f(convert, err)

	case "int32":
		result, err := strconv.ParseInt(sv, 10, 32)
		convert := int32(result)
		return f(convert, err)

	case "int64":
		result, err := strconv.ParseInt(sv, 10, 64)
		return f(result, err)

	case "uint":
		result, err := strconv.ParseUint(sv, 10, 0)
		convert := uint(result)
		return f(convert, err)

	case "uint8":
		result, err := strconv.ParseUint(sv, 10, 8)
		convert := uint8(result)
		return f(convert, err)

	case "uint16":
		result, err := strconv.ParseUint(sv, 10, 16)
		convert := uint16(result)
		return f(convert, err)

	case "uint32":
		result, err := strconv.ParseUint(sv, 10, 32)
		convert := uint32(result)
		return f(convert, err)

	case "uint64":
		result, err := strconv.ParseUint(sv, 10, 64)
		return f(result, err)

	case "float32":
		result, err := strconv.ParseFloat(sv, 32)
		convert := float32(result)
		return f(convert, err)

	case "float64":
		result, err := strconv.ParseFloat(sv, 64)
		return f(result, err)

	case "complex64":
		result, err := strconv.ParseComplex(sv, 64)
		convert := complex64(result)
		return f(convert, err)

	case "complex128":
		result, err := strconv.ParseComplex(sv, 128)
		return f(result, err)

	case "time":
		result, err := time.Parse(c.QueryTimeFormat, sv)
		if err != nil {
			cause := fmt.Errorf("valid form is: %s", c.QueryTimeHint)
			err = e.Wrap(cause, err)
		}
		return f(result, err)

	default:
		return fmt.Errorf("unsupported type(%s) with value(%v)", t, q.Value)
	}
}

func (q *Query) FromURL(input url.Values) ([]Query, error) {

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
