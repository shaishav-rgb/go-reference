package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

func login(uid, pwd string) (string, error) {
	// world's worst login system
	if uid == "admin" && pwd == "admin" {
		return "user:admin", nil
	}
	return "", errors.New("bad user")
}

func getData(token, file string) ([]byte, error) {
	// world's worst access control
	if token == "user:admin" {
		switch file {
		case "secrets.txt":
			return []byte("passwords aplenty!"), nil
		case "payroll.csv":
			return []byte("everyone's salary"), nil
		}
	}
	return nil, os.ErrNotExist
}

func GenerateErrorBroken(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func GenerateErrorOKReturnNil(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

func GenerateErrorUseErrorVar(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	// data, err := LoginAndGetData("admin", "admin", "secrets.txt")
	// fmt.Println(string(data))
	// fmt.Println(err)

	// data, err = LoginAndGetData("admin", "admin", "chicken_recipe.txt")
	// fmt.Println(string(data))
	// fmt.Println(err)

	// data, err = LoginAndGetData("jon", "password", "secrets.txt")
	// fmt.Println(string(data))
	// fmt.Println(err)

	err := GenerateErrorBroken(true)
	fmt.Println("GenerateErrorBroken(true) returns non-nil error:", err != nil)
	err = GenerateErrorBroken(false)
	fmt.Println("GenerateErrorBroken(false) returns non-nil error:", err != nil)
	err = GenerateErrorOKReturnNil(true)
	fmt.Println("GenerateErrorOKReturnNil(true) returns non-nil error:", err != nil)
	err = GenerateErrorOKReturnNil(false)
	fmt.Println("GenerateErrorOKReturnNil(false) returns non-nil error:", err != nil)
	err = GenerateErrorUseErrorVar(true)
	fmt.Println("GenerateErrorUseErrorVar(true) returns non-nil error:", err != nil)
	err = GenerateErrorUseErrorVar(false)
	fmt.Println("GenerateErrorUseErrorVar(false) returns non-nil error:", err != nil)
}
