package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/rhul-compsoc/compsoc-api-go/internal/database"
	"github.com/rhul-compsoc/compsoc-api-go/internal/router"
	"github.com/rhul-compsoc/compsoc-api-go/pkg/util"
)

// Stores the router
var r router.Router

// Don't come for me I can't be bothered
var s database.Store

func Run() {
	printAscii()
	log.Println("Starting app")

	go shutdown()

	log.Println("Load env")
	err := godotenv.Load()
	util.ErrOut(err)

	log.Println("Creating Store")
	s = database.New(database.PostgresDB())
	err = s.AutoMigrate()
	util.ErrOut(err)

	log.Println("Creating & starting Router")
	r = router.New()

	r.Use(cors.Default())
	r.RegisterRoutes(routes)
	r.NoRoute(reverseProxy())
	r.Run()
}

func shutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("")

	log.Println("Shutting down app")

	log.Println("Database closing")
	err := s.Close()
	util.ErrLog(err)

	log.Println("Router closing")
	r.Close()

	log.Println("Exiting app")

	os.Exit(0)
}

func printAscii() {
	fmt.Println(`
  _____                      _____            
 / ____|                    / ____|           
| |     ___  _ __ ___  _ __| (___   ___   ___ 
| |    / _ \| '_ \ _ \| '_ \\___ \ / _ \ / __|
| |___| (_) | | | | | | |_) |___) | (_) | (__ 
 \_____\___/|_| |_| |_| .__/_____/ \___/ \___|
                      | |                     
                      |_|       `,
	)
}
