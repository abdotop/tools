package dbcrudops

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	return db
}

func TestCreate(t *testing.T) {
	db := setupDatabase(t)
	operator := New(db)

	// Define a mock data structure
	type MockData struct {
		ID   uint
		Name string
	}

	// Migrate the schema
	err := operator.Migrate(&MockData{})
	assert.NoError(t, err)

	mock := &MockData{Name: "Test"}
	err = operator.Create(mock)
	assert.NoError(t, err)

	var result MockData
	db.First(&result, 1)
	assert.Equal(t, "Test", result.Name)
}

func TestRead(t *testing.T) {
	db := setupDatabase(t)
	operator := New(db)

	type MockData struct {
		ID   uint
		Name string
	}

	err := operator.Migrate(&MockData{})
	assert.NoError(t, err)

	mock := &MockData{Name: "ReadTest"}
	db.Create(mock)

	var readMock MockData
	err = operator.Read(&readMock, mock.ID)
	assert.NoError(t, err)
	assert.Equal(t, "ReadTest", readMock.Name)
}

func TestUpdate(t *testing.T) {
	db := setupDatabase(t)
	operator := New(db)

	type MockData struct {
		ID   uint
		Name string
	}

	err := operator.Migrate(&MockData{})
	assert.NoError(t, err)

	mock := &MockData{Name: "BeforeUpdate"}
	db.Create(mock)

	mock.Name = "AfterUpdate"
	err = operator.Update(mock)
	assert.NoError(t, err)

	var updatedMock MockData
	db.First(&updatedMock, mock.ID)
	assert.Equal(t, "AfterUpdate", updatedMock.Name)
}

func TestDelete(t *testing.T) {
	db := setupDatabase(t)
	operator := New(db)

	type MockData struct {
		ID   uint
		Name string
	}

	err := operator.Migrate(&MockData{})
	assert.NoError(t, err)

	mock := &MockData{Name: "ToDelete"}
	db.Create(mock)

	err = operator.Delete(mock)
	assert.NoError(t, err)

	var deletedMock MockData
	result := db.First(&deletedMock, mock.ID)
	assert.Error(t, result.Error) // Expect an error because the record should be deleted
}

func TestFindByKey(t *testing.T) {
	db := setupDatabase(t)
	operator := New(db)

	type MockData struct {
		ID   uint
		Name string
	}

	err := operator.Migrate(&MockData{})
	assert.NoError(t, err)

	mock := &MockData{Name: "FindByKeyTest"}
	db.Create(mock)

	var foundMock MockData
	err = operator.FindByKey(&foundMock, "name", "FindByKeyTest")
	assert.NoError(t, err)
	assert.Equal(t, "FindByKeyTest", foundMock.Name)
}

func TestOnError(t *testing.T) {
	db := setupDatabase(t) // Assuming setupDatabase is a function that sets up your DB
	operator := New(db)

	// Create a channel to receive errors from the callback
	receivedErrors := make(chan error, 1)

	// Define the callback function to capture errors
	callback := func(err error) {
		receivedErrors <- err
	}

	// Start listening for errors
	operator.OnError(callback)

	// Simulate an error
	testError := errors.New("test error")
	operator.errChan <- testError

	// Wait for the error to be received or timeout
	select {
	case err := <-receivedErrors:
		assert.Equal(t, testError, err, "The received error should match the test error")
	case <-time.After(1 * time.Second):
		t.Fatal("Expected an error but didn't receive one")
	}

	// Clean up: Close the error channel to prevent goroutine leaks
	close(operator.errChan)
}
