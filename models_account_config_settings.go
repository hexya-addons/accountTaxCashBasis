package account_tax_cash_basis

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AccountConfigSettings().DeclareModel()

	h.AccountConfigSettings().AddFields(map[string]models.FieldDefinition{
		"TaxCashBasisJournalId": models.Many2OneField{
			RelationModel: h.AccountJournal(),
			Related:       `CompanyId.TaxCashBasisJournalId`,
			String:        "Tax Cash Basis Journal",
		},
	})
}
