package postgres

import (
	"cfd/myapp/config"
	"cfd/myapp/pkg/db/postgres/repository"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db             *gorm.DB
	UserRepository *repository.UserRepository
	RoleRepository *repository.RoleRepository
}

func NewDatabase(dbConfig config.DbConfig) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Dbname,
		dbConfig.Port,
		dbConfig.Sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to DB")
	}

	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	return &Database{
		db:             db,
		UserRepository: userRepository,
		RoleRepository: roleRepository,
	}, nil
}

func (d *Database) Migrate(entities ...interface{}) error {
	err := d.db.AutoMigrate(entities...)
	if err != nil {
		return err
	}

	return nil
}
