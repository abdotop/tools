package dbcrudops

import (
	"html"

	"gorm.io/gorm"
)

// type db_type any

// var all_types = []db_type{}

type Operator struct {
	db *gorm.DB
	// _type   db_type
	errChan chan error // Channel to send errors
}

func New(db *gorm.DB) *Operator {
	return &Operator{
		db:      db,
		errChan: make(chan error),
	}
}

func (o *Operator) GetDb() *gorm.DB {
	return o.db
}

func (o *Operator) OnError(callback func(error)) {
	go func() {
		for err := range o.errChan { // Correctly range over the channel
			if err != nil {
				callback(err) // Call the callback function with the error
			}
		}
	}()
}

func (o *Operator) Migrate(models ...interface{}) error {
	for _, model := range models {
		err := o.db.AutoMigrate(model)
		if err != nil {
			o.errChan <- err
			return err
		}
	}
	return nil
}

func (o *Operator) Exec(sql string, values ...interface{}) error {
	result := o.db.Exec(sql, values...)
	if result.Error != nil {
		o.errChan <- result.Error
		return result.Error
	}
	return nil
}

func (o *Operator) Create(data interface{}) error {
	result := o.db.Create(data)
	if result.Error != nil {
		o.errChan <- result.Error
		return result.Error
	}
	return nil
}

func (o *Operator) Read(data interface{}, id interface{}) error {
	result := o.db.First(data, id)
	if result.Error != nil {
		o.errChan <- result.Error
		return result.Error
	}
	return nil
}

func (o *Operator) Update(data interface{}) error {
	result := o.db.Save(data)
	if result.Error != nil {
		o.errChan <- result.Error
		return result.Error
	}
	return nil
}

func (o *Operator) Delete(data interface{}) error {
	result := o.db.Delete(data)
	if result.Error != nil {
		o.errChan <- result.Error
		return result.Error
	}
	return nil
}

func (o *Operator) FindByKey(data interface{}, key string, value interface{}) error {
	result := o.db.Where(html.EscapeString(key)+" = ?", value).Find(data)
	if result.Error != nil {
		o.errChan <- result.Error
		return result.Error
	}
	return nil
}
