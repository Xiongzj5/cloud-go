package model

import (
	"errors"
)

// User is the type with all info of users in it
type User struct{
	Username string
	Stuid	 string
	Email    string
	Phone    string
}

var storage []User

//CheckUsername checks whether the username has been used
func CheckUsername(username string) bool {
	for _, v := range storage{
		if v.Username == username{
			return false
		}
	}
	return true
}

//CheckStuID checks whether the student id has been used
func CheckStuID(stuid string) bool{
	for _,v := range storage{
		if v.Stuid == stuid{
			return false
		}
	}
	return true
}

//CheckEmail checks whether the Email has been used
func CheckEmail(email string) bool{
	for _, v:=range storage{
		if v.Email == email{
			return false;
		}
	}
	return true;
}

//CheckPhone checks whether the Phone number has been used
func CheckPhone(phone string) bool{
	for _, v:=range storage{
		if v.Phone == phone{
			return false
		}
	}
	return true
}

//AddUser can add a User to the storage
func AddUser(username, stuid, email, phone string){
	storage = append(storage, User{
		Username: username,
		Stuid:	  stuid,
		Email: 	  email,
		Phone:	  phone,
	})
}

//FetchInfo use username to get all infos of a single user
//If no User is found, err will not be nil
func FetchInfo(username string) (user User, err error){
	for _, v := range storage{
		if v.Username == username{
			return v, nil
		}
	}
	return User{},errors.New("No Such User")
}
