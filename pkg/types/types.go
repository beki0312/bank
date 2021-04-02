package types


//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирами и т.д.)
type Money int64

//Category представялет собой категорию, в которой был совершён платёж(авто, аптеки, рестораны и т.д.).
type Category string
//Payment представляет информации о платеже.
type Payment struct{
	ID int
	Amount Money
	Category Category
}

//Currency представляет код валюты
type Currency string

//коды валют
const(
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"

)
//PAN представляет номер карты
type PAN string

//Card представляет  информацию о платёжном карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	MinBalance Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
type PaymentSource struct{
	Type string
	Number  PAN
	Balance  Money
}

