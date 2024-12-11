package user

type User struct {
	Fname string
	Lname string
	Age   int
}
type UserMapType = map[string]interface{}

var UserSlice = make([]UserMapType, 0, 10)

func CreateUser(user User) {
	var userMap = make(UserMapType)
	userMap["fname"] = user.Fname
	userMap["lname"] = user.Lname
	userMap["age"] = user.Age
	UserSlice = append(UserSlice, userMap)
}
func GetUser() []UserMapType {
	return UserSlice
}