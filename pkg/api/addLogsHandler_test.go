package api

import (
	"context"
	"errors"
	
	"main/pkg/common/models"
	"net/http"
	"time"

	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/assert.v1"
)


type MockDBPostgre struct {
    mock.Mock
}

func (m *MockDBPostgre) CheckKeyExist(ctx context.Context, key string) (bool, error) {
	args := m.Called(ctx, key)
	return args.Bool(0), args.Error(1)
   }
   
type MockDBMongo struct {
	mock.Mock

}

func (m *MockDBMongo) GetLogs(ctx context.Context, key string) ([]models.LogsModel, error) {
	args := m.Called(ctx, key)
	return args.Get(0).([]models.LogsModel), args.Error(1)
}
func (m *MockDBMongo) GetLogsById(ctx context.Context, key string, reqid string) ([]models.LogsModel, error) {
	args := m.Called(ctx, key)
	return args.Get(0).([]models.LogsModel), args.Error(1)
}

func (m *MockDBMongo) AddLog(ctx context.Context, logData models.LogsModel) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, logData)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
   }
   

func TestAcceptLogs(t *testing.T) {
    tests := []struct {
        name             string
        key              string
        expectedResult    bool
        expectedError error
        expectedStatus int
    }{
        {
            name: "Key exists",
            key: "existing_key",
            expectedResult: true,
            expectedError: nil,
            expectedStatus: http.StatusOK,
        },
        {
            name: "Key does not exist",
            key: "non_existing_key",
            expectedResult: false,
            expectedError: errors.New("Key not found"),
            expectedStatus: http.StatusNotFound,
        },
    }


	
	addLogTests := []struct {
        name           string
        logData        models.LogsModel
        expectedResult *mongo.InsertOneResult
        expectedError  error
    }{
        {
            name: "AddLog success",
            logData: models.LogsModel{
                Key: "31313132sasda",
                Timestamp: time.Time{},
                Level: "INFO",
                Message: "successfully added",
            },
            expectedResult: &mongo.InsertOneResult{InsertedID: 1},
            expectedError: nil,
        },
        {
            name: "AddLog unsuccess",
            logData: models.LogsModel{
                Key: "31313132sasda",
                Timestamp: time.Time{},
                Level: "INFO",
                Message: "failed to add",
            },
            expectedResult: nil,
            expectedError: errors.New("Error! While adding a log to the database"),
        },
    }

	
	
    mockDBPostgre := new(MockDBPostgre)
    mockDBMongo := new(MockDBMongo)

    h := handler{
        DB:         mockDBPostgre,
        MongoDB:    mockDBMongo,
        Validator:  validator.New(),
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            mockDBPostgre.On("CheckKeyExist", mock.Anything, test.key).Return(test.expectedResult, test.expectedError)

            keyExists, err := h.DB.CheckKeyExist(context.Background(), test.key)

            actualStatus := http.StatusOK
            if err != nil {
                actualStatus = http.StatusNotFound
            } else if !keyExists {
                actualStatus = http.StatusInternalServerError
            }

            assert.Equal(t, test.expectedResult, keyExists)
            assert.Equal(t, test.expectedError, err)
            assert.Equal(t, test.expectedStatus, actualStatus)
        })
    }

    for _, addLogTest := range addLogTests {
        t.Run(addLogTest.name, func(t *testing.T) {
            expectedStatus := http.StatusOK
            if addLogTest.expectedError != nil {
                expectedStatus = http.StatusInternalServerError
            }

            mockDBMongo.On("AddLog", mock.Anything, addLogTest.logData).Return(addLogTest.expectedResult, addLogTest.expectedError)

            result, err := h.MongoDB.AddLog(context.Background(), addLogTest.logData)

            actualStatus := expectedStatus
            if err != nil {
                actualStatus = http.StatusInternalServerError
            }

            assert.Equal(t, addLogTest.expectedResult, result)
            assert.Equal(t, addLogTest.expectedError, err)
            assert.Equal(t, expectedStatus, actualStatus)
        })
    }

	
}