package app

import (
	"log"
	"os"

	"github.com/fahribaharudin/petstore_restapi/app/controllers"
	"github.com/fahribaharudin/petstore_restapi/app/repositories"
	"github.com/fahribaharudin/petstore_restapi/app/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite driver
	"github.com/spf13/viper"
)

// Container as the wrapper of the app components
// inject all the components here..
type Container struct {
	PetController controllers.PetController
}

// config handler
var config *viper.Viper

// database handler
var db *gorm.DB

// the container instance
var container Container

func init() {

	// init the config
	config = viper.New()
	config.SetConfigFile("./config/default.json")
	if err := config.ReadInConfig(); err != nil {
		log.Println(err)
		os.Exit(0)
	}

	// init database
	var err error
	if config.GetString("env") == "dev" {
		db, err = gorm.Open(config.GetString("db.development.driver"), config.GetString("db.development.database"))
	} else if config.GetString("env") == "prod" {
		db, err = gorm.Open(config.GetString("db.production.driver"), config.GetString("db.production.database"))
	}

	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	// init container
	container = initCointainer()
}

// init the main container
func initCointainer() Container {
	container := Container{
		PetController: controllers.PetController{
			PetService: services.PetService{
				PetRepository: repositories.PetRepository{
					DbHandler: db,
				},
			},
		},
	}

	return container
}
