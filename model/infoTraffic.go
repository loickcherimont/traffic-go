package model

type InfoTraffic struct {
	Nature      string
	TrafficType string
	Date        string
	Lignes      []Ligne
}
