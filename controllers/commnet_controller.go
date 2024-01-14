package controllers

import (
	"encoding/json"
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
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
