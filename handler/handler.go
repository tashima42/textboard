package handler

import "github.com/tashima42/textboard/app"

type handler struct {
  boardRepo app.BoardRepository  
  postRepo app.PostRepository
}

func NewHandler() handler {
  return handler{
    boardRepo: app.BoardRepository{},
    postRepo: app.PostRepository{},
  }
}

