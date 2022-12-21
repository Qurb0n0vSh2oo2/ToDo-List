package users

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"topro/pkg/utils"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var DaB = "postgres://postgres:3302@localhost:5432/topro"
var dbconn, _ = pgxpool.Connect(context.Background(), DaB)

func CreatingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var account utils.User
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response, err := utils.BadRequest()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))

		return
	}
	if err := json.Unmarshal(byt, &account); err != nil {
		response, err := utils.BadRequest()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))
		return
	}

	token, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return

	}
	pattern := `^\w`
	MatchedName, err := regexp.Match(pattern, []byte(account.Username))
	utils.Check(err)
	MatchedPassword, err := regexp.Match(pattern, []byte(account.Password))
	utils.Check(err)

	if MatchedName && MatchedPassword {
		query := fmt.Sprintf(`
		insert into token(token, username)	values('%v', '%v')
	`, string(token), account.Username)
		_, err = dbconn.Exec(context.Background(), query)

		if err != nil {
			response, err := utils.Exist()
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Exist")
			fmt.Fprintf(w, string(response))
			return
		}
		response, err := utils.Success()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))
		fmt.Println("DONE!")

		query = fmt.Sprintf(`
		insert into users(username, password) values('%v', '%v')
		`, account.Username, token)
		_, err = dbconn.Exec(context.Background(), query)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		response, err := utils.FillBkanks()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, string(response))
		fmt.Println("Fill all blanks, please!")

	}
}
