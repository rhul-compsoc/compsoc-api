package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rhul-compsoc/compsoc-api-go/internal/models"
	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
)

// Gets a Guild with id given in the parameter.
//   - /guild/:guild
func GuildGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("guild")
		url := fmt.Sprintf("https://discord.com/api/guilds/%s?with_counts=true", i)

		client := http.Client{
			Timeout: time.Second * 2,
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		util.LogErr(err)

		req.Header.Add("Authorization", os.Getenv("DISCORD_TOK"))
		req.Header.Set("User-Agent", "compsoc-api")

		res, err := client.Do(req)
		util.LogErr(err)

		if res.Body != nil {
			defer res.Body.Close()
		}

		var g models.Guild
		err = json.NewDecoder(res.Body).Decode(&g)
		util.LogErr(err)

		c.JSON(http.StatusOK, g)
	}
}
