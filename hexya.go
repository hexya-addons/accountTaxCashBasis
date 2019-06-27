package account_tax_cash_basis

import (
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "account_tax_cash_basis"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
