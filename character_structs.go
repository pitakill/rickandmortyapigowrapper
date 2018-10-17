package rickandmorty

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Origin  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"origin"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
	URL     string   `json:"url"`
	Created string   `json:"created"`
}

type AllCharacters struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}

type MultipleCharacters []Character

func (c Character) GetID() int {
	return c.ID
}

func (ac AllCharacters) GetInfo() Info {
	return ac.Info
}

func (ac AllCharacters) GetResults() []Character {
	return ac.Results
}
