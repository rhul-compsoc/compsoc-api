package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
)

// /file
func ListFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := make([]models.File, 0)

		w := func(path string, i os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			f := models.File{
				Path:  path,
				Size:  i.Size(),
				IsDir: i.IsDir(),
			}

			r = append(r, f)
			return nil
		}

		err := filepath.Walk("./uploads", w)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// Get a single file.
// /file/:file
func GetFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("file")

		n = strings.ReplaceAll(n, "|", "/")
		path := path.Join("uploads/", n)
		fmt.Println(path)

		e, err := checkExists(path)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		f, err := os.Open(path)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		defer f.Close()

		h := make([]byte, 512)
		_, err = f.Read(h)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		t := http.DetectContentType(h)
		i, err := f.Stat()
		if err != nil {
			c.Status(http.StatusBadRequest)
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

// Upload a single file.
// /file
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

// Upload multiple files.
// /files
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

// Upload or update a file.
// /file
func PutFile() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Update a file.
// /file
func PatchFile() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Delete a file.
// /file/:file
func DeleteFile() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Check if a file exists.
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
