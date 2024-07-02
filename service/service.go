package service

import (
	"github.com/islombay/noutbuk_seller/config"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	redisdb "github.com/islombay/noutbuk_seller/storage/redis"
)

type ServiceInterface interface {
	Auth() *Auth
	Category() *Category
	Brand() *Brand
	Seller() *Seller
	Files() *Files
}

type Service struct {
	storage storage.StorageInterface
	log     logs.LoggerInterface
	cfg     config.Config
	redis   *redisdb.RedisStore

	auth     *Auth
	category *Category
	brand    *Brand
	seller   *Seller
	files    *Files
}

func New(storage storage.StorageInterface, log logs.LoggerInterface, cfg config.Config) ServiceInterface {
	srv := Service{storage: storage, log: log}

	redis := redisdb.NewRedisStore(&cfg.Redis, log)

	srv.auth = NewAuth(storage, log, redis)
	srv.category = NewCategory(storage, log)
	srv.brand = NewBrand(storage, log)
	srv.seller = NewSeller(storage, log)
	srv.files = NewFiles(storage, log, cfg)

	srv.redis = redis

	srv.cfg = cfg

	return &srv
}

func (s *Service) Auth() *Auth {
	if s.auth == nil {
		s.auth = NewAuth(s.storage, s.log, s.redis)
	}
	return s.auth
}

func (s *Service) Category() *Category {
	if s.category == nil {
		s.category = NewCategory(s.storage, s.log)
	}
	return s.category
}

func (s *Service) Brand() *Brand {
	if s.brand == nil {
		s.brand = NewBrand(s.storage, s.log)
	}
	return s.brand
}

func (s *Service) Seller() *Seller {
	if s.seller == nil {
		s.seller = NewSeller(s.storage, s.log)
	}
	return s.seller
}

func (s *Service) Files() *Files {
	if s.files == nil {
		s.files = NewFiles(s.storage, s.log, s.cfg)
	}
	return s.files
}
