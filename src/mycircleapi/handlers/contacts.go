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

type ContactsController struct {
	data *db.Dynamodb
}

func NewContactsController() Handler {
	c := new(ContactsController)
	c.data = new(db.Dynamodb)
	return c
}

func (this *ContactsController) Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var msg model.Contact
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

	if er := this.data.AddContact(msg); er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (this *ContactsController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, k := vars["id"]
	if !k {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	msgs, err := this.data.GetContact(id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, _ := json.MarshalIndent(msgs, "", " ")
	w.Write(b)

}

func (this *ContactsController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, k := vars["id"]
	if !k {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := this.data.Delete(id, db.IdAttribute, db.TblContactList)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		b, _ := json.MarshalIndent("deleted", "", " ")
		w.Write(b)
	}

}

func (this *ContactsController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, k := vars["id"]
	if !k {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var msg model.Affected
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

	err = this.data.UpdateEffected(id, msg.AffectedStatus)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		b, _ := json.MarshalIndent("Updated", "", " ")
		w.Write(b)
	}
}
