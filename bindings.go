package main

// JSON:API model for new comments
type CommentModel struct {
  Data  struct {
    Type string `json:"type"`
    Attributes struct {
      User string `json:"user"`
      Timestamp int64 `json:"timestamp"`
      Content string `json:"content"`
    } `json:"attributes"`
    Relationships struct {
      Article struct {
        Data struct {
          Id string `json:"id"`
          Type string `json:"type"`
        } `json:"data"`
      } `json:"article"`
    } `json:"relationships"`
  } `json:"data"`
}

// Datastore schema for persistence
type CommentEntity struct {
  Id         string  `datastore:"id"`
  User       string  `datastore:"user"`
  Timestamp  int64   `datastore:"timestamp"`
  Content    string  `datastore:"content,noindex"`
  ArticleId  string  `datastore:"articleId"`
}
