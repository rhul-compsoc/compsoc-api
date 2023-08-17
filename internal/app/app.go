package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/router"
	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
)

// Stores the router
var r *router.Router

// Don't come for me I can't be bothered
var s *database.Store

func Run() {
	log.Println("Starting app")

	log.Println("Load env")
	err := godotenv.Load()
	util.ErrOut(err)

	log.Println("Creating Store")
	s = database.New()

	log.Println("Creating & starting Router")
	r = router.New()
	//r.Use(middleware.MakeAuth())
	r.RegisterRoutes(makeRoutes())
	r.NoRoute(reverseProxy())
	r.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down app")
}
