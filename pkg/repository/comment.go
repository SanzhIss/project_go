package repository

import (
	"fmt"
	"github.com/Beksultan15/project_go"
	"github.com/jmoiron/sqlx"
)


type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres{
	return &CommentPostgres{db:db}
}


func (a *CommentPostgres) CommentProduct(comment project_go.Comment) (int,error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2,$3,$4) RETURNING id", comment_table)
	row := a.db.QueryRow(query,comment.Product_id,comment.User_id,comment.Rating,comment.Comment)
	
	if err:= row.Scan(&id); err!= nil {
		return 0,err
	}

	return id,nil
}