package auth

import "golang.org/x/crypto/bcrypt"

//*****************password************************************************************
// hash password for newly joining user
func HashPassword(password string) (string, error){
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

// compare hashed password with plain text password.
// password is plain text from frontend.
// hash is hashed password from hasura.
func CheckPasswordHash(password, hash string) error{
    err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
    return err
}
//********************************************************************************************
