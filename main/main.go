package main

import (
  "fmt"
  "github.com/ariauakbar/techtalks/models"
)

func main() {
  u := models.CreateUser("Ariauakbar")
  topics := models.CreateBulkTopics(u)

  fmt.Print(topics)

  // t := models.CreateTopic("Building Real Time Web App with Go", u)
  // fmt.Printf("%s by %s \n", t.Title, t.Creator.Name)
  //
  // t.IncrementVote()
  // t.IncrementVote()
  //
  // fmt.Printf("VoteCount: %d \n", t.VoteCount)
}

// func IndexHandler(w http.ResponseWriter, r *http.Request)  {
//   fp := path.Join("templates", "index.html")
//   tmpl, err := templates.ParseFiles(fp)
//   if err != nil {
//     http.Error(w, err.Error(), http.StatusInternalServerError)
//     return
//   }
//
//
//
// }
