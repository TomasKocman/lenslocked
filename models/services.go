package models

import "github.com/jinzhu/gorm"

type Services struct {
	Gallery GalleryService
	User    UserService
	Image   ImageService
	db      *gorm.DB
}

func NewServices(connectionInfo string) (*Services, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &Services{
		User:    NewUserService(db),
		Gallery: NewGalleryService(db),
		Image:   NewImageService(),
		db:      db,
	}, nil
}

func (s *Services) Close() error {
	return s.db.Close()
}

func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}, &Gallery{}).Error
}

func (s *Services) DestructiveReset() error {
	if err := s.db.DropTableIfExists(&User{}, &Gallery{}).Error; err != nil {
		return err
	}
	return s.AutoMigrate()
}
