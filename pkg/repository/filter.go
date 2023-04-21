package repository

import (
	"fmt"
	"github.com/Beksultan15/project_go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)




type FilterPostgres struct {
	dba *sqlx.DB
}

func NewFilterPostgres(dba *sqlx.DB) *FilterPostgres{
	return &FilterPostgres{dba:dba}
}


func (a * FilterPostgres) FilteringProduct(c *gin.Context,lte,gte int) ([]project_go.Products,error) {
	dsn := "host=localhost user=docker password=docker dbname=go port=5434 sslmode=disable"
	
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	var products []project_go.Products
		
	fmt.Println(lte)
	sql := "SELECT fp.id,fp.product_title,fp.price,fp.quantity,fp.product_description FROM products as fp "
	
	sql = fmt.Sprintf("%s WHERE fp.price BETWEEN %d AND %d",sql,lte,gte)

	db.Raw(sql).Scan(&products)
	return	products,nil
	
}
