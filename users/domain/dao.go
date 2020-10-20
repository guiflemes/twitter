package domain

import "github.com/guiflemes/twitter_clone/utils/errors"

func (u *User) CheckUSerExist() bool{
	return true
}

func (u *User) Create() *errors.RestErr{
	return nil
}
