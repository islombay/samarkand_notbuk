package postgresql

import (
	"fmt"
	"os"

	"github.com/islombay/noutbuk_seller/config"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DuplicateKeyError = "23505"
)

type Storage struct {
	user     storage.UserI
	category storage.CategoryI
	brand    storage.BrandI
	seller   storage.SellerI
	files    storage.FilesI
	product  storage.ProductI

	db  *gorm.DB
	log logs.LoggerInterface
}

func NewPostgresStore(cfg config.DBConfig, log logs.LoggerInterface) storage.StorageInterface {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, os.Getenv("DB_USER"), os.Getenv("DB_PWD"),
		cfg.DBName, cfg.Port, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	st := Storage{db: db, log: log}

	st.user = NewUserRepo(db, log)
	st.category = NewCategoryRepo(db, log)
	st.brand = NewBrandRepo(db, log)
	st.seller = NewSellerRepo(db, log)
	st.files = NewFilesRepo(db, log)
	st.product = NewProductRepo(db, log)

	return &st
}

func (s *Storage) Close() {
}

func (s *Storage) User() storage.UserI {
	return s.user
}

func (s *Storage) Category() storage.CategoryI {
	return s.category
}

func (s *Storage) Brand() storage.BrandI {
	return s.brand
}

func (s *Storage) Seller() storage.SellerI {
	return s.seller
}

func (s *Storage) Files() storage.FilesI {
	return s.files
}

func (s *Storage) Product() storage.ProductI {
	return s.product
}
