package db

import (
	"github.com/Andres-MK/internal/domain/entities"
	"github.com/Andres-MK/internal/domain/repositories"
	"gorm.io/gorm"
)

type EmailRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) repositories.EmailRepository {
	return &EmailRepository{ db: db }
}

func (repo *EmailRepository) Create(email *entities.ValidatedEmail) (error) {
	
	dbEmail := toDBEmail(email)
	if err := repo.db.Table(dbEmail.TableName()).Create(&dbEmail).Error; err != nil {
		return err
	}

	return nil
}

func (repo *EmailRepository) GetAll() ([]*entities.Email, error) {
	var dbEmails []*entities.Email
	if err := repo.db.Table("EmailsRegister").Find(&dbEmails).Error; err != nil {
		return nil, err
	}

	
	
	return dbEmails, nil
}

