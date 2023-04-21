package repository

import (
	"fmt"
	"github.com/Beksultan15/project_go"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




type FilterPostgres struct {
	dba *sqlx.DB
}

func NewFilterPostgres(dba *sqlx.DB) *FilterPostgres{
	return &FilterPostgres{dba:dba}
}


func (a * FilterPostgres) FilteringProduct(c *gin.Context,lte,gte uint) ([]project_go.Products,error) {
	dsn := "host=postgres-db user=docker password=docker dbname=go port=5434 sslmode=disable"
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	var products []project_go.Products
		
	if err != nil {
		panic("Could not connect to the database")
	}
	sql := fmt.Sprintf("SELECT fp.id,fp.producttitle,fp.price,fp.quantity,fp.pddesc, pr.image FROM products AS fp WHERE fp.price BETWEEN '%%%lte%%%' and '%%%gte%%%'",productsTable)
	row := a.dba.Get(&products,sql,lte,gte)
	
	db.Raw(sql).Scan(&products)
	return products,row
	
}
