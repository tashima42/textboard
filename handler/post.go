package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tashima42/textboard/app"
)

func (h *handler) PostHandler (c echo.Context) error {
  board := c.Param("board")
  ref := c.Param("ref")

  page := app.PostPage{
    Boards: []string{ "misc", "news", "discussion", "board" },
    Board: board,
    Post: app.Post{
      ID: 1,
      BoardID: 1,
      Ref: ref,
      Title: "This is my post",
      Body: "Lorem Ipsum Dolor Sit Amet",
      Date: "26/09/2024 10:21",
    },
    Replies: []app.Post{
      {
        ID: 3,
        Ref: "3",
        Title: "this is my other title",
        Body: "Lorem Ipsum Dolor Sit Amet",
        Date: "26/09/2024 10:21",
      },
    },
  }

  return c.Render(http.StatusOK, "post", page)
}

