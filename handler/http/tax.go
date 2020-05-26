package handler

import (
	std "TAX/CRUD"
	"os"

	"io/ioutil"
	"net/http"
)

func NewOrgHandler() *OrgHandler {
	return &OrgHandler{Repo: std.NewStudents()}
}

// OrgHandler ..
type OrgHandler struct {
	Repo *std.Repo
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (p *OrgHandler) GetClientInfo(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if r.Method != "GET" {

		http.Error(w, http.StatusText(405), 405)

		return
	}
	url := os.Getenv("URL")
	apikey := os.Getenv("API_KEY")
	uri := r.URL.Query().Get("uri")
	newurl := url + uri[1:len(uri)-1] + "?access_token=" + apikey
	println(newurl)

	println(newurl)
	rr, err := http.Get(newurl)
	if err != nil {
		s := "<error>Server not found</error>"
		RespondWithJSON(w, http.StatusNotFound, []byte(s))
	}

	response, _ := ioutil.ReadAll(rr.Body)
	res := string(response)
	println(res)
	RespondWithJSON(w, http.StatusOK, []byte(res))

}
