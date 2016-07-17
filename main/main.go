package main

import (
  //"fmt"
  "github.com/ariauakbar/techtalks/models"
  "net/http"
  "path"
  "html/template"
  "encoding/json"
  //"log"
)

func main() {

  models.DBInit()
  models.CreateUserTable()
  http.HandleFunc("/", IndexHandler)
  http.HandleFunc("/auth", AuthHandler)
  http.HandleFunc("/api/v1/topics", TopicsHandler)
  http.HandleFunc("/api/v1/vote", VoteHandler)
  http.HandleFunc("/api/v1/auth", AuthAPIHandler)
  http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
  fp := path.Join("templates", "indexreact.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  //topics := models.GetAllTopics()
  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    //resp := []string{"Hello": "Hau", "Cool":"yoyoy"}
    var t models.Topic
    err := json.NewDecoder(r.Body).Decode(&t)
    if err != nil {
      panic(err)
    }

    t.IncrementVote()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    resp := models.GetAllTopicsJSON()
    w.Write(resp)
  }
}

func TopicsHandler(w http.ResponseWriter, r *http.Request) {
  resp := models.GetAllTopicsJSON()
  w.Write(resp)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    fp := path.Join("templates", "login.html")
    tmp, err := template.ParseFiles(fp)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    if err := tmp.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

func AuthAPIHandler(w http.ResponseWriter, r *http.Request)  {
  if r.Method == "POST" {
    var u models.User
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
      panic(err.Error())
    }

    if models.CheckIfExisting(&u) == false {
      models.Register(&u)
    } else {
      byt := []byte(`{"status":"Success"}`)
      //resp := json.Marshal(byt)
      w.Write(byt)
    }

  }
}
