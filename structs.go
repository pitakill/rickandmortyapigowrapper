package rickandmorty

// General structs

type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type Endpoints struct {
	Characters string `json:"characters"`
	Locations  string `json:"locations"`
	Episodes   string `json:"episodes"`
}

type ResponseItem interface {
	GetID() int
}

type ResponseAll interface {
	GetInfo() Info
}
