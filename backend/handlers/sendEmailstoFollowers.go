package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"local-event-management-backend/models/email"
	models "local-event-management-backend/models/organization"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)



type EmailEventPayload struct{
	By_organization_id uuid.UUID `json:"by_organization_id"`
	Description string `json:"description"`
	Event_date string `json:"event_date"`
	Event_id uuid.UUID `json:"event_id"`
	Title string `json:"title"`

}


func mapData (orgname string, event EmailEventPayload) map[string]string{
	data := map[string]string{
	"Organization":orgname,
	"EventTitle": event.Title,
	"EventDate":event.Event_date,
	"EventDescription": event.Description,
	"Id":event.Event_id.String(),
	}

	return data
}

func SendEmalToFollowers(w http.ResponseWriter, r *http.Request){
	reqBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	  }

	var payload EmailEventPayload
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		fmt.Println("json erro ", err)
	http.Error(w, "invalid payload", http.StatusBadRequest)
	return
	}

	// do some thing 
	// get organization followrs
	// map message 
	// send email to them
	if payload.By_organization_id == uuid.Nil{
		fmt.Println("id error")

	http.Error(w, "invalid payload", http.StatusBadRequest)
	return
	}

	data,err :=  models.GetFollower(payload.By_organization_id)
	if err != nil {
		http.Error(w, "un able to get org", http.StatusInternalServerError)
		return
	}

	emailData := mapData(data.Data_organizations[0].Organization_name, payload)
	filePath := "./models/email/emailtemplate.html"
	userid := "me"
	systemEmail := os.Getenv("SYSTEM_EMAIL_ADDRESS")
	subject := "You're Invited"
	for _, user := range data.Data_organizations[0].Follows{
		email.SendEmail(userid,systemEmail,user.User.Email,subject, filePath,emailData)
		time.Sleep(500 * time.Millisecond)
	}

	w.WriteHeader(http.StatusOK)

}