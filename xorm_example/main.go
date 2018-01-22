// ref: http://xorm.io/docs

package main

import (
	_ "github.com/lib/pq"
	"github.com/go-xorm/xorm"
	"fmt"
	"errors"
	"time"
	"strconv"
)


type UserInfo struct {
	Id   int64 // If field name is Id and type is int64, xorm makes it as auto increment primary key
	Name string
	CityId int64
	UpdatedAt time.Time `xorm:"updated"`
}


type CityInfo struct {
	CityId int64 `xorm:"pk autoincr"`
	CityName string
}

type UserCity struct {
	UserInfo `xorm:"extends"`
	CityName string
}

func (UserCity) TableName() string { // used for join
	return "user_info"
}

func (CityInfo) BeforeInsert(){ // event method on struct
	fmt.Println("event: before data insert in CityInfo")
}

func (CityInfo) AfterInsert(){ // event method on struct
	fmt.Println("event: after data insert in CityInfo")
}

func (CityInfo) BeforeSet(name string, cell xorm.Cell){ // event method on struct
	fmt.Println("event: after data find and before set in CityInfo", name)
}


func showTables(engine *xorm.Engine){
	fmt.Println("\n### showTables")

	tables, err := engine.DBMetas()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, table := range tables {
		fmt.Println(table.Name)
	}

	// engine.IsTableExist(&User{})
	// engine.IsTableEmpty(&User{})
}

func createTables(engine *xorm.Engine){
	fmt.Println("\n### createTables")

	err := engine.CreateTables(&UserInfo{}, &CityInfo{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tables created")
}

func dropTables(engine *xorm.Engine){
	fmt.Println("\n### dropTables")

	err := engine.DropTables(&UserInfo{}, &CityInfo{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("tables deleted")
}

func insertData(engine *xorm.Engine){
	fmt.Println("\n### insertData")

	_, err := engine.Insert(
		&UserInfo{1, "abc", 1, time.Time{}},
		&UserInfo{2, "def", 2, time.Time{}},
		&CityInfo{1, "Dhaka"},
		&CityInfo{2, "Ctg"},
		&CityInfo{3, "Khulna"})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data inserted")
}

func getSingleData(engine *xorm.Engine){
	fmt.Println("\n### getSingleData")

	info := CityInfo{CityName:"Ctg"}
	_, err := engine.Get(&info)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)

	/*
		get value using primary key

		info := CityInfo{}
		_, err := engine.Id(1).Get(&info)
	*/
}

func findMultipleData(engine *xorm.Engine){
	fmt.Println("\n### findMultipleData")


	var cities []CityInfo // using array
	err := engine.Find(&cities)

	// err = engine.SQL("select * from city_info").Find(&cities) //custom sql

	// err = engine.Where("city_id > ? or city_name = ?", 2, "Dhaka").Find(&cities) // conditional

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cities)

	/*
		citiesMap := make(map[int64]CityInfo) // using map
		err = engine.Find(&citiesMap)
		fmt.Println(citiesMap)
	*/
}

func iterateTable(engine *xorm.Engine){
	fmt.Println("\n### iterateTable")

	// Iterate, like find, but handle records one by one

	err := engine.Iterate(new(UserInfo), func(i int, bean interface{}) (err error){

		defer func(){
			if r := recover(); r != nil {
				err = errors.New("error: " + fmt.Sprintf("%v", r))
			}
		}()

		user := bean.(*UserInfo)
		fmt.Println(i, *user)

		//panic("panic")

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func joinTable(engine *xorm.Engine){
	fmt.Println("\n### joinTable")

	var users []UserCity
	err := engine.Join("INNER", "city_info", "user_info.city_id = city_info.city_id").Find(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}

func updateData(engine *xorm.Engine){
	fmt.Println("\n### updateData")

	info := UserInfo{}
	_, err := engine.Id(1).Get(&info)
	if err != nil {
		fmt.Println(err)
		return
	}

	info.Name = "new-name"
	_, err = engine.Id(info.Id).Update(&info)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data updated")
}

func sqlQuery(engine *xorm.Engine){
	fmt.Println("\n### sqlQuery")

	// If select then use Query

	sql := "select * from user_info"
	results, err := engine.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert results []map[string][]byte to users []UserInfo

	var users []UserInfo

	for _,result := range results{
		var user UserInfo

		user.Id, err = strconv.ParseInt(string(result["id"]), 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		user.Name = string(result["name"])

		user.CityId, err = strconv.ParseInt(string(result["city_id"]), 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		// ref: http://stackoverflow.com/questions/25845172/parsing-date-string-in-golang
		// fmt.Println(string(result["updated_at"]))
		layout := "2006-01-02T15:04:05Z"
		user.UpdatedAt, err = time.Parse(layout, string(result["updated_at"]))
		if err != nil {
			fmt.Println(err)
			return
		}

		users = append(users, user)
	}

	fmt.Println(users)
}

func sqlCommand(engine *xorm.Engine) {
	fmt.Println("\n### sqlCommand")

	// If insert, update or delete then use Exec
	sql := "update user_info set name=? where id=?"
	res, err := engine.Exec(sql, "xiaolun", 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	rosAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Rows Affected:", rosAffected)
}

func sessionTransaction(engine *xorm.Engine){
	fmt.Println("\n### sessionTransaction")

	session := engine.NewSession()
	defer session.Close()

	session.Begin()

	_, err := session.Insert(&UserInfo{3, "asd", 3, time.Time{}})
	if err != nil {
		fmt.Println("rollback", err)
		session.Rollback()
		return
	}

	// sql := "update user_info set name=? where id=?"
	sql := "update user_info set username=? where id=?" //column "username" not exists

	_, err = engine.Exec(sql, "pqr", 3)
	if err != nil {
		fmt.Println("rollback", err)
		session.Rollback()
		return
	}

	err = session.Commit()
	if err != nil {
		fmt.Println("commit-error", err)
		return
	}
	fmt.Println("transection commited")
}

func sessionProcessEvents(engine *xorm.Engine){
	fmt.Println("\n### sessionProcessEvents")

	session := engine.NewSession()
	defer session.Close()

	session.Begin()

	before := func(bean interface{}){
		fmt.Println("before", bean)
	}
	after := func(bean interface{}){
		fmt.Println("after", bean)
	}
	session.Before(before).After(after).Insert(&UserInfo{4, "aaa", 3, time.Time{}})
	session.Before(before).After(after).Insert(&UserInfo{5, "abb", 1, time.Time{}})

	fmt.Println("before commit")

	session.Commit()
}


func main() {

	engine, err := xorm.NewEngine("postgres", "host=localhost port=5432 user=postgres password=vagrant dbname=test-dipta")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer  engine.Close()

	// engine.ShowSQL(true)

	dropTables(engine)
	createTables(engine)
	showTables(engine)
	insertData(engine)
	getSingleData(engine)
	findMultipleData(engine)
	iterateTable(engine)
	joinTable(engine)
	updateData(engine)
	sqlQuery(engine)
	sqlCommand(engine)
	sessionTransaction(engine)
	sessionProcessEvents(engine)
}
