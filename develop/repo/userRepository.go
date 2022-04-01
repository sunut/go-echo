package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	errorHandler "fake.com/develop/error"
	model "fake.com/develop/models"
	util "fake.com/develop/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := viper.GetString("database.dbuser")
	dbPass := viper.GetString("database.dbpassword")
	dbName := "test_go"
	//dbName := "tcp(127.0.0.1:3306)/test_go"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	//db, err := sql.Open(dbDriver, `dbname=test_go host=localhost user=root password=123`)
	if err != nil {
		panic(err.Error())
	}
	boil.SetDB(db)
	//models.Users().Count(context.Background(), db)
	return db
}

func GetAll() ([]model.User, error) {
	db := DbConn()
	defer db.Close()

	//user := model.User{ID: 1113, Name: "supass"}
	//user.Insert(context.Background(), db, boil.Infer())
	result, err := model.Users().All(context.Background(), db)
	if err != nil {
		return nil, errorHandler.HTTPError(http.StatusBadRequest, 1899, "InvalidID", "invalid user id")
	}
	util.ErrorCheck(err)
	jsonData, err := json.Marshal(&result)
	util.ErrorCheck(err)
	fmt.Println(string(jsonData))
	var users []model.User
	err = json.Unmarshal(jsonData, &users)
	return users, nil
}

func GetError() ([]model.User, error) {
	db := DbConn()
	defer db.Close()

	result, err := model.Users().All(context.Background(), db)
	fmt.Println(result)
	fmt.Println(err)
	return nil, errorHandler.HTTPError(http.StatusBadRequest, 1899, "InvalidID", "invalid user id")
}

func AddNewUser(reqUser model.User) {
	db := DbConn()
	defer db.Close()
	//user := model.User{ID: 1113, Name: "supass", Mail: null.String{String:"123"}}
	reqUser.Insert(context.Background(), db, boil.Infer())
}

func UpdateUser(reqUser model.User) {
	db := DbConn()
	defer db.Close()
	user, err := model.FindUser(context.Background(), db, reqUser.ID)
	util.ErrorCheck(err)
	fmt.Println(user)
	user.Name = reqUser.Name
	rows, err := user.Update(context.Background(), db, boil.Infer())
	//rows, err := user.Update(context.Background(), db, boil.Whitelist(model.UserColumns.Name))
	fmt.Println(rows)
}

func DeleteUser(reqUser model.User) {
	db := DbConn()
	defer db.Close()

	tx, err := boil.BeginTx(context.Background(), nil)

	fmt.Println("test: ", err)
	user, err := model.FindUser(context.Background(), tx, reqUser.ID)
	util.ErrorCheck(err)
	rows, err := user.Delete(context.Background(), tx)
	//tx.Commit()
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	// DELETE FROM "pilots" WHERE "id"=$1;
	//err, _ = model.Users(qm.Where("id=?", reqUser.ID)).DeleteAll(context.Background(), db)
	fmt.Println(rows)
}

//func GetAllOlder() []model.User {
//	db := DbConn()
//	defer db.Close()
//	results, e := db.Query("select * from user")
//	util.ErrorCheck(e)
//	var users []model.User
//	for results.Next() {
//		var user model.User
//		// for each row, scan the result into our tag composite object
//		//err := results.Scan(&user.Id, &user.Name, &user.Address, &user.UCode, &user.Mail)
//		//util.ErrorCheck(err)
//		//// and then print out the tag's Name attribute
//		//log.Printf(user.Name)
//		users = append(users, user)
//	}
//
//	return users
//}

//func AddNewUser(reqUser model.User) {
//	db := DbConn()
//	defer db.Close()
//	fmt.Println("done..")
//
//	insert, err := db.Prepare("INSERT INTO user (`id`, `name`, `address`, `ucode`) VALUES (?, ?, ?, ?)")
//	util.ErrorCheck(err)
//	//res, err := insert.Exec(reqUser.Id, reqUser.Name, reqUser.Address, reqUser.UCode)
//	//util.ErrorCheck(err)
//	//id, e := res.LastInsertId()
//	//util.ErrorCheck(e)
//	//fmt.Println("Insert id", id)
//	defer insert.Close()
//}
