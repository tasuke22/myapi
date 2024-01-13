package services

import (
	"github.com/tasuke/myapi/models"
	"github.com/tasuke/myapi/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}
