package model

type BoardPage struct {
  Boards []string
  Board Board
  Pages []int
}

type PostPage struct {
  Boards []string
  Board string
  Post Post
  Replies []Post
}

type Board struct {
  ID int
  Ref string
  Name string
  Posts []Post
}

type Post struct {
  ID int
  Ref string
  Title string
  Body string
  Date string
}

