package mutations

import "github.com/google/uuid"

type InsetPayment struct {
	Insert_data_payments_one struct {
		Tx_ref uuid.UUID
	} `graphql:"insert_data_payments_one(object: {amount: $amount, currency: $currency, email: $email, first_name: $first_name, last_name: $last_name, paid_for_event_id: $paid_for_event_id, payer_user_id: $payer_user_id, phone_number: $phone_number, tx_ref: $tx_ref, status:$status})"`
}