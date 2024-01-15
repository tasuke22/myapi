package controllers

import (
	"encoding/json"
	"github.com/tasuke/myapi/apperrors"
	"github.com/tasuke/myapi/controllers/services"
	"github.com/tasuke/myapi/models"
	"net/http"
)

type CommentController struct {
	service services.CommentService
}

func NewCommentController(s services.CommentService) *CommentController {
	return &CommentController{
		service: s,
	}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
