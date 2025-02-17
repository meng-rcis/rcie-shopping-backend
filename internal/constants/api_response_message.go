package constants

const (
	DefaultAPISuccessMsg = "Success"
	DefaultAPIErrorMsg   = "Something went wrong"

	// Sample Service
	GetSampleSuccessMsg       = "Get sample successfully"
	CreateSampleSuccessMsg    = "Create sample successfully"
	UpdateSampleSuccessMsg    = "Update sample successfully"
	EmptySampleDescSuccessMsg = "Empty sample description successfully"
	DeleteSampleSuccessMsg    = "Delete sample successfully"

	// Cart Service
	GetCartItemsSuccessMsg   = "Get cart items successfully"
	AddCartItemSuccessMsg    = "Add item to the cart successfully"
	UpdateCartItemSuccessMsg = "Update cart item successfully"
	RemoveCartItemSuccessMsg = "Delete cart item successfully"

	// Product Service
	GetProductSuccessMsg = "Get product successfully"

	// Order Service
	GetOrdersSuccessMsg   = "Get order successfully"
	CreateOrderSuccessMsg = "Create order successfully"

	// Search Service
	SearchProductSuccessMsg = "Search product successfully"

	// Admin Service
	UpdateOrderStatusSuccessMsg     = "Update order status successfully"
	UpdateProductQuantitySuccessMsg = "Update product quantity successfully"
	BulkUpdateProductSuccessMsg     = "Bulk update product successfully"

	// Common
	DecodingJSONError       = "Decoding JSON Error"
	ConvertStringToIntError = "Convert string to int error"
	MissingParamError       = "Missing parameter"
)
