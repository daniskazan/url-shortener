package users

import (
	"github.com/daniskazan/url-shortener/internal/core/users"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Save(user users.User) {
	repo.DB.Create(&user)
}

func (repo *UserRepository) FindByID(userID int) (*users.User, error) {
	// Создаем объект пользователя, который будет заполнен данными из базы данных
	user := &users.User{}

	// Используем GORM для поиска пользователя по идентификатору
	if err := repo.DB.First(user, userID).Error; err != nil {
		// Если произошла ошибка при поиске, возвращаем её
		return nil, err
	}

	// Возвращаем найденного пользователя
	return user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}
