package orderbook

type OrderType string

const (

	GoodTillCancel OrderType="GoodTillCancel"
	FillAndKill OrderType="FillAndKill"
)