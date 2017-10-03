package apirouter

import (
	"../pgmodels"
	"encoding/json"
	"github.com/go-pg/pg"
	"log"
	"net/http"
)

type userInput struct {
	Name     string
	Email    string
	Password *string
}

type userOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	in := userInput{}
	err := decoder.Decode(&in)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
	u := pgmodels.User{
		Name:  in.Name,
		Email: in.Email,
	}
	if in.Password != nil {
		u.SetPassword(*in.Password)
	}

	db := r.Context().Value("db").(*pg.DB)
	err = db.Insert(&u)
	if err != nil {
		panic(err)
	}

	e := json.NewEncoder(w)
	out := userOutput{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
	e.Encode(out)
}
