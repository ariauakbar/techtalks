package main

import (
  //"fmt"
  "github.com/ariauakbar/techtalks/models"
  "net/http"
  "path"
  "html/template"
  "encoding/json"
  "log"
)

func main() {
//  models.CreateTopicTable()

  models.DBInit()

  http.HandleFunc("/", IndexHandler)
  http.HandleFunc("/api/v1/topics", TopicsHandler)
  http.HandleFunc("/api/v1/vote", VoteHandler)
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

    log.Print(t.ID)
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
