package tasks

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"topro/pkg/utils"
// )

// func ReadingTask(w http.ResponseWriter, r *http.Request) {


// 	w.Header().Add("Content-Type", "application/json")
// 	var allrows []utils.Task
// 	var checkbody utils.Task
// 	byt, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		response, err := utils.BadRequest()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		fmt.Fprintf(w, string(response))
// 		return
// 	}
// 	if err := json.Unmarshal(byt, &checkbody); err != nil {
// 		response, err := utils.BadRequest()
// 		if err != nil {
// 			log.Println(log.Llongfile, err)
// 		}

// 		fmt.Fprintf(w, string(response))
// 		return
// 	}

// 	query := fmt.Sprintf("select id, name, description, status, user_id from tasks where user_id=%v", checkbody.UserID)
// 	// _, err := dbconn.Exec(context.Background(), query)
// 	rows, err := dbconn.Query(context.Background(), query)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	} else if !rows.Next() {
// 		w.WriteHeader(http.StatusNotFound)
// 		response, err := utils.NotFound()
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		fmt.Fprintf(w, string(response))
// 		return
// 	}

// 	// fmt.Println("rows.Next():", rows.Next())
// 	// fmt.Println("err:", err)
// 	// fmt.Println("Этап проверки пройден")

// 	firsttask := utils.Task{}
// 	err = rows.Scan(&firsttask.ID, &firsttask.Name, &firsttask.Description, &firsttask.Status, &firsttask.UserID)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	allrows = append(allrows, firsttask)

// 	for rows.Next() {
// 		saverows := utils.Task{}
// 		err = rows.Scan(&saverows.ID, &saverows.Name, &saverows.Description, &saverows.Status, &saverows.UserID)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		allrows = append(allrows, saverows)
// 	}

// 	fmt.Println("allrows:", allrows)

// 	w.WriteHeader(http.StatusOK)
// 	response, err := utils.Success()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	fmt.Fprintf(w, string(response))
// 	fmt.Println("DONE!")
// }
