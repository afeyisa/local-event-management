package mutations

import "github.com/google/uuid"

type UdatePayment struct {
	Update_data_payments_by_pk struct {
		Paid_for_event_id uuid.UUID `json:"paid_for_event_id"`
		Status            string    `json:"status"`
		Payer_user_id     uuid.UUID `json:"payer_user_id"`
		Tx_ref 			  uuid.UUID `json:"tx_ref"`
	} `graphql:"update_data_payments_by_pk(pk_columns: {tx_ref: $tx_ref}, _set: {chapa_reference: $chapa_reference, chapa_charge: $chapa_charge, status: $status})"`
}

type UpdatePyamentMethod struct {
	Update_data_payments_by_pk struct {
		Payment_method string `json:"payment_method"`
	} `graphql:"update_data_payments_by_pk(pk_columns: {tx_ref: $tx_ref}, _set: {payment_method: $payment_method})"`
}
