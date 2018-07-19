package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
  "github.com/unrolled/render"
)

// Ping status for Go server
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Ping string }{"OK"})
	}
}

// ETL top news stories
func newCommentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
    ctx := appengine.NewContext(req)
    payload, _ := ioutil.ReadAll(req.Body)
    // parse request
    var model CommentModel
    if err := json.Unmarshal(payload, &model); err != nil {
		  log.Errorf(ctx, "bad request <%v>\n", err)
			formatter.Text(w, http.StatusBadRequest,
			  "Request format could not be read.")  // 400
			return
		}
    // insert data
    entity := buildComment(&model)
		if err := saveComment(ctx, entity); err != nil {
      log.Errorf(ctx, "persistence failed <%v>\n", err)
			formatter.Text(w, http.StatusInternalServerError,
			  "Failed to post comment.")  // 500
			return
    }
    formatter.JSON(w, http.StatusNoContent, nil)  // 204
	}
}
