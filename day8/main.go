package main

type User struct {
    ID int64
    Name string
    Avatar string
}

func GetUserInfo() *User {
    return &User{
        ID: 6666,
        Name: "test",
        Avatar: "232",
    }
}
func main1()  {
    u := GetUserInfo()
    println(u.Name)
}

//https://blog.csdn.net/qcrao/article/details/118004683
