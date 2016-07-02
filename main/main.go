package main

import (
  //"fmt"
  "github.com/ariauakbar/techtalks/models"
  "net/http"
  "path"
  "html/template"
)

func main() {
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

  u := models.CreateUser("Ariauakbar")
  topics := models.CreateBulkTopics(u)

  if err := tmpl.Execute(w, topics); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
