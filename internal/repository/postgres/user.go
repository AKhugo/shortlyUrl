package postgres

import (
	"shortlyurl/internal/domain"

	"gorm.io/gorm"
)



type UserRepository struct {
	db *gorm.DB // database connection
	// redis *redis.Client // redis connection

}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}


// CreateUser creates a new user
func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}


func (r *UserRepository) UpdateUserInfo(email string, userInfo *domain.User) error {
	return r.db.Model(&domain.User{}).Where("email = ?", email).Updates(userInfo).Error
}


func (r *UserRepository) DeleteUser(id uint) error {
	// soft delete the user
	return r.db.Delete(&domain.User{}, id).Error
}

// hard delete the user
func (r *UserRepository) HardDeleteUser(id uint) error {
	return r.db.Unscoped().Delete(&domain.User{}, id).Error
}