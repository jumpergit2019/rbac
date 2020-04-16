package main

import (
	"database/sql"
	"fmt"

	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type User struct {
//	gorm.Model
//	Name         string
//	Age          sql.NullInt64
//	Birthday     *time.Time
//	Email        string  `gorm:"type:varchar(100);unique_index"`
//	Role         string  `gorm:"size:255"`        // set field size to 255
//	MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
//	Num          int     `gorm:"AUTO_INCREMENT"`  // set num to auto incrementable
//	Address      string  `gorm:"index:addr"`      // create index with name `addr` for address
//	IgnoreMe     int     `gorm:"-"`               // ignore this field
//}
//
//func (User) TableName() string {
//	return "usertable"
//}
type Animal struct {
	ID   int64
	Name string        `gorm:"default:'galeone'"`
	Age  sql.NullInt32 `gorm:"default:18"`
}

//func (animal *Animal) BeforeCreate(scope *gorm.Scope) error {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	scope.SetColumn("id", r.Int63n(100000))
//	return nil
//}

func main() {
	db, err := gorm.Open("mysql", "jumper:J1u2m3p!e@r#@tcp(192.168.1.35:3306)/test?loc=Local&parseTime=true&charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema

	////////////////////////////////////create data
	//db.DropTableIfExists("animals")
	//db.AutoMigrate(&Animal{})
	//
	//animal := Animal{
	//	Name: "",
	//	Age: sql.NullInt32{
	//		Int32: 0,
	//		Valid: false,
	//	},
	//}
	//fmt.Println("is new: ", db.NewRecord(&animal))
	//db.Create(&animal)
	//fmt.Println("is new: ", db.NewRecord(&animal))

	//var a Animal
	//db.First(&a)
	//fmt.Println("a: ", a)
	//
	//var a10 Animal
	//db.First(&a10, 10)
	//fmt.Println("a10: ", a10)
	//
	//var b Animal
	//db.Last(&b)
	//fmt.Println("b: ", b)
	//
	//var as []Animal
	//db.Find(&as)
	//fmt.Println("as: ", as)

	//var aa Animal
	//db.Where("id = ?", 16).First(&aa)
	//fmt.Println(aa)
	//
	//var aas []Animal
	//db.Where("id > ?", 10).Find(&aas)
	//fmt.Println(aas)
	//
	//var abs []Animal
	//names := []string{"wang", "ru", "zhang", "zhao"}
	//db.Where("name in (?)", names).Find(&abs)
	//fmt.Println(abs)
	//
	//var acs []Animal
	//db.
	//	//Debug().
	//	Where("age > ?", 30).
	//	Where("name in (?)", names).
	//	Or("age < ?", 18).
	//	Find(&acs)
	//fmt.Println("acs: ", acs)
	//
	//var ads []Animal
	//db.
	//	//Debug().
	//	Where("age between ? and ?", 30, 60).
	//	Find(&ads)
	//fmt.Println("ads: ", ads)
	//
	//var ae Animal
	//db.
	//	Debug().
	//	Set("gorm:query_option", "for update").
	//	Where("id = ?", 1).
	//	Find(&ae)
	//fmt.Println(ae)

	//需要order by才使用 first / last
	//不需要排序直接使用 find , 查找结果多个就使用 [], 一个就使用变量
	//使用 where 方式，更加清晰，容易修改， 尽量不使用其他方式

	//var af Animal
	//ra := db.
	//	Debug().
	//	Where(&Animal{Name: "llllllll"}).
	//	Attrs(&Animal{Age: struct {
	//		Int32 int32
	//		Valid bool
	//	}{Int32: 111, Valid: true}}).
	//	FirstOrInit(&af).RowsAffected
	//
	//fmt.Println(af, ra)
	//
	//if ra == 0 {
	//	db.Create(&af)
	//}
	//
	//var ag Animal
	//db.
	//	Debug().
	//	Where(&Animal{Name: "ffffff"}).
	//	Attrs(&Animal{Age: struct {
	//		Int32 int32
	//		Valid bool
	//	}{Int32: 999, Valid: true}}).
	//	FirstOrCreate(&ag)
	//fmt.Println(ag)

	//firstorinit 用于判断是否存在该记录，若是存在，就获取内容，若是不存在就使用 attrs 创建内存对象，
	//之后可以判断 rows affected 来确定是否获取到值，没有则可以使用create 来插入数据。
	//上述 firstorinit + rows affected + create 等价于 firstorcreate
	//总结： 有就获取，没有就创建

	//var ahs []Animal
	//db.Debug().
	//	Where("age > ?", db.Table("animals").Select("avg(age)").SubQuery()).
	//	Find(&ahs)
	//fmt.Println(ahs)
	//
	////子查询需要使用 subquery, 使用方式如上
	//
	//var ais []Animal
	//db.Debug().Select("name, age").Where("age > ?", 30).Find(&ais)
	//fmt.Println(ais)
	//
	//var ajs []Animal
	//var aks []Animal
	//db.Debug().Where("name in (?)", []string{"wang", "zhang"}).Order("name desc").Find(&ajs).Order("age desc", true).Find(&aks)
	//fmt.Println(ajs)
	//fmt.Println(aks)

	//order("x asc, y desc") 等价于 order("x asc").order("y desc")
	//order("x asc").find(&a1).order("y desc", true).find(&a2)
	//相当于进行了两个查询，分别使用两个order

	//var als []Animal
	////db.Limit(3).Find(&als)
	//db.Debug().Offset(2).Limit(-1).Find(&als)
	//fmt.Println(als)

	//发现 limit 可以单独使用， 但是 offset 不能单独使用，只能与 limit 一起使用
	//即便使用 offset(3).limit(-1) 也不行
	//也就是说不能实现 从某个记录开始往后面获取所有记录，发现mysql5.7 也不支持 limit 3,-1

	//var cnt int32
	//db.Model(&Animal{}).Where("age > ?", 30).Count(&cnt)
	//fmt.Println(cnt)
	//
	//var ams []Animal
	//db.Debug().Where("age > ?", 30).Find(&ams).Count(&cnt)
	//fmt.Println(ams)
	//fmt.Println(cnt)
	//
	//db.Debug().Table("animals").Select("count(distinct(name))").Count(&cnt)
	//fmt.Println(cnt)

	//需要同时获取记录和数量，使用 find + count
	//只需要获取数量，使用 model + count
	//需要计算聚合函数，使用 table + select + count

	rows, err := db.Table("animals").Select("id, name, age").Rows()

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var t Animal
		err := rows.Scan(&t.ID, &t.Name, &t.Age)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(t)
	}

	fmt.Println("------------------")

	rows2, err := db.Table("animals").Select("name, sum(age) as ages").Group("name").Having("sum(age) > ?", 50).Rows()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows2.Close()
	for rows2.Next() {
		var t Animal
		err := rows2.Scan(&t.Name, &t.Age)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(t)
	}

	fmt.Println("------------------")

	type Tmp struct {
		Name string
		Ages int32
	}
	var tmps []Tmp
	var cnt int32
	db.Table("animals").Select("name, sum(age) as ages").Group("name").Having("sum(age) > ?", 50).Scan(&tmps).Count(&cnt)
	fmt.Println("tmps: ", tmps)
	fmt.Println("cnt: ", cnt)

	//一般查询也可以使用  table + select + rows, 相对于find/first 可以对每个数据进行进一步操作之后再存入内存（rows.next rows.scan)
	//分组使用 table + select + group + having + rows/scan

}
