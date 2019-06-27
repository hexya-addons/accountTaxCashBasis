package account_tax_cash_basis

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AccountMove().DeclareModel()

	h.AccountMove().AddFields(map[string]models.FieldDefinition{
		"TaxCashBasisRecId": models.Many2OneField{
			RelationModel: h.AccountPartialReconcile(),
			String:        "Tax Cash Basis Entry of",
			Help: "Technical field used to keep track of the tax cash basis" +
				"reconciliation.This is needed when cancelling the source:" +
				"it will post the inverse journal entry to cancel that part too.",
		},
	})
	h.AccountMoveLine().DeclareModel()

	h.AccountMoveLine().Methods().Create().Extend(
		`Create`,
		func(rs m.AccountMoveLineSet, vals models.RecordData) {
			//        taxes = False
			//        if vals.get('tax_line_id'):
			//            taxes = [{'use_cash_basis': self.env['account.tax'].browse(
			//                vals['tax_line_id']).use_cash_basis}]
			//        if vals.get('tax_ids'):
			//            taxes = self.env['account.move.line'].resolve_2many_commands(
			//                'tax_ids', vals['tax_ids'])
			//        if taxes and any([tax['use_cash_basis'] for tax in taxes]) and not vals.get('tax_exigible'):
			//            vals['tax_exigible'] = False
			//        return super(AccountMoveLine, self).create(vals)
		})
}
