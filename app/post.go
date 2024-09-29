package app

import "golang.org/x/net/context"

type PostPage struct {
  Boards []string
  Board string
  Post Post
  Replies []Post
}

type Post struct {
  ID int
  BoardID int
  Ref string
  Title string
  Body string
  Date string
}

var posts = []Post{
  {
    ID: 1,
    BoardID: 1,
    Ref: "1",
    Title: "this is my title",
    Body: "Lorem Ipsum Dolor Sit Amet",
    Date: "26/09/2024 10:21",
  },
  {
    ID: 2,
    BoardID: 1,
    Ref: "2",
    Title: "this is my title",
    Body: "Lorem Ipsum Dolor Sit Amet",
    Date: "26/09/2024 10:21",
  },
  {
    ID: 3,
    BoardID: 2,
    Ref: "3",
    Title: "this is my title",
    Body: "Lorem Ipsum Dolor Sit Amet",
    Date: "26/09/2024 10:21",
  },
}

type PostRepository struct { }

func (p *PostRepository) BoardPosts(ctx context.Context, boardID int) ([]Post, error) {
  boardPosts := make([]Post, 0)
  for _, p := range(posts) {
    if p.BoardID == boardID {
      boardPosts = append(boardPosts, p)
    }
  }
  return boardPosts, nil
}

