package main

import (
	"errors"
	"fmt"
)

func LogOutput(message string) {
  fmt.Println(message)
}

type SimpleDataStore struct {
  userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
  name, ok := sds.userData[userID]
  return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
  return SimpleDataStore{
    userData: map[string]string{
      "1": "Fred",
      "2": "Mary",
      "3": "Pat",
    },
  }
}

type DataStore interface {
  UserNameForID(userID string) (string, bool)
}

type Logger interface {
  Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
  lg(message)
}

type SimpleLogic struct {
  l Logger
  ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
  sl.l.Log("in SayHello for " + userID)
  name, ok := sl.ds.UserNameForID(userID)
  if !ok {
    return "", errors.New("unknown user")
  }
  return "Goodbye, " + name, nil
}
