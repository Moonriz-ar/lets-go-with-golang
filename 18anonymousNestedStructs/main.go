package main

type Address struct {
	Province   string
	City       string
	CreateTime string
}

type Email struct {
	Account    string
	CreateTime string
}

type User struct {
	Name    string
	Gender  string
	Address // 匿名字段 anonymous fields
	Email   // 匿名字段 anonymous fields
}

func main() {
	var user2 User
	user2.Name = "Mario"
	user2.Gender = "Male"
	// user2.CreateTime ="2019" // ambiguous select user2.CreateTime
	user2.Address.CreateTime = "2000" // CreateTime inside nested struct Address
	user2.Email.CreateTime = "2000"   // CreateTime inside nested struct Email
}
