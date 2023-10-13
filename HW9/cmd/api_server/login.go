package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func signIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	fmt.Println("debug222")
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	var newAttempt attempt

	err = json.Unmarshal(body, &newAttempt)
	fmt.Printf("newAttempt user = %s\n", newAttempt.Login)
	fmt.Printf("newAttempt password = %s\n", newAttempt.Password)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	originPassword, loginExists := registered[newAttempt.Login]
	if !loginExists {
		http.Error(w, "login doesn't exists", http.StatusBadRequest)
		return
	}

	h := sha1.New()
	h.Write([]byte(newAttempt.Password))
	guessPassword := password(hex.EncodeToString(h.Sum(nil)))
	fmt.Printf("original pass:%s\n", originPassword)
	fmt.Printf("guessPassword pass:%s\n", guessPassword)

	if originPassword != guessPassword {
		http.Error(w, "Password doesn't match", http.StatusBadRequest)
		return
	}

	if password(newAttempt.Role) != "teacher" {
		http.Error(w, "Sorry you aren't teacher", http.StatusBadRequest)
		return
	}

	logined_users = append(logined_users, login(newAttempt.Login))
	fmt.Println("HERE====")
	fmt.Println(logined_users)
	*r = *r.Clone(context.WithValue(r.Context(), CONTEXT_AUTH_KEY, true))
}
