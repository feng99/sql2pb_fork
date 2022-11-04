package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Mikaelemmmm/sql2pb/core"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbType := flag.String("db", "mysql", "the database type")
	host := flag.String("host", "rm-2ze1gfhk0hmq7l3ov.mysql.rds.aliyuncs.com", "the database host")
	port := flag.Int("port", 3306, "the database port")
	user := flag.String("user", "youth_mall", "the database user")
	password := flag.String("password", "A8Zh_7pQG_eaiLD-pWLc77t", "the database password")
	//database
	schema := flag.String("schema", "crmeb", "the database schema")
	//table
	table := flag.String("table", "eb_live_room", "the table schema，multiple tables ',' split. ")
	serviceName := flag.String("service_name", *schema, "the protobuf service name , defaults to the database schema.")
	packageName := flag.String("package", *schema, "the protocol buffer package. defaults to the database schema.")
	goPackageName := flag.String("go_package", "", "the protocol buffer go_package. defaults to the database schema.")
	ignoreTableStr := flag.String("ignore_tables", "", "a comma spaced list of tables to ignore")

	flag.Parse()

	if *schema == "" {
		fmt.Println(" - please input the database schema ")
		return
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *user, *password, *host, *port, *schema)
	db, err := sql.Open(*dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ignoreTables := strings.Split(*ignoreTableStr, ",")

	s, err := core.GenerateSchema(db, *table, ignoreTables, *serviceName, *goPackageName, *packageName)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}
