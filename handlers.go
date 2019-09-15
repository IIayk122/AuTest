package main

import (
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/IIayk122/AuTest/go.uuid"
)

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func play(res http.ResponseWriter, req *http.Request) {
	//Ищем печеньки, если нет - выдаём.
	c, err := req.Cookie("play")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "play",
			Value: sID.String(),
		}
	}
	http.SetCookie(res, c)

	// Продолжить игру если есть печеньки
	if _, ok := data[c.Value]; !ok {
		data[c.Value] = game{generateSecret(), []set{}, false}
	}

	if req.Method == "POST" {
		var newgame game
		if req.FormValue("play_again") == "Играть заново!" || req.FormValue("restart") == "Начать с начала" {
			newgame = game{generateSecret(), []set{}, false}
		} else {
			temp := data[c.Value]
			guess, _ := strconv.Atoi(req.FormValue("guess"))
			newSet := checkGuess(guess, temp.secretNumber)
			if newSet.B == 4 {
				newgame = game{temp.secretNumber, append(temp.Sets, newSet), true}
			} else {
				newgame = game{temp.secretNumber, append(temp.Sets, newSet), false}
			}
		}
		data[c.Value] = newgame
	}
	fmt.Println(data[c.Value])
	tpl.ExecuteTemplate(res, "play.gohtml", data[c.Value])
}
