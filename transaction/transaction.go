package transaction

import "fmt"

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func (t Transaction) ToString() string {
	return fmt.Sprintf("%s-%s-%f", t.From, t.To, t.Amount)
}
