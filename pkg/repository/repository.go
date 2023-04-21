package repository

import (
	"github.com/Beksultan15/project_go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


type Authorization interface {
	CreateUser(user project_go.User) (int, error)
	GetUser(username,password string) (project_go.User, error)
}

type Searching interface {
	GetSearchingProduct(c *gin.Context) ([]project_go.Products, error)
}

type Filtering interface {
	FilteringProduct(c *gin.Context,lte,gte uint) ([]project_go.Products, error)
}

type Repository struct {
	Authorization
	Searching
	Filtering
}

func NewRepository(db *sqlx.DB) *Repository {
    return &Repository{
		Authorization: NewAuthPostgres(db),
		Searching: NewSearchPostgres(db),
		Filtering: NewFilterPostgres(db),
	}
}