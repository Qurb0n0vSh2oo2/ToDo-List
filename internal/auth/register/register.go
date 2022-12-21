package auth

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"topro/pkg/utils"
)

func (t IAuthRegProvider) IRegist(account utils.User) (err error) {
	fmt.Println(2)
	// w.Header().Add("Content-Type", "application/json")
	// var account utils.User
	// byt, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	response, err := utils.BadRequest()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	fmt.Fprintf(w, string(response))

	// 	return
	// }
	// if err := json.Unmarshal(byt, &account); err != nil {
	// 	response, err := utils.BadRequest()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	fmt.Fprintf(w, string(response))
	// 	return
	// }

	pattern := `^\w`
	MatchedName, err := regexp.Match(pattern, []byte(account.Username))
	utils.Check(err)
	MatchedPassword, err := regexp.Match(pattern, []byte(account.Password))
	utils.Check(err)

	if MatchedName && MatchedPassword {
		// 	token, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

		// 	if err != nil {
		// 		log.Println(err)
		// 		return err

		// 	}
		// 	query := fmt.Sprintf(`
		// 	insert into token(token, username)	values('%v', '%v')
		// `, string(token), account.Username)
		// 	_, err = t.Postgres.Exec(context.Background(), query)

		// 	if err != nil {
		// 		_, err := utils.Exist()
		// 		if err != nil {
		// 			log.Println(err)
		// 		}
		// 		return errors.New("Exist!")
		// 	}

		// 	_, err = utils.Success()
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// 	return errors.New("Success")

		// 	tokenUser, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
		// 	if err != nil {
		// 		log.Println(err)
		// 		return err
		// 	}

		// 	query = fmt.Sprintf(`
		// 	insert into users(username, password) values('%v', '%v')
		// 	`, account.Username, string(tokenUser))
		// 	_, err = t.Postgres.Exec(context.Background(), query)
		// 	if err != nil {
		// 		log.Println(err)
		// 		return err
		// 	}
		t.RAuthRegist.RAuthRegist(account)
		return

	} else {
		_, err = utils.FillBkanks()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Fill all blanks, please!")
	}
}
