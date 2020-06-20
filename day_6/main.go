package main

import (
        "github.com/gin-gonic/gin"
        "strings"
        "path/filepath"
        "fmt"
        "os"
)


func main() {
    pwd, _ := os.Getwd()
    router := gin.Default()

    router.GET("/download/:filename", func (ctx *gin.Context) {

            fileName := ctx.Param("filename")
            targetPath := filepath.Join(pwd, fileName)
            fmt.Println(targetPath)
            //This ckeck is for example, I not sure is it can prevent all possible filename attacks - will be much better if real filename will not come from user side. I not even tryed this code
            if !strings.HasPrefix(filepath.Clean(targetPath), pwd) {
                    ctx.String(403, "Look like you attacking me")
                    return
            }
            //Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
            ctx.Header("Content-Description", "File Transfer")
            ctx.Header("Content-Transfer-Encoding", "binary")
            ctx.Header("Content-Disposition", "attachment; filename="+fileName )
            ctx.Header("Content-Type", "application/octet-stream")
            ctx.File(targetPath)
    })

    router.Run()
}
        