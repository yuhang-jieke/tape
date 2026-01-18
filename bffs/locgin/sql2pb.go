package locgin

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	core "github.com/yuhang-jieke/tape/bffs/pkg"
)

func DB2proto(
	dbType string,
	host string,
	user string,
	password string,
	port int,
	schema string,
	table string,
	ignoreTableStr string,
	ignoreColumnStr string,
	serviceName string,
	goPackageName string,
	packageName string,
	fieldStyle string,
	outPath string,
) {
	fmt.Println("接收入参：", dbType, host, user, password, port, schema, table)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, schema)
	db, err := sql.Open(dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var ignoreTables []string
	var ignoreColumns []string
	if ignoreTableStr != "" {
		ignoreTables = strings.Split(ignoreTableStr, ",")
	}
	if ignoreColumnStr != "" {
		ignoreColumns = strings.Split(ignoreColumnStr, ",")
	}
	s, err := core.GenerateSchema(
		db,
		table, // 注意：这里用 table，而不是 schema
		ignoreTables,
		ignoreColumns,
		serviceName,
		goPackageName,
		packageName,
		fieldStyle,
	)
	if err != nil {
		log.Fatal(err)
	}
	if s != nil {
		// 如果指定了输出路径，则写入文件；否则打印到控制台
		if outPath != "" {
			if err = os.WriteFile(outPath, []byte(s.String()), 0644); err != nil {
				log.Fatal(err)
			}
			fmt.Println("proto written to", outPath)
		} else {
			fmt.Println(s)
		}
	}
}
