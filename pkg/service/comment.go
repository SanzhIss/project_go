package service 

import (
	"github.com/Beksultan15/project_go"
	"github.com/Beksultan15/project_go/pkg/repository"
	
)

type CommentService struct {
	repo repository.Commenting
}

func NewCommentService(repo repository.Commenting) *CommentService {
	return &CommentService{repo: repo}
}

func (a *CommentService) CommentProduct(comment project_go.Comment) (int,error) {
	return a.repo.CommentProduct(comment)
}
 