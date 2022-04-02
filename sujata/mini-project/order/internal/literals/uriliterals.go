package literals

const (
	APIVersion1 = "/v1"

	CreateOrderAPIName  = "Create Order"
	CreateOrderEndpoint = APIVersion1 + "/order"

	GetOrderAPIName  = "Get Order"
	GetOrderEndpoint = APIVersion1 + "/order"

	SetOrderStatusAPIName  = "Set status"
	SetOrderStatusEndpoint = APIVersion1 + "/status"
)
