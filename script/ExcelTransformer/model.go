package main

import "encoding/json"

type Product struct {
	Code          string
	Name          string
	BuyPrice      int64
	SellPrice     int64
	Quantity      int32
	CategoryID    int64
	Category      Category
	UnitID        int64
	Unit          Unit
	ImageUrl      string
	Discount      float32
	ProductPrices ProductPrices
}

type ProductPrice struct {
	QuantityMultiplier int32
	SellPrice          int64
}

type ProductPrices []ProductPrice

func (p ProductPrices) ToString() string {
	maps := map[int32]int64{}
	for _, val := range p {
		maps[val.QuantityMultiplier] = val.SellPrice
	}
	jsonString, _ := json.Marshal(maps)
	return string(jsonString)
}

type Unit struct {
	Name     string
	Code     string
	TotalPcs int32
}

type Category struct {
	Name string
}
