/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuhang-jieke/tape/bffs/locgin"
)

// sql2pbCmd represents the sql2pb command
var pb2protoCmd = &cobra.Command{
	Use:   "sql2pb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		locgin.DB2proto(
			dbType,
			host,
			user,
			password,
			port,
			schema,
			table,
			ignoreTableStr,
			ignoreColumnStr,
			serviceName,
			goPackageName,
			packageName,
			fieldStyle,
			outPath,
		)

		fmt.Println("sql2pb called")
	},
}

var dbType, host, user, password, schema, table, ignoreTableStr, ignoreColumnStr string
var serviceName, goPackageName, packageName, fieldStyle string
var outPath string
var port int

func init() {
	rootCmd.AddCommand(pb2protoCmd)
	pb2protoCmd.Flags().StringVarP(&dbType, "db", "D", "mysql", "the database type")
	pb2protoCmd.Flags().StringVarP(&host, "host", "H", "localhost", "the database host")
	pb2protoCmd.Flags().IntVarP(&port, "port", "P", 3306, "the database port")
	pb2protoCmd.Flags().StringVarP(&user, "user", "U", "root", "the database user")
	pb2protoCmd.Flags().StringVarP(&password, "password", "p", "root", "the database password")
	pb2protoCmd.Flags().StringVarP(&schema, "schema", "S", user, "the database schema")
	pb2protoCmd.Flags().StringVarP(&table, "table", "T", "2307aim", "the database table")

	pb2protoCmd.Flags().StringVarP(&ignoreTableStr, "ignore-tables", "I", "", "a comma spaced list of tables to ignore")
	pb2protoCmd.Flags().StringVarP(&ignoreColumnStr, "ignore-columns", "C", "", "a comma spaced list of mysql columns to ignore")
	pb2protoCmd.Flags().StringVar(&serviceName, "service", "MyService", "the grpc service name")
	pb2protoCmd.Flags().StringVar(&goPackageName, "go_package", "", "the go_package option in proto")
	pb2protoCmd.Flags().StringVar(&packageName, "package", "pb", "the proto package name")
	pb2protoCmd.Flags().StringVar(&fieldStyle, "field_style", "", "field naming style (e.g. snake, camel)")
	pb2protoCmd.Flags().StringVar(&outPath, "out", "", "output .proto file path (if empty, print to stdout)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pb2protoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pb2protoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
