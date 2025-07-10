package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
    bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    return string(bs), err
}

func CheckPassword(hash, pw string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}