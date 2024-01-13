package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	// いいね数を取得
	article_id := 1
	const sqlGetNice = `
		select nice from articles where article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// いいね数を取得
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// いいね数を更新
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}
	tx.Commit()

	//result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//row := db.QueryRow(sqlStr, articleID)
	//if err := row.Err(); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())

	//var article models.Article
	//var createdTime sql.NullTime
	//
	//err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//if createdTime.Valid {
	//	article.CreatedAt = createdTime.Time
	//}
	//
	////for rows.Next() {
	////	err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	////
	////	if createdTime.Valid {
	////		article.CreatedAt = createdTime.Time
	////	}
	////
	////	if err != nil {
	////		fmt.Println(err)
	////	}
	////}
	//
	//fmt.Printf("%+v\n", article)
}
