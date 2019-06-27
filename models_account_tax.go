package account_tax_cash_basis

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.AccountTax().DeclareModel()

	h.AccountTax().AddFields(map[string]models.FieldDefinition{
		"UseCashBasis": models.BooleanField{
			String: "Use Cash Basis",
			Help: "Select this if the tax should use cash basis,which will" +
				"create an entry for this tax on a given account during reconciliation",
		},
		"CashBasisAccount": models.Many2OneField{
			RelationModel: h.AccountAccount(),
			String:        "Tax Received Account",
			Filter:        q.Deprecated().Equals(False),
			Help:          "Account use when creating entry for tax cash basis",
		},
	})
}
