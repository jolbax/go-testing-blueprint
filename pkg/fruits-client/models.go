package fruits_client

type Fruit string

type Fruits struct {
	Fruits []Fruit `json:"fruits"`
}

type FruitsResponse struct {
	Fruits *Fruits
	Code   int
}
