package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tashima42/textboard/app"
)

func (h *handler) BoardHandler (c echo.Context) error {
  boardRef := c.Param("board")
  // query := r.URL.Query()
  // requestedPage := query.Get("page")

  boards, err := h.boardRepo.BoardNames(c.Request().Context())
  if err != nil {
    return err
  }

  board, err := h.boardRepo.BoardByRef(c.Request().Context(), boardRef)
  if err != nil {
    return err
  }

  posts, err := h.postRepo.BoardPosts(c.Request().Context(), board.ID)
  if err != nil {
    return err
  }

  page := app.BoardPage{
    Boards: boards,
    Board: board,
    Posts: posts,
    Pages: []int{ 1, 2, 3, 4, 5, 6 },
  }

  return c.Render(http.StatusOK, "index", page)
}

