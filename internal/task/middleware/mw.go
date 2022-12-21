package middleware

import (
	"errors"
	"log"
	"regexp"
	"topro/pkg/utils"
)

func (t IMwProvider) IMiddleware(password string) (err error) {
	log.Println(2)
	pattern := `^\S`
	MatchedPassword, err := regexp.Match(pattern, []byte(password))
	utils.Check(err)
	
	if MatchedPassword {
		err = t.RMiddleware.RMW(password)
		if err != nil {
			return err
		}
		return
	} else {
		_, err := utils.FillBkanks()
		if err != nil {
			log.Println(err)
		
		}
		return errors.New("Fill all headers, please!")
	}
}
