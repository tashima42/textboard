package app

import (
	"context"
	"errors"
)

type Board struct {
  ID int
  Ref string
  Name string
}

type BoardPage struct {
  Boards []string
  Board Board
  Posts []Post
  Pages []int
}

var boards = []Board{
  {
    ID: 1,
    Ref: "misc",
    Name: "Mischelaneous",
  },
  {
    ID: 2,
    Ref: "news",
    Name: "World News",
  },
  {
    ID: 3,
    Ref: "br",
    Name: "Brazil mentioned",
  },
  {
    ID: 3,
    Ref: "us",
    Name: "Eagle",
  },
}

type BoardRepository struct { }

func (b *BoardRepository) BoardNames(ctx context.Context) ([]string, error) {
  boardNames := make([]string, len(boards))
  for i, b := range(boards) {
    boardNames[i] = b.Ref
  }
  return boardNames, nil 
}

func (b *BoardRepository) BoardByRef(ctx context.Context, ref string) (Board, error) {
  for _, b := range boards {
    if ref == b.Ref {
      return b, nil
    }
  }
  return Board{}, errors.New("board not found")
}

