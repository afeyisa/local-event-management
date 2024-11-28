package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"local-event-management-backend/types"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func InitiatePayment(price_amount float64, userInfo types.TicketchekoutArgs ) (response  types.Ticketcheckout, err error) {

	url := os.Getenv("CHAPA_URL")
	tx_ref := uuid.New().String()
	payload := map[string]interface{}{
		"amount":   price_amount,
		"currency": "ETB",
		 "first_name":userInfo.First_name,
		 "last_name": userInfo.Last_name,
		 "phone_number": userInfo.Phone_number ,
		 "return_url": os.Getenv("CHAPA_RETURN_URL")+tx_ref,
		 "tx_ref": tx_ref,
		 "email": userInfo.Email,		
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return response , fmt.Errorf("un able to create payment payload")
	}
	chapaHttpClient := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return response , fmt.Errorf("unable prepare payment request")

	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("CHAPA")))
	req.Header.Set("Content-Type", "application/json")

	res, err := chapaHttpClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("chapa API error: %s", err)
	}
	defer res.Body.Close()

	var chapaResponse types.ChapaResponse
	if err := json.NewDecoder(res.Body).Decode(&chapaResponse); err != nil {
		return response, err
	}

	if chapaResponse.Status != "success" {
		return response, fmt.Errorf("chapa API error: %s", chapaResponse.Message)
	}
	
	// pay ment initial insertion
	insertionRes,err := InsertPayments(price_amount, userInfo, tx_ref)

	if err != nil || insertionRes != tx_ref{
		return response, fmt.Errorf("error during making connetion to database")

	}

	response = types.Ticketcheckout{
		Url: chapaResponse.Data.CheckoutURL,
	}

	// save that the payment is started and pending
	return response, nil

}