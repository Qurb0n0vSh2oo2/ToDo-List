package auth

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"topro/pkg/utils"
)

func (t IAuthLoginProvider) ILogin(account utils.User) (err error) {
	fmt.Println(2)
	pattern := `^\w`
	MatchedName, err := regexp.Match(pattern, []byte(account.Username))
	utils.Check(err)
	MatchedPassword, err := regexp.Match(pattern, []byte(account.Password))
	utils.Check(err)

	if MatchedName && MatchedPassword {

		t.RAuthLogin.RAuthLogin(account)
		return

	} else {
		_, err := utils.FillBkanks()
		if err != nil {
			log.Println(err)
		}
		return errors.New("Fill all blanks, pls!")
	}
}
