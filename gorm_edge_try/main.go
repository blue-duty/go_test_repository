package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Cxk{}, &Chicken{})
	if err != nil {
		return
	}

	db.Model(&Cxk{}).Create(&Cxk{
		Id:   1,
		Name: "cxk",
		Chickens: []Chicken{
			{
				Id:   1,
				Name: "chicken1",
			},
			{
				Id:   2,
				Name: "chicken2",
			},
		},
	})
	db.Model(&Cxk{}).Create(&Cxk{
		Id:   2,
		Name: "cxk2",
		Chickens: []Chicken{
			{
				Id:   1,
				Name: "chicken1",
			},
			{
				Id:   4,
				Name: "chicken4",
			},
		},
	})
	/*
		db.Model(&Cxk{}).Create(&Cxk{
			Id:   1,
			Name: "cxk",
			Chickens: []Chicken{
				{
					Id:   1,
					Name: "chicken1",
				},
				{
					Id:   2,
					Name: "chicken2",
				},
			},
		})

		db.Model(&Chicken{}).Create(&Chicken{
			Id:    3,
			Name:  "chicken3",
			CxkId: 1,
		})

		db.Model(&Cxk{}).Create(&Cxk{
			Id:   1,
			Name: "cxk",
			Chicken: Chicken{
				Id:   1,
				Name: "chicken",
			},
		})

		db.Model(&Language{}).Create(&Language{
			Name: "language" + strconv.Itoa(0),
			Users: []User{
				{
					Name: "user" + strconv.Itoa(0),
				},
				{
					Name: "user" + strconv.Itoa(1),
				},
			},
		})

		var c []User
		sql := "SELECT u.id,u.name,u.created_at FROM  users AS u JOIN user_languages AS ul ON u.id = ul.user_id WHERE ul.language_id = ?"
		db.Raw(sql, 1).Scan(&c)
		var c []User
		var l Language
		err = db.Preload(clause.Associations, "name not IN (?)", "user1").Find(&l).Error
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(l)
		fmt.Println(l.Users)
		for _, v := range l.Users {
			fmt.Println(v.Name)
		}

		var cxk Cxk
		db.Preload("Chickens").Where("id = ?", 1).First(&cxk)
		fmt.Printf("%#v", cxk)
	*/
	var cxks []Cxk
	db.Preload("Chickens", "id in (1)").Find(&cxks)
	fmt.Printf("%#v", cxks)
}

/*
type Cxk struct {
	Id      uint   `gorm:"primary_key"`
	Name    string `gorm:"column:name"`
	Chicken Chicken
}
func (Cxk) TableName() string {
	return "cxk"
}

type Chicken struct {
	Id    uint   `gorm:"primary_key"`
	Name  string `gorm:"column:name"`
	CxkId uint
}
func (Chicken) TableName() string {
	return "chicken"
}

type Cxk struct {
	Id       uint   `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Chickens []Chicken
}
func (Cxk) TableName() string {
	return "cxk"
}

type Chicken struct {
	Id    uint   `gorm:"primary_key"`
	Name  string `gorm:"column:name"`
	CxkId uint
}
func (Chicken) TableName() string {
	return "chicken"
}
*/

type Cxk struct {
	Id       uint      `gorm:"primary_key"`
	Name     string    `gorm:"column:name"`
	Chickens []Chicken `gorm:"many2many:cxk_chicken;"`
}

func (Cxk) TableName() string {
	return "cxk"
}

type Chicken struct {
	Id   uint   `gorm:"primary_key"`
	Name string `gorm:"column:name"`
}

func (Chicken) TableName() string {
	return "chicken"
}
