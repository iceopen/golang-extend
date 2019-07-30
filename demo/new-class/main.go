package main

func main() {

}

type user struct {
	name string
	age  uint32
}

func (u *user) SetInfo(name string, age uint32) {
	u.name = name
	u.age = age
}

func (u *user) GetAge() uint32 {
	return u.age
}

// NewEve 每次都声明一个变量
func NewEve() {
	i := 1
	for {
		u := user{}
		u.SetInfo("kkk", 123)
		u.GetAge()
		i++
		if i > 10000 {
			break
		}
	}
}

// OldEve 已经声明过的
func OldEve() {
	u := user{}
	i := 1
	for {
		u.SetInfo("kkk", 123)
		u.GetAge()
		i++
		if i > 10000 {
			break
		}
	}
}
