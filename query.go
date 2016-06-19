package druid

import (
	"encoding/json"
	"time"
)

// Aggregation Queries

const (
	Timeseries = "timeseries"
	TopN       = "topN"
	GroupBy    = "groupBy"
)

// Possible Granularities
// all, none, minute, fifteen_minute, thirty_minute, hour and day.

const (
	GranularityAll           = "all"
	GranularityNone          = "none"
	GranularityMinute        = "minute"
	GranularityFifteenMinute = "fifteen_minute"
	GranularityThirtyMinute  = "thirty_minute"
	GranularityHour          = "hour"
	GranularityDay           = "day"
)

// Filter Types

const (
	FilterSelector   = "selector"
	FilterRegex      = "regex"
	FilterAnd        = "and"
	FilterOr         = "or"
	FilterNot        = "not"
	FilterJavascript = "javascript"
	FilterExtraction = "extraction" //TODO
	FilterSearch     = "search"
	FilterIn         = "in"
	FilterBound      = "bound"
)

// Search Query Types

const (
	SearchInsensitiveContains = "insensitive_contains"
	SearchFragment            = "fragment"
	SearchContains            = "contains"
)

// Aggregator Types

const (
	AggregatorCount = "count"

	AggregatorLongSum   = "longSum"
	AggregatorDoubleSum = "doubleSum"

	AggregatorDoubleMin = "doubleMin"
	AggregatorDoubleMax = "doubleMax"
	AggregatorLongMin   = "longMin"
	AggregatorLongMax   = "longMax"
)

// Post Aggregator
const (
	PostAggregationTypeArithmatic  = "arithmetic"
	PostAggregationTypeFieldAccess = "fieldAccess"
	PostAggregationTypeConstant    = "constant"
)

// Post Aggregator Function Types

const (
	PostAggregatorFnAdd      = "+"
	PostAggregatorFnSubtract = "-"
	PostAggregatorFnMultiply = "*"
	PostAggregatorFnDivide   = "/"
	PostAggregatorFnQuotient = "quotient"
)

// Post Aggregator Field Types
const (
	PostAggregatorFieldFieldAccess = "fieldAccess"
	PostAggregatorFieldConstant    = "constant"
)

type Aggregation struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	FieldName string `json:"fieldName"`
}

type PostAggregatorField struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	FieldName string `json:"fieldName,omitempty"`
	Value     string `json:"value,omitempty"`
}

type PostAggregation struct {
	Type     string                `json:"type"`
	Name     string                `json:"name"`
	Fn       string                `json:"fn"`
	Fields   []PostAggregatorField `json:"fields"`
	Ordering string                `json:"ordering,omitempty"`
}

type SearchQuery struct {
	Type          string   `json:"type"`
	Value         string   `json:"value"`
	Values        []string `json:"values"`
	CaseSensitive bool     `json:"caseSensitive"`
}

type LimitSpec struct {
}

type Having struct {
}

type Filter struct {
	Type      string `json:"type"`
	Dimension string `json:"dimension,omitempty"`
	Value     string `json:"value,omitempty"`

	Fields []Filter `json:"fields,omitempty"`

	// Regex Filter
	Pattern string `json:"pattern,omitempty"`

	// In Filter
	Values []string `json:"values,omitempty"`

	// Javascript Filter
	Function string `json:"function,omitempty"`

	// Bound Filter
	Lower        string `json:"lower,omitempty"`
	Upper        string `json:"upper,omitempty"`
	LowerStrict  bool   `json:"lowerStrict,omitempty"`
	UpperStrict  bool   `json:"upperStrict,omitempty"`
	AlphaNumeric bool   `json:"alphaNumeric,omitempty"`

	// Search Filter
	Query *SearchQuery `json:"query,omitempty"`
}

type AggregationQuery struct {
	QueryType   string  `json:"queryType"`
	DataSource  string  `json:"dataSource"`
	Dimension   string  `json:"dimension,omitempty"`
	Descending  bool    `json:"descending"`
	Threshold   int     `json:"threshold,omitempty"`
	Metric      string  `json:"metric,omitempty"`
	Granularity string  `json:"granularity,omitempty"`
	Filter      *Filter `json:"filter"`

	Aggregations []Aggregation `json:"aggregations"`

	PostAggregations []PostAggregation `json:"postAggregations"`
	Intervals        []string          `json:"intervals"`

	LimitSpec *LimitSpec `json:"limitSpec,omitempty"`
	Having    *Having    `json:"having,omitempty"`
}

type TimeInterval struct {
	Start time.Time
	End   time.Time
}

func TimeseriesQuery(dataSource string, descending bool, granuarity string) *AggregationQuery {
	return &AggregationQuery{
		QueryType:   Timeseries,
		DataSource:  dataSource,
		Descending:  descending,
		Granularity: granuarity,
	}
}

func (q *AggregationQuery) AddInterval(interval string) {
	q.Intervals = append(q.Intervals, interval)
}

func (q *AggregationQuery) SetFilters(f Filter) {
	q.Filter = &f
}

func (q *AggregationQuery) AddAggregator(aggregator Aggregation) {
	q.Aggregations = append(q.Aggregations, aggregator)
}

func (q *AggregationQuery) AddPostAggregator(postAggregator PostAggregation) {
	q.PostAggregations = append(q.PostAggregations, postAggregator)
}

func (q *AggregationQuery) GetJSONString() (string, error) {
	j, err := json.Marshal(q)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
