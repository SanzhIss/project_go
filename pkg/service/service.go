package service

import (
	"github.com/Beksultan15/project_go"
	"github.com/Beksultan15/project_go/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(user project_go.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	RefreshToken(username,password string) (string,error)
}

type Searhing interface {
	GetSearchingProduct(c *gin.Context) ([]project_go.Products,error)
}


type Filtering interface{
	FilteringProduct(c *gin.Context,lte,gte int) ([]project_go.Products,error)
}

type Commenting interface{
	CommentProduct(project_go.Comment) (int,error)
}
type Service struct {
	Authorization
	Searhing
	Filtering
	Commenting
}

func NewService(repos *repository.Repository) *Service {
    return &Service{
		Authorization: newAuthService(repos.Authorization),
		Searhing: newSearchService(repos.Searching),
		Filtering: NewFilterService(repos.Filtering),
		Commenting: NewCommentService(repos.Commenting),
	}  
}