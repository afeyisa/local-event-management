package models

import (
	"encoding/json"
	"fmt"
	"io"
	"local-event-management-backend/models/email"
	models "local-event-management-backend/models/ticket"
	"time"

	"net/http"
	"os"
	"strings"
)

type Data  struct {
	// FirstName     string    `json:"first_name"`
	// LastName      string    `json:"last_name"`
	// Email         string    `json:"email"`
	// Currency      string    `json:"currency"`
	// Amount        float64   `json:"amount"`
	Charge        float64   `json:"charge"`
	// Mode          string    `json:"mode"`
	// Method        string    `json:"method"`
	// Type          string    `json:"type"`
	Status        string    `json:"status"`
	Reference     string    `json:"reference"`
	TxRef         string    `json:"tx_ref"`
	// Customization struct {
	// 	Title       string `json:"title"`
	// 	Description string `json:"description"`
	// 	Logo        *string `json:"logo"`
	// } `json:"customization"`
	// Meta      interface{} `json:"meta"`
	// CreatedAt time.Time   `json:"created_at"`
	// UpdatedAt time.Time   `json:"updated_at"`
} 

type PaymentResponse struct {
	// Message string `json:"message"`
	Status  string `json:"status"`
	ChapaData Data `json:"data"`
}

type Response struct{
	Status string `json:"status"`
}

func sendAlert(issue, info string) {
	data := map[string]string{
		"Issue":issue,
		"Time":time.Now().Format("2006-01-02 15:04:05"),
		"Info": fmt.Sprintf("<strong>%v</strong>", info),
	}

		filePath := "./models/email/alert.html"
		userid := "me"
		systemEmail := os.Getenv("SYSTEM_EMAIL_ADDRESS")
		subject := "Server Alert"
		email.SendEmail(userid,systemEmail,systemEmail,subject, filePath,data)
}

func VerifyPayment(tx_ref string) (gqlres Response, err error) {
	fmt.Println("verufy model is called")
	url := os.Getenv("CHAPA_TRANSACTION_VIRIFICATION_URL")+tx_ref
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add( "Authorization", fmt.Sprintf("Bearer %v", os.Getenv("CHAPA")))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var reponseBody PaymentResponse
	err = json.Unmarshal(body, &reponseBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(reponseBody.ChapaData.TxRef)
	paymentInfo, err := UPdatePaymentStatus(reponseBody.ChapaData)

	if err != nil{
		// send alert email
		// send response to user thier ticket is under review
		sendAlert("error occurred during updating payment status", "tx_rf for payment:"+tx_ref)

	
		return Response{Status: "please wait util we review your payment "}, nil

	}
	
	// by checking the status of the payment 
	//  we should be create ticket for the user
	// so we need to call ticket creater here 
	err = models.Ticket(paymentInfo)
	// depending on the error we may alert the admin about the ticket
	// or recording the payment
	if err != nil{
		// send alert email
		// send response to user thier ticket is under review
		sendAlert("error occurred during revoking or awarding ticket", "tx_rf for payment:"+tx_ref)
		return Response{Status: "please wait util we get you ticket "}, nil

	}


	return Response{Status: reponseBody.Status}, nil

} 