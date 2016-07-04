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

  http.HandleFunc("/", IndexHandler)
  http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
  fp := path.Join("templates", "index.html")
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
