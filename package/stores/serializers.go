package stores

import (
	"github.com/steelthedev/go-commerce/package/accounts"
)

type StoresSerializer struct {
	ID    uint                    `json:"id"`
	Name  string                  `json:"store_name"`
	Image string                  `json:"image"`
	User  accounts.UserSerializer `json:"owner"`
}
