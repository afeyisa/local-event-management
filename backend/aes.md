// package main

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"encoding/hex"
// 	"io"
// 	"log"
// 	"os"
// )


// func SaveHex(fileName string, content string){

// 	file, err := os.Create(fileName)
// 	if err != nil {
//         panic(err)
//     }
// 	_,err = file.WriteString(content)
// 	if err != nil {
// 	panic(err)
// 	}
// }



// func GenerateKey() []byte {
// 	key := make([]byte, 32)
// 	_, err := rand.Reader.Read(key)
// 	if err != nil{
// 		log.Fatal("un able to generate key")
// 	}
// 	return key
// }

// func Encrypt( keyname , filePath string){
// 	plaintext, err := os.ReadFile(filePath)
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	key := GenerateKey()
// 	sk := hex.EncodeToString(key)
// 	SaveHex(keyname+".txt", sk)
// 	ciph, err := aes.NewCipher(key)
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	gcm, err := cipher.NewGCM(ciph)
// 	if err != nil {
//         log.Fatal(err)
//     }

// 	nonce := make([]byte, gcm.NonceSize())
// 	_, err = io.ReadFull(rand.Reader, nonce)
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	os.WriteFile(keyname+"cipher.txt", (gcm.Seal(nonce, nonce, plaintext, nil)),0600)
// }

// func Decrypt(filepath, key string){
// 	ciphher, err := os.ReadFile(filepath)
// 	if err != nil {
//         log.Fatal(err)
//     }

// 	bytekey,err := hex.DecodeString(key)
// 	if err != nil {
//         log.Fatal(err)
//     }

// 	ciph, err := aes.NewCipher(bytekey)
// 	if err != nil {
//         log.Fatal(err)
//     }

// 	gcm, err := cipher.NewGCM(ciph)
// 	if err != nil {
//         log.Fatal(err)
//     }

// 	nonceSize := gcm.NonceSize()

// 	if len(ciphher) < nonceSize {
//         log.Fatal(err)
//     }

// 	nonce, ciphertext := ciphher[:nonceSize], ciphher[nonceSize:]
//     plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
//     if err != nil {
//         log.Fatal(err)
//     }

// 	os.WriteFile("plaintext.json", plaintext,0777)

// }



// func main(){

// 	// Encrypt("credentialskey","credentials copy.json")
// 	Decrypt("credentialskeycipher.txt","56fa7b6e83c091cf96d051cad7ca730687afa784a2ae826bfbbe1380c77fae75")

// }
// // generate key
// // save key
// // encrypt data
// // save encrpted data