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

	// here we're sending nil instead of map. Coz nil represents as zero-value for map,interface,int,string,slice,array,map
	return nil, "User not found"
}
