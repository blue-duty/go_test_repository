### 1. 对于一对一的关系的尝试
初始化一张cxk表和一张chicken表，关联关系是一对一的关系，关联关系的类型是has_one，关联关系的方向是has_one的方向。
``` go
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
```
跳过连接数据库的过程，我们初始化完成cxk表以及chicken表后，添加一只cxk的信息，并且添加一只chicken的信息
我们可以在添加cxk的时候直接为其绑定一只chicken。
``` go
    err = db.AutoMigrate(&Cxk{}, &Chicken{})
    if err != nil {
        return
    }

    db.Model(&Cxk{}).Create(&Cxk{
        Id:   1,
        Name: "cxk",
        Chicken: Chicken{
            Id:   1,
            Name: "chicken",
        },
    })
```
这样我们就可以在查询cxk的时候直接查询出其绑定的chicken的信息。
``` go
var cxk Cxk
db.First(&cxk)
fmt.Printf("%#v", cxk)
```
使用上面的方法我们只能得到这样的结果
``` go
main.Cxk{Id:0x1, Name:"cxk", Chicken:main.Chicken{Id:0x0, Name:"", CxkId:0x0}}
```
我们只查询出cxk所对应的信息，却没有查出他拥有的chicken的信息，我们可以通过gorm的Preload方法来解决这个问题。
``` go
var cxk Cxk
db.Preload("Chicken").First(&cxk)
fmt.Printf("%#v", cxk)
// 结果：main.Cxk{Id:0x1, Name:"cxk", Chicken:
// main.Chicken{Id:0x1, Name:"chicken", CxkId:0x1}}
```
### 2. 对于一对多的关系的尝试
一只cxk肯定不止拥有一只chicken，我们可以通过has_many的方式来实现一对多的关系。
``` go
type Cxk struct {
    Id      uint   `gorm:"primary_key"`
    Name    string `gorm:"column:name"`
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
```
如法炮制，我们初始化了一张cxk表和一张chicken表，关联关系是一对多的关系，关联关系的类型是has_many，关联关系
的方向是has_many的方向。并为表初始化了一组数据。
``` go
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
```
这次我们不在添加cxk的时候直接为其绑定一只chicken，而是在添加chicken的时候直接为其绑定一个cxk。
``` go
    db.Model(&Chicken{}).Create(&Chicken{
        Id:   3,
        Name: "chicken3",
        CxkId: 1,
    })
```
这样我们就可以在查询cxk的时候直接查询出其绑定的chicken的信息。
``` go
var cxk Cxk
db.Preload("Chickens").First(&cxk)
fmt.Printf("%#v", cxk)
// 结果：main.Cxk{Id:0x1, Name:"cxk", Chickens:[]main.Chicken{
// main.Chicken{Id:0x1, Name:"chicken1", CxkId:0x1}, 
// main.Chicken{Id:0x2, Name:"chicken2", CxkId:0x1}, 
// main.Chicken{Id:0x3, Name:"chicken3", CxkId:0x1}}}
```
### 3. 对于多对多的关系的尝试
一只cxk肯定不止拥有一只chicken，一只chicken肯定不止属于一个cxk，我们可以通过many_to_many的方式来
实现多对多的关系。
``` go
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
```
此时将会生成一张cxk_chicken表，保存了cxk和chicken的关系，我们为其添加值。
``` go
    err = db.AutoMigrate(&Cxk{}, &Chicken{})
    if err != nil {
        return
    }
    db.Model(&Cxk{}).Create(&Cxk{
        Id:   1,
        Name: "cxk1",
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
```
此时cxk1拥有chicken1和chicken2，cxk2拥有chicken1和chicken4。
查询cxk1以及chicken1的信息
``` go
// 查询cxk1的信息
var cxk Cxk
db.Preload("Chickens").Where("id = ?", 1).First(&cxk)
fmt.Printf("%#v", cxk)
// 结果：main.Cxk{Id:0x1, Name:"cxk", Chickens:[]main.Chicken{
// main.Chicken{Id:0x1, Name:"chicken1"}, 
// main.Chicken{Id:0x2, Name:"chicken2"}}}

// 查询chicken1的所有cxk的信息
var cxks []Cxk
db.Preload("Chickens", "id in (1)").Find(&cxks) //预加载的条件是id in (1)
fmt.Printf("%#v", cxks)
//结果：[]main.Cxk{main.Cxk{Id:0x1, Name:"cxk", Chickens:[]main.Chicken{
// main.Chicken{Id:0x1, Name:"chicken1"}}}, 
// main.Cxk{Id:0x2, Name:"cxk2", Chickens:[]main.Chicken{
// main.Chicken{Id:0x1, Name:"chicken1"}}}}
```
