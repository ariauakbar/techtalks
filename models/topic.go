package models

// import "github.com/ariauakbar/techtalks/models"
import (
  "fmt"
  "encoding/json"
)

type Topic struct {
  Title string `json:"title"`
  VoteCount int `json:"vote_count"`
  Creator *User `json:"creator"`
}

type Response struct {
  Topics []*Topic `json:"topics"`
}

func CreateTopic(title string, creator *User) *Topic {
  t := &Topic{Title: title, Creator: creator}
  t.VoteCount = 0
  return t
}

func (t *Topic) IncrementVote() {
  t.VoteCount += 1
}

func CreateBulkTopics(creator *User) string {
  const total = 5
  var topics = make([]*Topic, total)

  for i:=0;i < total;i++ {
    title := fmt.Sprintf("Fundamental Go %d", i)
    t := CreateTopic(title, creator)
    topics[i] = t
  }

  return string(jsonFormat(topics))
}

func jsonFormat(topics *Topics) string {
  resp := &Response{
    Topics : topics,
  }
  jsonRsp, _ := json.Marshal(resp)
  return string(jsonRsp)
}
