package modifiers

import(
	"financialExchange/enum"
)
//функция абертка на анонимные функции которая работает с ордерами 
func ModifierForOrder(sellerQuantity float64) func(*enum.Order) *enum.Order{
	return func(order *enum.Order)*enum.Order{
		order.Filled += sellerQuantity

		if order.Filled >= order.Quantity{ 
			order.Status = enum.OrderStatusFilled
		}else{
			order.Status = enum.OrderStatusFilled
		}
		return order
	}
}