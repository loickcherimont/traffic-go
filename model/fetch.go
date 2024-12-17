package model

// Create a model
// To retrieve JSON data
type Fetch struct {
	Total   int           `json:"total_count"`
	Results []InfoTraffic `json:"results"`
}
