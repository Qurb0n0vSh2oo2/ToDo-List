package utils

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"regexp"
// 	"topro/pkg/utils"
// )

// func UpdatingUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")
// 	fmt.Println("Updating")
// 	var a int
// 	var account utils.User
// 	byt, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := json.Unmarshal(byt, &account); err != nil {
// 		response, err := utils.BadRequest()
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Fprintf(w, string(response))
// 		return
// 	}

// 	pattern := `^\w`
// 	MatchedName, err := regexp.Match(pattern, []byte(account.Username))
// 	utils.Check(err)

// 	correctId := fmt.Sprintf("select id from users where id = %v", account.ID)
// 	err = dbconn.QueryRow(context.Background(), correctId).Scan(&a)
// 	if err != nil {
// 		response, err := utils.NotFound()
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Fprintf(w, string(response))
// 		fmt.Println(">>>NOT FOUND<<<")
// 		return
// 	}

// 	if MatchedName{
// 		query := fmt.Sprintf("update users set username='%v' where id='%v'",
// 			account.Username, account.ID)

// 		_, err := dbconn.Exec(context.Background(), query)
// 		if err != nil {
// 			response, err := utils.NotFound()
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(w, string(response))
// 			return
// 		} else {
// 			w.WriteHeader(http.StatusOK)
// 			response, err := utils.Success()
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(w, string(response))
// 			fmt.Println("Updated!")
// 		}
// 	} else {
// 		response, err := utils.FillBkanks()
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Fprintf(w, string(response))
// 		fmt.Println("Fill all blanks, please!")
// 	}

// }
