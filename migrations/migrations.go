package migrations

import (
	"log"

	"github.com/Edwinfpirajan/practice.git/config"
	"github.com/Edwinfpirajan/practice.git/models"
)

func GenerateMigration() {
	config.DB.AutoMigrate(
		&models.Team{},
		&models.User{},
		&models.Country{},
		&models.City{},
	)

	config.DB.Exec("ALTER TABLE users ADD CONSTRAINT fk_team_id FOREIGN KEY (fk_team_id) REFERENCES teams(id) ON DELETE CASCADE ON UPDATE CASCADE")
	config.DB.Exec("ALTER TABLE teams ADD CONSTRAINT fk_city_id FOREIGN KEY (fk_city_id) REFERENCES cities(id) ON DELETE CASCADE ON UPDATE CASCADE")
	config.DB.Exec("ALTER TABLE cities ADD CONSTRAINT fk_country_id FOREIGN KEY (fk_country_id) REFERENCES countries(id) ON DELETE CASCADE ON UPDATE CASCADE")

	log.Println("Migración completada.")
}
