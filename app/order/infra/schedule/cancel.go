package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/order/biz/model"
)

type OrderCancellationJob struct{}

func (j *OrderCancellationJob) Execute() {
	checkAndCancelOrders()
}

func (j *OrderCancellationJob) Description() string {
	return "OrderCancellationJob checks and cancels overdue orders"
}

func checkAndCancelOrders() {
	ctx := context.Background()
	orders, err := model.GetPendingOrders(ctx, mysql.DB)
	if err != nil {
		fmt.Println("Error fetching pending orders:", err)
		return
	}

	for _, order := range orders {
		if time.Since(order.CreatedAt) > 30*time.Minute { // 超过30分钟未支付
			err := model.CanceledByOrderId(ctx, mysql.DB, order.OrderId)
			if err != nil {
				fmt.Println("Error cancelling order:", err)
			} else {
				fmt.Println("Order cancelled:", order.OrderId)
			}
		}
	}
}
