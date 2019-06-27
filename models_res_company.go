package account_tax_cash_basis

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.Company().AddFields(map[string]models.FieldDefinition{
		"TaxCashBasisJournalId": models.Many2OneField{
			RelationModel: h.AccountJournal(),
			String:        "Tax Cash Basis Journal",
		},
	})
}
