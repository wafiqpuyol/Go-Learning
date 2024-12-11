package auth

import (
	"github.com/wafiqpuyol/Go-Learning/user"
)

func IsUseExist(username string) (user.UserMapType, string) {
	for _, userCred := range user.UserSlice {
		if userCred["fname"] == username || userCred["lname"] == username {
			return userCred, ""
		}
	}

	return nil, "User not found"
}
