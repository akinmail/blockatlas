package observer

type Subscription struct {
	Coin     uint     `json:"coin"`
	Address  string   `json:"address"`
	Webhooks []string `json:"webhook"`
}

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64) error
}

type Storage interface {
	Tracker
	Lookup(coin uint, addresses ...string) ([]Subscription, error)
	Add([]Subscription) error
	Delete([]Subscription) error
}
