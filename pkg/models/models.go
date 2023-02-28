package models

type Order struct{
	OrderId uint `json:"orderId" gorm:"primaryKey"`
	CustomerName string `json:"customerName"`
	Items []Item `json:"items" gorm:"foreignKey:OrderId"`
}

type Item struct{
	ItemId uint `json:"itemId" gorm:"primaryKey"`
	ItemCode string `json:"itemCode"`
	Description string `json:"description"`
	Quantity uint `json:"quantity"`
	OrderId uint `json:"-"`
}