package controllers

import (
    "io/ioutil"
    "encoding/json"
	"net/http"
	"github.com/samuelhornsey/file-upload/models"
	"strconv"
)

// File Upload handler function
func FileUpload(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Redirect(w, req, "/", http.StatusSeeOther)
        return
	}

	file, handle, err := req.FormFile("file")

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)

    if err != nil {
        panic(err)
	}
	
	f := models.Create(handle.Filename)

	err = ioutil.WriteFile("./uploads/" + f.Hash, data, 0666)
	
    if err != nil {
        panic(err)
    }

	f.Insert()

	data, err = json.Marshal(f)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// File api endpoint
func GetFile(w http.ResponseWriter, req *http.Request) {
	params, ok := req.URL.Query()["ID"]

	if !ok {
		f := models.GetAll()


		data, err := json.Marshal(f)

		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}

	if len(params) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
		return
	}

	id, err := strconv.Atoi(params[0])

	if err != nil {
		panic(err)
	}

	f := models.Get(id)

	data, err := json.Marshal(f)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}