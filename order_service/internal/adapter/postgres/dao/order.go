package dao

type Order struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

func ToOrder() Order {

}
