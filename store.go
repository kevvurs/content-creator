package main

import (
  "golang.org/x/net/context"
  "google.golang.org/appengine/datastore"
)

func saveComment(ctx context.Context, entity *CommentEntity) error {
  key := datastore.NewKey(ctx, "Comment", entity.Id, 0, nil)
  _, err := datastore.Put(ctx, key, entity)
  return err
}
