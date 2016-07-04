package models

// import "github.com/ariauakbar/techtalks/models"
import (
  "fmt"
  "github.com/jinzhu/gorm"
  // "encoding/json"
)

type Topic struct {
  gorm.Model
  Title string `json:"title"`
  VoteCount int `gorm:"default:0" json:"vote_count"`
  //Creator *User `json:"creator"`
}

// type Response struct {
//   Topics []*Topic `json:"topics"`
// }

/*
func CreateTopic(title string, creator *User) *Topic {
  t := &Topic{Title: title, Creator: creator}
  t.VoteCount = 0
  return t
}
*/

func CreateTopic(title string) *Topic {
  t := &Topic{Title: title}
  t.VoteCount = 0
  return t
}

func (t *Topic) IncrementVote() {
  t.VoteCount += 1
}

func CreateBulkTopics() []*Topic {
  const total = 5
  var topics = make([]*Topic, total)

  for i:=0;i < total;i++ {
    title := fmt.Sprintf("Fundamental Go %d", i)
    t := CreateTopic(title)
    topics[i] = t
  }

  return topics
}

func CreateTopicTable() {
  db, err := DB()
  if err != nil {
    panic(err.Error())
  }
  db.Debug().AutoMigrate(&Topic{})
  // db.Debug().CreateTable(&Topic{})
  defer db.Close()
}



/*
func CreateBulkTopics(creator *User) []*Topic {
  const total = 5
  var topics = make([]*Topic, total)

  for i:=0;i < total;i++ {
    title := fmt.Sprintf("Fundamental Go %d", i)
    t := CreateTopic(title, creator)
    topics[i] = t
  }

  return topics
}
*/

// func CreateBulkTopics(creator *User) string {
//   const total = 5
//   var topics = make([]*Topic, total)
//
//   for i:=0;i < total;i++ {
//     title := fmt.Sprintf("Fundamental Go %d", i)
//     t := CreateTopic(title, creator)
//     topics[i] = t
//   }
//
//   return string(jsonFormat(topics))
// }

// func jsonFormat(topics []*Topic) []byte {
//   resp := &Response{
//     Topics : topics,
//   }
//   jsonRsp, _ := json.Marshal(resp)
//   return jsonRsp
// }
