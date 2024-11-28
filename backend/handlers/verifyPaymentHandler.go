package handlers

import (
	"encoding/json"
	"io"
	httpHelper "local-event-management-backend/helpers/http"
	models "local-event-management-backend/models/payment"
	"net/http"
)

type verifyPaymentArgs struct{
	Tx_ref string `json:"tx_ref"`
}



type VerifyPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            verifyPaymentArgs `json:"input"`
  }


func VerifyPaymentHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
	  return
	}
  
	var verifyPayload VerifyPayload
	err = json.Unmarshal(reqBody, &verifyPayload)
	if err != nil {
		httpHelper.WriteError(w,http.StatusBadRequest,"invalid payload")
	  return
	}
	defer r.Body.Close()

	result, err := models.VerifyPayment(verifyPayload.Input.Tx_ref)
  
	if err != nil {
	  httpHelper.WriteError(w,http.StatusBadRequest,"failed to verify the payment")

	  return
	}

	json.NewEncoder(w).Encode(result)

}