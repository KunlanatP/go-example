package cmd

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kunlanat/go-example/api/v1/books"
	"github.com/kunlanat/go-example/config"
	"github.com/kunlanat/go-example/migration"
	"github.com/kunlanat/go-example/repository"
	"github.com/kunlanat/go-example/service"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Print the start number of Hugo",
	Long:  `All software has starts. This is Hugo's`,
	Run:   startServer,
}

var port uint16

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().Uint16VarP(&port, "port", "p", 8000, "Port to run application server")
}

func startServer(cmd *cobra.Command, args []string) {
	// dsn := "postgresql://postgres:password@localhost:5432/example?sslmode=disable"
	db, err := gorm.Open(postgres.Open(config.Default.DB_URL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	migration.AutoMigrate(db)

	app := fiber.New()

	apiGroup := app.Group("/api")
	GroupV1 := apiGroup.Group("/v1")

	repo := repository.BookRepositoryWithGORM(db)
	serv := service.BookServiceImp(repo)
	GroupV1.Mount("/", books.BookController(serv))

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
