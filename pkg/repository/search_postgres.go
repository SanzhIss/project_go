package repository

import (
	"fmt"
	"github.com/Beksultan15/project_go"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



type SearchPostgres struct {
	dba *sqlx.DB
}


func NewSearchPostgres(dba *sqlx.DB) *SearchPostgres {
	return &SearchPostgres{dba:dba}
}


func (a *SearchPostgres) GetSearchingProduct(c *gin.Context) ([]project_go.Products, error) {
	dsn := "host=localhost user=docker password=docker dbname=go port=5434 sslmode=disable"
	
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	var products []project_go.Products
		
	if err != nil {
		panic("Could not connect to the database")
	}
	sql := "SELECT fp.id,fp.product_title,fp.price,fp.quantity,fp.product_description FROM products as fp"
		
		if s := c.Query("s"); s != "" {
			sql = fmt.Sprintf("%s WHERE fp.product_title LIKE '%%%s%%'", sql, s)
		}
		db.Raw(sql).Scan(&products)

		return products,nil
}


