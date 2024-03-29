package constants

const QueryTimeFormat = "2006-01-02_15:04"

// QueryTimeHint is the human-readable explanation of QueryTimeFormat. Intended for use with error reporting.
const QueryTimeHint = "YYYY-MM-DD_hh:mm"

var ValidRelationalOperators = map[string]bool{
	"EQ": true,
	"NE": true,
	"LT": true,
	"GT": true,
	"LE": true,
	"GE": true,
}
