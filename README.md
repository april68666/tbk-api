# 淘宝客api基础封装

```
api := NewTbkApi("appKey", "appSecret", "session 如果不调用服务商接口可空")
// 所有订单查询,其他接口自己看吧
orderDetails, err := api.POrderDetailsGet(OrderDetailsGetParams{
	EndTime:    "2021-04-13 11:11:18",
	StartTime:  "2021-04-13 11:11:18",
	QueryType:  1,
	OrderScene: 1,
})
fmt.Println(orderDetails, err)
}
```