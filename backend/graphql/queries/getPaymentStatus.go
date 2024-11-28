package queries

type DataPayment []struct{
    Status string `json:"status"`
}
type PaymentStatus struct{
     Data_Payment DataPayment `graphql:"data_payments(where: {tx_ref: {_eq: $tx_ref}})"`
}

   