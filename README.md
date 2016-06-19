# Druid Client for Go
Currently Supports Timeseries and TopNQuery.

## Example Usage

### TopN query

```
client := druid.New("HOST_URL")
countryFilter := druid.Filter{Type: druid.FilterSelector, Dimension: "country_name", Value: "United States"}
country2Filter := druid.Filter{Type: druid.FilterSelector, Dimension: "country_name", Value: "Australia"}
filter := druid.Filter{Type: druid.FilterOr, Fields: []druid.Filter{countryFilter, country2Filter}}
q := druid.TopNQuery("wikipedia", "city", "count", 20, druid.GranularityThirtyMinute)
q.SetFilters(filter)
q.AddAggregator(druid.Aggregation{Type: druid.AggregatorLongSum, Name: "total_count", FieldName: "count"})
q.AddAggregator(druid.Aggregation{Type: druid.AggregatorCount, Name: "count"})
q.AddInterval("2013-08-01T00:00:00.000/2013-08-08T21:22:48.989")
data, err := client.RunQuery(q)
```