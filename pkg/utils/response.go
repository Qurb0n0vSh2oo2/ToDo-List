package utils

import (
    "encoding/json"
)

type Response struct {
    Response string
}

func Success()([]byte, error) {
    Success := "Success"
    d := Response{Success}
    return json.MarshalIndent(d, "", "  ")
}

func NotFound()([]byte, error) {
    NotFound := "NotFound"
    d := Response{NotFound}
    return json.MarshalIndent(d, "", "  ")
}

func FillBkanks()([]byte, error) {
    FillBkanks := "Please, fill all blanks"
    d := Response{FillBkanks}
    return json.MarshalIndent(d, "", "  ")
}

func BadRequest()([]byte, error) {
    BadRequest := "Bad Request"
    d := Response{BadRequest}
    return json.MarshalIndent(d, "", "  ")
}

func Exist()([]byte, error) {
    BadRequest := "This accaunt already exist!"
    d := Response{BadRequest}
    return json.MarshalIndent(d, "", "  ")
}
