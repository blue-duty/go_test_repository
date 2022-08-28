package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	// 创建一个新的ssh配置对象
	//config := &ssh.ClientConfig{
	// // 设置超时时间
	// Timeout: time.Second * 10,
	// // 设置ssh的私钥
	// // PrivateKey: ssh.NewSignerFromKey(key),
	// // 设置ssh的私钥的密码
	// // Password: "123456",
	// User:            "cos",
	// Auth:            []ssh.AuthMethod{ssh.Password("we9621895")},
	// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	//}
	//dialect, err := ssh.Dial("tcp", "192.168.28.180:22", config)
	//if err != nil {
	// fmt.Println("ssh.Dial error:", err)
	// return
	//}
	//defer func(dialect *ssh.Client) {
	// err := dialect.Close()
	// if err != nil {
	//  fmt.Println("dialect.Close error:", err)
	// }
	//}(dialect)
	//session, err := dialect.NewSession()
	//if err != nil {
	// fmt.Println("dialect.NewSession error:", err)
	// return
	//}
	//defer func(session *ssh.Session) {
	// err := session.Close()
	// if err != nil && err != io.EOF {
	//  fmt.Println("session.Close error:", err)
	// }
	//}(session)
	//// 运行命令
	//// 读取命令的输出
	//stdin, err := session.StdinPipe()
	//if err != nil {
	// log.Fatal("输入错误", err)
	// return
	//}
	////设置session的标准输出和错误输出分别是os.stdout,os,stderr.就是输出到后台
	////stdout, err := session.StdoutPipe()
	//session.Stderr = os.Stderr
	//session.Stdout = os.Stdout
	//// 写入命令
	//err = session.Start("passwd")
	//if err != nil {
	// return
	//}
	//time.Sleep(time.Second * 5)
	////普通用户更改密码时需要输入旧密码
	//_, err = fmt.Fprintf(stdin, "we9621895\n")
	//if err != nil {
	// log.Fatal("输入错误", err)
	// return
	//}
	//time.Sleep(time.Second * 5)
	//_, err = fmt.Fprintf(stdin, "trunkey123\n")
	//if err != nil {
	// log.Fatal("执行错误", err)
	// return
	//}
	//time.Sleep(time.Second * 5)
	//_, err = fmt.Fprintf(stdin, "trunkey123\n")
	//if err != nil {
	// log.Fatal("执行错误", err)
	// return
	//}
	//time.Sleep(time.Second * 5)
	//fmt.Println("执行完成")

	//type test struct {
	// name string
	//}
	//
	//var t test
	//t.name = "test"
	//
	//fmt.Println(t)
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//err = db.AutoMigrate(&Language{}, &User{})
	//	//if err != nil {
	//	//	return
	//	//}
	//	//db.Model(&Language{}).Create(&Language{
	//	//	Name: "language" + strconv.Itoa(0),
	//	//	Users: []User{
	//	//		{
	//	//			Name: "user" + strconv.Itoa(0),
	//	//		},
	//	//		{
	//	//			Name: "user" + strconv.Itoa(1),
	//	//		},
	//	//	},
	//	//})

	//var c []User
	//sql := "SELECT u.id,u.name,u.created_at FROM  users AS u JOIN user_languages AS ul ON u.id = ul.user_id WHERE ul.language_id = ?"
	//db.Raw(sql, 1).Scan(&c)
	//var c []User
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

}

type User struct {
	Name string
	gorm.Model
	LanguageId uint
}

type Language struct {
	gorm.Model
	Name  string
	Users []User
}

//type Test struct {
// Name string json:"name"
//}

/*
type User struct {
 Id   string gorm:"primary_key"
 Name string gorm:"column:name"
}
func (User) TableName() string {
 return "user"
}

type User2 struct {
 Id   string gorm:"primary_key"
 User User
}

func (User2) TableName() string {
 return "user2"
}


//type User struct {
// gorm.Model
// CreditCard []CreditCard
//}
//
//type CreditCard struct {
// gorm.Model
// Number string
// User   User `gorm:"ForeignKey:UserID"`
//}
//
*/
