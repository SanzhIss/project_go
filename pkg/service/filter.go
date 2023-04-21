package service 

import (
	"github.com/Beksultan15/project_go"
	"github.com/Beksultan15/project_go/pkg/repository"
	"github.com/gin-gonic/gin"
)

type FilterService struct {
	repo repository.Filtering
}

func NewFilterService(repo repository.Filtering) *FilterService {
	return &FilterService{repo: repo}
}

func (a *FilterService) FilteringProduct(c *gin.Context,lte,gte uint) ([]project_go.Products,error) {
	return a.repo.FilteringProduct(c,lte,gte)
	
}
 