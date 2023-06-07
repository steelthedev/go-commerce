package shops

type ShopsSerializer struct {
	ID    uint   `json:"id"`
	Name  string `json:"shop_name"`
	Image string `json:"image"`
}
