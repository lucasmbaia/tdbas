package repository

import (
	"reflect"
	"errors"
	"time"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	DB  *gorm.DB
}

type RepositoryConfig struct {
	Username          string
	Password          string
	Host              string
	Port              string
	DBName            string
	Timeout           string
	Debug             bool
	ConnsMaxIdle      int
	ConnsMaxOpen      int
	ConnsMaxLifetime  int
}

func NewRepository(cfg RepositoryConfig) (r Repository, err error) {
	if r.DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Timeout)); err != nil {
		return
	}

	r.DB.LogMode(cfg.Debug)
	r.DB.DB().SetMaxIdleConns(cfg.ConnsMaxIdle)
	r.DB.DB().SetMaxOpenConns(cfg.ConnsMaxOpen)
	r.DB.DB().SetConnMaxLifetime(time.Duration(cfg.ConnsMaxLifetime))

	return
}

func (r *Repository) Create(entity interface{}) (err error) {
	if reflect.ValueOf(entity).Kind() != reflect.Ptr {
		err = errors.New("The target struct is required to be a pointer")
		return
	}

	err = r.DB.Create(entity).Error
	return
}

func (r *Repository) ReadAll(condition, entity interface{}) (ok bool, err error) {
	if reflect.ValueOf(entity).Kind() != reflect.Ptr {
		err = errors.New("The target struct is required to be a pointer")
		return
	}

	var operation *gorm.DB = r.DB.Set("gorm:auto_preload", true).Find(entity, condition)

	if operation.Error != nil {
		err = operation.Error
	} else {
		ok = true
	}

	return
}

func (r *Repository) Update(condition, entity interface{}) (err error) {
	if reflect.ValueOf(entity).Kind() != reflect.Ptr || reflect.ValueOf(condition).Kind() != reflect.Ptr {
		err = errors.New("The target struct is required to be a pointer")
		return
	}

	if operation := r.DB.Model(entity).Where(condition).Update(entity); operation.Error != nil {
		err = operation.Error
	}

	return
}

func (r *Repository) ReadByConditions(entity, conditions interface{}) (exists bool, err error) {
	if reflect.ValueOf(entity).Kind() != reflect.Ptr {
		err = errors.New("The target struct is required to be a pointer")
		return
	}

	var operation *gorm.DB = r.DB.First(entity, conditions)

	if operation.RecordNotFound() {
		return
	}

	if operation.Error != nil {
		err = operation.Error
	}
	exists = true

	return
}

func (r *Repository) Delete(condition interface{}) (exists bool, err error) {
	var (
		entity    reflect.Value
		operation *gorm.DB
	)

	if reflect.ValueOf(condition).Kind() != reflect.Ptr {
		err = errors.New("The target struct is required to be a pointer")
		return
	}

	entity = reflect.New(reflect.ValueOf(condition).Type().Elem()).Elem()
	operation = r.DB.Where(condition).Delete(entity.Interface())

	if operation.RecordNotFound() {
		return
	}

	if operation.Error != nil {
		exists = true
	}

	return
}
