package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tashima42/textboard/model"
)

func BoardHandler (c echo.Context) error {
  board := c.Param("board")
  // query := r.URL.Query()
  // requestedPage := query.Get("page")

  page := model.BoardPage{
    Boards: []string{ "misc", "news", "discussion", "board" },
    Board: model.Board{
      Ref: board,
      Name: board,
      Posts: []model.Post{
        {
          ID: 1,
          Ref: "1",
          Title: "this is my title",
          Body: "Lorem Ipsum Dolor Sit Amet",
          Date: "26/09/2024 10:21",
        },
        {
          ID: 2,
          Ref: "2",
          Title: "this is my title",
          Body: "Lorem Ipsum Dolor Sit Amet",
          Date: "26/09/2024 10:21",
        },
      },
    },
    Pages: []int{ 1, 2, 3, 4, 5, 6 },
  }

  return c.Render(http.StatusOK, "index", page)
}

