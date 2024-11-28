package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	models "local-event-management-backend/models/payment"
	"net/http"
)

type PaymentWebhook struct {
	// Event         string           `json:"event"`
	// FirstName     string           `json:"first_name"`
	// LastName      string           `json:"last_name"`
	// Email         *string          `json:"email"`
	// Mobile        string           `json:"mobile"`
	// Currency      string           `json:"currency"`
	Amount        string           `json:"amount"`
	Charge 		  string 		   `json:"charge"`
	Status 		  string           `json:"status"`
	// Mode          string           `json:"mode"`
	Reference     string           `json:"reference"`
	// CreatedAt     string           `json:"created_at"`
	// UpdatedAt     string           `json:"updated_at"`
	// Type          string           `json:"type"`
	TxRef         string           `json:"tx_ref"`
	PaymentMethod string           `json:"payment_method"`
	// Customization Customization    `json:"customization"`
	// Meta          interface{}      `json:"meta"`
}

// type Customization struct {
// 	Title       *string `json:"title"`
// 	Description *string `json:"description"`
// 	Logo        *string `json:"logo"`
// }

func ChapaWebHookHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "invalid payload", http.StatusBadRequest)
	}
	var webHookPayload PaymentWebhook 
	err = json.Unmarshal(reqBody, &webHookPayload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid payload json  unmarshal error", http.StatusBadRequest)
		return
	  }
	  // if the payment is on the database
	// get status and check if the status is the same and if not verify the status with chapa verify api
	res, err := models.GetPaymentStatus(webHookPayload.TxRef)

	// if the paymet is not on the database verify the payment with chapa veify api
		// check the similarity of the status
		// if not the retun bad status
		// they are similar  save paymnt method from from webhook and return ok status
	if err != nil || !(len(res) > 0) || res[0].Status != webHookPayload.Status{
		verificationRes,err := models.VerifyPayment(webHookPayload.TxRef)
		if err != nil || verificationRes.Status != webHookPayload.Status{
			http.Error(w, "invalid tx_ref", http.StatusBadRequest)
			return
		}else{
			_ = models.UPdatePaymentMethod(webHookPayload.TxRef, webHookPayload.PaymentMethod)
			w.WriteHeader(http.StatusOK)
			return
		}
	}else{
		_ = models.UPdatePaymentMethod(webHookPayload.TxRef, webHookPayload.PaymentMethod)
		w.WriteHeader(http.StatusOK)
	}
}