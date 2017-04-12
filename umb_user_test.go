package umb_user

import (
	"testing"
)


func TestWorkWithHash(t *testing.T) {

	cases := []struct {
		email, password, hash string
		err                   error
	}{
		{"nikonor@nikonor.ru", "MyTheBestSecretPassword", "", nil},
		{"NiKoNoR@nikonor.ru", "MyTheBestSecretPassword", "", nil},
		{"NiK.oN.oR@nikonor.ru", "MyTheBestSecretPassword", "", nil},
	}

	for i, c := range cases {
		cases[i].hash, cases[i].err = GenHash(c.email, c.password)
		if cases[i].err != nil {
			t.Errorf("Error on GenHash: %s", cases[i].err)
		}
	}

	for _, c := range cases {
		if !CheckPass(c.email, c.password, c.hash) {
			t.Errorf("Error on CheckPass")
		}
	}

}

func TestGenHash(t *testing.T) {
    cases := []struct {
        email, password, hash string
        err                   bool
    }{
        {"", "MyTheBestSecretPassword", "", false},
        {"NiKoNoR@nikonor.ru", "12345", "", false},
        {"NiK.oN.oR@nikonor.ru", "12345678", "", true},
    }
    for _, c := range cases {
        _, err := GenHash(c.email, c.password)
        if ( err != nil && c.err == true ) || ( err == nil && c.err == false ) {
            t.Errorf("Error on GenHash:\n\temail=%s\n\tpass=%s\n\twant=%s\n\tgot=%s\n",c.email,c.password,c.err,err)
        }
    }

}