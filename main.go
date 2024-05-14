package main

import (
	"fmt"
	"reflect"
)

/**
  反射的优点：更具有通用性、适用于不同的类型，适用于编译阶段不能确定类型，运行时动态的确定类型，json.Unmarshal()底层用到了反射根据字段名去匹配和解析
  反射的缺点：代码不易阅读、不易维护、易发生线上panic; 性能很差，比正常代码慢一到两个数量级
  反射的Type和Value: type用于获取类型相关的信息（slice的长度、struct的成员、函数的参数个数);
  value用于获取和修改原始数据的值（修改slice和map中的元素，修改struct的成员变量）
*/

type People interface {
	GetName()
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"`
}

func (u *User) GetName() {
	//fmt.Println(u.Name)
}

func (u *User) GetAge() int {
	//fmt.Println(u.Age)
	return u.Age
}

func (u *User) GetSex() {
	//fmt.Println(u.Sex)
}

func main() {

	//reflect.type
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test5_1()

	//reflect.value
	//test6()
	//test7()
	//test8()
	//test9()
	//test10()
	//test11()

	//test12()
	//test13()
	test14()
}

//reflect.type
func test1() {

	fmt.Println("==================")
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("hello")
	fmt.Println(typeI)
	fmt.Println(typeS)

	typeUser := reflect.TypeOf(&User{})
	fmt.Println(typeUser)
	fmt.Println(typeUser.Kind())
	fmt.Println(typeUser.Elem())
	fmt.Println(typeUser.Elem().Kind())

}

//reflect.type
func test2() {
	fmt.Println("==================")
	typeUser := reflect.TypeOf(&User{})
	typeUser2 := reflect.TypeOf(User{})

	fmt.Println(typeUser)
	fmt.Println(typeUser2)

	fmt.Println(typeUser.Elem() == typeUser2)

}

//获取结构体成员变量
func test3() {

	fmt.Println("==================")
	typeUser := reflect.TypeOf(User{})
	fieldNum := typeUser.NumField() // 成员变量的个数
	for i := 0; i < fieldNum; i++ {
		field := typeUser.Field(i) // 返回一个struct, 该结构体描述了这个字段的属性等
		fmt.Printf("%d %s offset %d anonymous %t type %s exported %t json tag %s\n", i,
			field.Name,         // 变量名称
			field.Offset,       // 相对于结构体首地址的内存偏移量，string类型会占据16个字节, int类型占用8个字节,
			field.Anonymous,    // 是否为匿名成员
			field.Type,         // 数据类型
			field.IsExported(), // 包外是否可见
			field.Tag.Get("json"))
	}

	fmt.Println()
	if nameFiled, ok := typeUser.FieldByName("Name"); ok {
		fmt.Printf("name is exported %t\n", nameFiled.IsExported())
	}

	thirdField := typeUser.FieldByIndex([]int{2})
	fmt.Printf("third field name %s\n", thirdField.Name)

}

//获取结构体成员方法
func test4() {

	fmt.Println("==================")
	typeUser := reflect.TypeOf(User{})
	methodNum := typeUser.NumMethod() //成员方法的个数 接受者为值的放
	for i := 0; i < methodNum; i++ {
		method := typeUser.Method(i)
		fmt.Printf("method name: %s, type: %s, exported:%t\n", method.Name, method.Type, method.IsExported())
	}

	if method, ok := typeUser.MethodByName("GetSex"); ok {
		fmt.Println(method)
	}

	fmt.Println()
	typeUser2 := reflect.TypeOf(&User{})
	methodNum = typeUser2.NumMethod() //接受者为指针和值的方法都包含在内
	for i := 0; i < methodNum; i++ {
		method := typeUser2.Method(i)
		fmt.Printf("method name: %s, type: %s, exported:%t\n", method.Name, method.Type, method.IsExported())
	}

}

func Add(a, b int) int {
	return a + b
}

//直接获取函数的信息
func test5() {
	typeFunc := reflect.TypeOf(Add)
	fmt.Println(typeFunc.Kind() == reflect.Func)
	argInNum := typeFunc.NumIn()   //输入参数的个数
	argOutNum := typeFunc.NumOut() //输出参数的个数
	for i := 0; i < argInNum; i++ {
		argType := typeFunc.In(i)
		fmt.Printf("第%d个输入参数的类型%s, kind: %s\n", i, argType, argType.Kind())
	}

	for i := 0; i < argOutNum; i++ {
		argType := typeFunc.Out(i)
		fmt.Printf("第%d个输入参数的类型%s, kind: %s\n", i, argType, argType.Kind())
	}
}

// 判断类型是否实现了某接口
func test5_1() {

	typeOfPeople := reflect.TypeOf((*People)(nil)).Elem()
	fmt.Println(typeOfPeople)
	fmt.Printf("typeOfPeople kind is interface %t \n", typeOfPeople.Kind() == reflect.Interface)
	//t1 := reflect.TypeOf(User{})
	t2 := reflect.TypeOf(&User{})
	fmt.Printf("t2 implements people interface %t \n", t2.Implements(typeOfPeople))
}

//reflect.value
func test6() {
	iValue := reflect.ValueOf(1)
	sValue := reflect.ValueOf("hello")
	userPtrValue := reflect.ValueOf(&User{
		Sex:  true,
		Name: "杰克逊",
		Age:  99,
	})

	fmt.Println(iValue)
	fmt.Println(sValue)
	fmt.Println(userPtrValue)

	//value转换为Type
	iType := iValue.Type()
	sType := sValue.Type()
	userType := userPtrValue.Type()

	//在type上调用kind()和在相应的value上调用相应的kind()的结果是一样的
	fmt.Println(iType.Kind() == reflect.Int, iValue.Kind() == reflect.Int, iType.Kind() == iValue.Kind())
	fmt.Println(sType.Kind() == reflect.String, sValue.Kind() == reflect.String, sType.Kind() == sValue.Kind())
	fmt.Println(userType.Kind() == reflect.Ptr, userPtrValue.Kind() == reflect.Ptr, userType.Kind() == userPtrValue.Kind())

	//指针value和非指针value相互转换
	userValue := userPtrValue.Elem() //指针value转换为非指针value
	fmt.Println(userValue.Kind(), userPtrValue.Kind())
	userPtrValue3 := userValue.Addr() //非指针value转化为指针value
	fmt.Println(userValue.Kind(), userPtrValue3.Kind())

	//得到value的原始数据
	fmt.Printf("origin value iValue is %d %d\n", iValue.Interface().(int), iValue.Int())
	fmt.Printf("origin value sValue is %s %s\n", sValue.Interface().(string), sValue.String())
	user := userValue.Interface().(User)
	fmt.Printf("sex=%t name=%s age=%d\n", user.Sex, user.Name, user.Age)
	user2 := userPtrValue.Interface().(*User)
	fmt.Printf("sex=%t name=%s age=%d\n", user2.Sex, user2.Name, user2.Age)
}

//空value的判断
func test7() {
	var i interface{}
	v := reflect.ValueOf(i)
	fmt.Printf("v持有值:%v %t, type of v is Invalid %t\n", v, v.IsValid(), v.Kind() == reflect.Invalid)

	var user *User = nil
	v = reflect.ValueOf(user) //Value指向一个nil
	if v.IsValid() {
		fmt.Printf("v持有的值是:%v %t\n", v, v.IsNil())
	}

	var u User
	v = reflect.ValueOf(u)
	if v.IsValid() {
		fmt.Printf("v持有的值是：%v %t\n", v, v.IsZero())
	}
}

//通过value修改原始数据的值
//未导出的成员不能通过反射进行修改
func test8() {
	var i int = 10
	var s string = "hello"
	user := User{
		Sex:  true,
		Name: "张三丰",
		Age:  22,
	}
	valueI := reflect.ValueOf(&i) //要修改就在这里要传指针
	valueS := reflect.ValueOf(&s)
	valueUser := reflect.ValueOf(&user)
	valueI.Elem().SetInt(8) //valueI对应的原始对象是指针，通过Elem()返回指针指向的对象
	valueS.Elem().SetString("golang")
	valueUser.Elem().FieldByName("Age").SetInt(77)
	fmt.Println(user)
}

//通过value修改slice
//通过value修改map
func test9() {

	users := make([]*User, 1, 5)
	users[0] = &User{
		Sex:  false,
		Name: "杰克逊",
		Age:  88,
	}
	sliceValue := reflect.ValueOf(&users)
	if sliceValue.Elem().Len() > 0 {
		sliceValue.Elem().Index(0).Elem().FieldByName("Name").SetString("令狐冲")
	}
	fmt.Println(*users[0])

	sliceValue.Elem().SetLen(2)
	sliceValue.Elem().Index(1).Set(reflect.ValueOf(&User{
		Sex:  true,
		Name: "张三丰",
		Age:  101,
	}))

	fmt.Println(*users[0], *users[1])

	//修改map
	u1 := &User{
		Sex:  true,
		Name: "杰克逊",
		Age:  55,
	}

	u2 := &User{
		Sex:  false,
		Name: "李四光",
		Age:  23,
	}

	userMap := make(map[int]*User, 5)
	userMap[1] = u1

	mapValue := reflect.ValueOf(&userMap)
	//往map添加一个key-value
	mapValue.Elem().SetMapIndex(reflect.ValueOf(2), reflect.ValueOf(u2))
	//根据key取出对应的map
	mapValue.Elem().MapIndex(reflect.ValueOf(1)).Elem().FieldByName("Name").SetString("张三丰")

	fmt.Println(*userMap[1], *userMap[2])

}

//调用函数
func test10() {

	valueFunc := reflect.ValueOf(Add) //函数也是一种数据类型
	typeFunc := reflect.TypeOf(Add)
	argsNum := typeFunc.NumIn()            //reflect.Type 获取输入参数的个数
	args := make([]reflect.Value, argsNum) //准函数的输入参数
	for i := 0; i < argsNum; i++ {
		if typeFunc.In(i).Kind() == reflect.Int {
			args[i] = reflect.ValueOf(3)
		}
	}

	sumValue := valueFunc.Call(args) //返回[]reflect.Value
	fmt.Println(sumValue)
	if typeFunc.Out(0).Kind() == reflect.Int {
		sum := sumValue[0].Interface().(int) //从value转换为原始类型
		fmt.Printf("sum=%d\n", sum)
	}

}

//调用成员方法
//指针类型可以调用值的方法，值不可以调用指针的方法
func test11() {
	user := User{
		Sex:  false,
		Name: "王麻子",
		Age:  34,
	}
	valueUser := reflect.ValueOf(&user)
	ageMethod := valueUser.MethodByName("GetAge")
	resultValue := ageMethod.Call([]reflect.Value{}) //无参数时传一个空的切片
	result := resultValue[0].Interface().(int)
	fmt.Println(result)

}

//创建struct
func test12() {
	t := reflect.TypeOf(User{})
	value := reflect.New(t)
	value.Elem().FieldByName("Age").SetInt(10)
	user := value.Interface().(*User)
	fmt.Println(user)
}

//创建slice
func test13() {
	var slice []User
	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.MakeSlice(sliceType, 1, 3)
	sliceValue.Index(0).Set(reflect.ValueOf(User{
		Age:  8,
		Name: "李达",
	}))
	users := sliceValue.Interface().([]User)
	fmt.Printf("1st user name %s\n", users[0].Name)
}

func test14() {

	var userMap map[int]*User
	mapType := reflect.TypeOf(userMap)

	mapValue := reflect.MakeMapWithSize(mapType, 10)
	user := &User{
		Name: "杰克逊",
		Age:  44,
		Sex:  false,
	}
	key := reflect.ValueOf(user.Age)
	mapValue.SetMapIndex(key, reflect.ValueOf(user))
	mapValue.MapIndex(key).Elem().FieldByName("Name").SetString("令狐一刀")
	userMap = mapValue.Interface().(map[int]*User)
	fmt.Printf("user name %s %s \n", userMap[44].Name, user.Name)
}
