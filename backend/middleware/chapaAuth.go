package middleware

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"local-event-management-backend/helpers/auth"
	"net/http"
	"os"
)

func ChapaAuth(next http.Handler) http.Handler{

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){

			reqBody, err := io.ReadAll(r.Body)
			fmt.Println(string(reqBody))
			if err != nil {
				http.Error(w, "failed to read payload", http.StatusBadRequest)
				return
			}

			keySignature := r.Header.Get("Chapa-Signature")
			payloadkeySignature := r.Header.Get("X-Chapa-Signature")

			if payloadkeySignature == "" || keySignature ==""{
				http.Error(w, "no headers", http.StatusBadRequest)
			}

			// convert the header to hex
			keySignatureHex, err := hex.DecodeString(keySignature)
			fmt.Println("key string",keySignature)
			fmt.Println("key hex",keySignatureHex)

			if err != nil{
				http.Error(w, "failed to read Chapa-Signature", http.StatusBadRequest)
			}

			payloadkeySignatureHex, err:= hex.DecodeString(payloadkeySignature)
			if err != nil{
				http.Error(w, "failed to read X-Chapa-Signature", http.StatusBadRequest)
			}

			// load the key
			key := []byte(os.Getenv("CHAPA_HASH_KEY"))
			if key == nil{
				http.Error(w, "no key", http.StatusInternalServerError)
			}

			// validate the headers
			isChapaKeyValid := auth.ValidMAC(key, keySignatureHex, key)
			fmt.Println("ischapakeyvalid: ",isChapaKeyValid)
			if !isChapaKeyValid{
				http.Error(w, "Unauthorized action", http.StatusUnauthorized)
				return
			}

			isPayLoadValid := auth.ValidMAC(reqBody, payloadkeySignatureHex,key)
			fmt.Println("isPayLoadValid: ",isPayLoadValid)

			if !isPayLoadValid{
				http.Error(w, "Unauthorized action", http.StatusUnauthorized)
				return
			}

			fmt.Println("chapa has been authorized")
			r.Body = io.NopCloser(bytes.NewBuffer(reqBody))
			next.ServeHTTP(w,r)
		},
	)
}