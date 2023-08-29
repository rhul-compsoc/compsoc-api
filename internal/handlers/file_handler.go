package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
)

func ListFile() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /file/:file
func GetFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("file")

		path := path.Join("uploads/", n)
		fmt.Println(path)

		e, err := checkExists(path)
		if err != nil {
			c.Status(http.StatusBadRequest)
			log.Println(err, 1)
			return
		}
		if !e {
			c.Status(http.StatusBadRequest)
			log.Println(util.ErrFileNotExist, 2)
			return
		}

		f, err := os.Open(path)
		if err != nil {
			c.Status(http.StatusBadRequest)
			log.Println(err, 3)
			return
		}
		defer f.Close()

		h := make([]byte, 512)
		_, err = f.Read(h)
		if err != nil {
			c.Status(http.StatusBadRequest)
			log.Println(err, 4)
			return
		}

		t := http.DetectContentType(h)
		i, err := f.Stat()
		if err != nil {
			c.Status(http.StatusBadRequest)
			log.Println(err, 5)
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", n))
		c.Header("Content-Type", t)
		c.Header("Content-Length", fmt.Sprintf("%d", i.Size()))
		c.File(path)
	}
}

// Upload file controller
func PostFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		f, err := c.FormFile("file")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		path := path.Join("uploads", f.Filename)
		fmt.Println(path)
		e, err := checkExists(path)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if e {
			c.Status(http.StatusBadRequest)
			return
		}

		err = c.SaveUploadedFile(f, path)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

func PostFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		files := form.File["files"]
		for _, f := range files {
			path := path.Join("uploads", f.Filename)

			e, err := checkExists(path)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
			if e {
				c.Status(http.StatusBadRequest)
				return
			}

			err = c.SaveUploadedFile(f, path)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
		}

		c.Status(http.StatusOK)
	}
}

func PutFile() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func PatchFile() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteFile() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func checkExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
