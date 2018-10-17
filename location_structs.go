package rickandmorty

type Location struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Dimension string   `json:"dimension"`
	Residents []string `json:"residents"`
	URL       string   `json:"url"`
	Created   string   `json:"created"`
}

type AllLocations struct {
	Info    Info       `json:"info"`
	Results []Location `json:"results"`
}

type MultipleLocations []Location

func (l Location) GetID() int {
	return l.ID
}
