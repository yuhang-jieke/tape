/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "115.190.57.118:9000" // 你的 MinIO 地址
		accessKeyID := "minioadmin"       // 管理员账号
		secretAccessKey := "minioadmin"   // 管理员密码
		useSSL := false

		// 初始化客户端
		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			log.Fatalln("连接 MinIO 失败:", err)
		}

		// 2. 准备上传的文件
		filePath := args[0]
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalln("打开文件失败:", err)
		}
		defer file.Close()

		// 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatalln("获取文件信息失败:", err)
		}
		fileName := fileInfo.Name()
		fileSize := fileInfo.Size()

		// 3. 上传到 MinIO 桶（你的桶名是 yuhang）
		bucketName := "yuhang"
		contentType := "image/png" // 可根据实际文件类型修改，或自动识别

		_, err = minioClient.PutObject(
			context.Background(),
			bucketName,
			fileName, // 存储在桶里的文件名
			file,
			fileSize,
			minio.PutObjectOptions{ContentType: contentType},
		)
		if err != nil {
			log.Fatalln("上传失败:", err)
		}

		fmt.Printf("✅ 文件 %s 已成功上传到 Minio 桶 %s\n", fileName, bucketName)
		fmt.Println("upload called")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
