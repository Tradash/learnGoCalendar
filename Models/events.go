package Models

type PlaceInfo string

type Events struct {
	DateStart  string
	DateEnd    string
	Members    []string
	Place      PlaceInfo
	InfoEvents string
	Enabled    bool
}
