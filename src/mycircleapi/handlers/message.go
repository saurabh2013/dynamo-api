package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"mycircleapi/db"
	model "mycircleapi/models"

	"github.com/gorilla/mux"
)

type MsgController struct{}

func NewMsgController() Handler {
	return new(MsgController)
}

func (this *MsgController) Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := new(db.DB)
	var msg model.Message
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err = json.Unmarshal(b, &msg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if er := data.AddMessage(msg.UserId, &msg); er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (this *MsgController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	email, k := vars["id"]
	if !k {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := new(db.DB)
	msgs, err := data.GetMessage(email)
	if err != nil {
		log.Print(err)
	}
	b, _ := json.MarshalIndent(msgs, "", " ")
	w.Write(b)

}

func (this *MsgController) Delete(w http.ResponseWriter, r *http.Request) {

}
func (this *MsgController) Update(w http.ResponseWriter, r *http.Request) {}
