package coinbase

type StartTickerParams struct {
	Type       string    `json:"type"`
	ProductIDs []string  `json:"product_ids"`
	Channels   []Channel `json:"channels"`
}

type Channel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}
