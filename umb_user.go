package umb_user

import (
	"errors"
	"fmt"
	u "github.com/nikonor/umb_lib"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
	"strings"
)

const (
	MinPasswordLength = 8
)

var ()

type User struct {
	Id    int64
	Email string
}

func GenHash(email, password string) (string, error) {

	if len(email) == 0 && !u.ValidateEmail(email) {
		return "", errors.New("Email is not set or not valid")
	}
	if len(password) < MinPasswordLength {
		return "", errors.New(fmt.Sprintf("Password is too short %d, %d", len(password), MinPasswordLength))
	}

	conf := u.ReadConf("")

	s := make_inner_password(email, password)

	cost, err := strconv.Atoi(conf["BCRYPTCOST"])
	if err != nil {
		cost = bcrypt.DefaultCost
	}

	h, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

func CheckPass(email, password, hash string) bool {
	s := make_inner_password(email, password)
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(s)) == nil {
		return true
	} else {
		return false
	}
}

func make_inner_password(email, password string) string {
	re := regexp.MustCompile("[.@]")
	email = re.ReplaceAllString(email, "")
	s := password + strings.ToLower(email)

	return s
}
