package main

import (
  "crypto/md5"
  "fmt"
)

func buildComment(model *CommentModel) *CommentEntity {
  var entity CommentEntity
  attr := model.Data.Attributes
  entity.Id = hashId(attr.User, attr.Timestamp)
  entity.User = attr.User
  entity.Timestamp = attr.Timestamp
  entity.Content = attr.Content
  entity.ArticleId = model.Data.Relationships.Article.Data.Id
  return &entity
}

func hashId(user string, timestamp int64) string {
  contentId := fmt.Sprintf("%s:%d", user, timestamp)
  b := md5.Sum([]byte(contentId))
  s := fmt.Sprintf("%x", b)
  return s
}
