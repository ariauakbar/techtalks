package main

import (
  //"fmt"
  "github.com/ariauakbar/techtalks/models"
  "net/http"
  "path"
  "html/template"
)

func main() {
//  models.CreateTopicTable()

  models.DBInit()

  http.HandleFunc("/", IndexHandler)
  http.HandleFunc("/api/v1/topics", TopicsHandler)
  http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
  fp := path.Join("templates", "indexreact.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  topics := models.GetAllTopics()
  if err := tmpl.Execute(w, topics); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TopicsHandler(w http.ResponseWriter, r *http.Request) {
  resp := models.GetAllTopicsJSON()
  w.Write(resp)
}
