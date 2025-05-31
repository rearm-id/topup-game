package viptopup

// ProfileData defines the structure for profile-related data.
// (Assuming a generic map for now as PHP example doesn't detail it, adjust if known)
type ProfileData map[string]interface{}

// ProfileResponse represents the response from the profile endpoint.
type ProfileResponse struct {
	Status  bool        `json:"status"`
	Data    ProfileData `json:"data"` // Changed from interface{}
	Message string      `json:"message"`
}

// OrderPrepaidData defines the structure for prepaid order data.
// (Assuming a generic map for now as PHP example doesn't detail it for prepaid, adjust if known)
type OrderPrepaidData map[string]interface{}

type OrderPrepaidResponse struct {
	Status  bool             `json:"status"`
	Data    OrderPrepaidData `json:"data"` // Changed from interface{}
	Message string           `json:"message"`
}

// StatusOrderPrepaidDataItem defines individual item structure for prepaid order status.
// (Assuming a generic map for now, adjust if known)
type StatusOrderPrepaidDataItem map[string]interface{}

type StatusOrderPrepaidResponse struct {
	Status  bool                         `json:"status"`
	Data    []StatusOrderPrepaidDataItem `json:"data"` // Changed from interface{}
	Message string                       `json:"message"`
}

// ServicePrepaidDataItem defines individual item structure for prepaid services.
// (Assuming a generic map for now, adjust if known)
type ServicePrepaidDataItem map[string]interface{}

type ServicePrepaidResponse struct {
	Status  bool                     `json:"status"`
	Data    []ServicePrepaidDataItem `json:"data"` // Changed from interface{}
	Message string                   `json:"message"`
}

// OrderGameData defines the structure for the data returned by OrderGame.
type OrderGameData struct {
	TrxID    string `json:"trxid"`
	Code     string `json:"code"`
	DataNo   string `json:"data_no"`
	DataZone string `json:"data_zone,omitempty"`
	Price    int    `json:"price"` // Assuming int, could be float64
	Note     string `json:"note"`
	Balance  int    `json:"balance"` // Assuming int, could be float64
}

type OrderGameResponse struct {
	Status  bool          `json:"status"`
	Data    OrderGameData `json:"data"`
	Message string        `json:"message"`
}

// StatusOrderGameDataItem defines the structure for each item in the StatusOrderGame response data.
type StatusOrderGameDataItem struct {
	TrxID   string `json:"trxid"`
	Code    string `json:"code"`
	DataNo  string `json:"data_no"`
	Price   int    `json:"price"` // Assuming int, could be float64
	Note    string `json:"note"`
	Status  string `json:"status,omitempty"`  // PHP example did not show this, but often present
	Message string `json:"message,omitempty"` // PHP example did not show this, but often present
}

type StatusOrderGameResponse struct {
	Status  bool                      `json:"status"`
	Data    []StatusOrderGameDataItem `json:"data"`
	Message string                    `json:"message"`
}

// ServiceGamePrice defines the nested price structure in ServiceGameDataItem.
type ServiceGamePrice struct {
	Basic   int `json:"basic"`   // Assuming int
	Premium int `json:"premium"` // Assuming int
	Special int `json:"special"` // Assuming int
}

// ServiceGameDataItem defines the structure for each item in the ServiceGame response data.
type ServiceGameDataItem struct {
	Code   string           `json:"code"`
	Game   string           `json:"game"`
	Name   string           `json:"name"`
	Price  ServiceGamePrice `json:"price"`
	Server string           `json:"server"`
	Status string           `json:"status"` // e.g., "available"
	Note   string           `json:"note"`
}

type ServiceGameResponse struct {
	Status  bool                  `json:"status"`
	Data    []ServiceGameDataItem `json:"data"`
	Message string                `json:"message"`
}
