package main

import (
	"fmt"
	"reflect"
)

func Type_reflect(a interface{}) {
	// 获取类型信息 reflect.TypeOf
	t := reflect.TypeOf(a)
	fmt.Println("type of a:", t)
	k := t.Kind()
	switch k {
	case reflect.Float64:
		fmt.Println(k)
	case reflect.String:
		fmt.Println("is a string")
	}

}

func Value_reflect(a interface{}) {
	// 获取值信息 reflect.ValueOf
	v := reflect.ValueOf(a)
	fmt.Println(v)
	// 获取类型，
	k := v.Kind()

	switch k {
	case reflect.Int64:
		fmt.Println("int:", v.Int()) //如果值为Int类型，可以用v.Int()获取值
	case reflect.Float64:
		fmt.Println("float:", v.Float())
	}
}

func Set_reflect(a interface{}) {
	// 通过反射设置变量的值，需要传入指针类型变量，不然会报错
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Int:
		v.SetInt(100)
		fmt.Println("Setting value successful!")
	case reflect.Float64:
		v.SetFloat(3.3) // 报错reflect.Value.SetFloat using unaddressable value
		fmt.Println("Setting value successful!")
	case reflect.Ptr: // 指针类型属于Ptr，修改时使用v.Elem().Set...进行设置
		v.Elem().SetFloat(3.2)
		fmt.Println("Setting value successful!")
	default:
		fmt.Println("default switch!")
	}

}

type User struct {
	Name string `json:"name" db:"name"`
	Age  int
}

func (s *User) SetName(name string) {
	s.Name = name
	fmt.Println(s)
}

func (s *User) Print() {
	fmt.Println("This is a struct")
}

// 通过反射读取结构体内的字段属性和值
func Struct_reflect(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Println("is a int64")
	case reflect.String:
		fmt.Println("is a string")
	case reflect.Struct:
		fmt.Println("is a struct")
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)           // 获取每个字段类型,名称
			val := v.Field(i).Interface() // 获取每个字段的值
			fmt.Printf("name:%s,type:%v value:%v\n", field.Name, field.Type,
				val)
		}
	}
}

// 通过反射设置结构体字段值
func Set_struct_reflect(s interface{}) {
	v := reflect.ValueOf(s)
	v.Elem().Field(0).SetString("godric")  // 通过索引设置
	v.Elem().FieldByName("Age").SetInt(25) // 通过字段名设置
	fmt.Printf("%#v", s)
}

// 通过reflect.Typeof()查询结构体方法,reflect.ValueOf()调用方法
func Method_struct_reflect(s interface{}) {
	t := reflect.TypeOf(s)
	fmt.Printf("This struct has %d method\n", t.NumMethod())
	// 结构体内方法排序安装方法名称英文字母顺序排列
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("%d method,name:%s,type:%v\n", i, method.Name, method.Type)
	}

	// 调用方法method
	v := reflect.ValueOf(s)
	m1 := v.MethodByName("Print")
	// 定义一个reflect.Value的空切片传入没有参数的方法
	var args []reflect.Value
	m1.Call(args)

	// 调用有参数method
	m2 := v.MethodByName("SetName")
	var args2 []reflect.Value
	name := "Godric"
	// 将传入的参数使用reflect.ValueOf()反射后添加到切片，调用时传入函数
	nameVal := reflect.ValueOf(name)
	args2 = append(args2, nameVal)
	m2.Call(args2)

}

// 通过refle.TpyeOf()获取结构体字段的tag
func Tag_struct_reflect(s interface{}) {
	t := reflect.TypeOf(s)
	field0 := t.Field(0)
	fmt.Println("json tag:", field0.Tag.Get("json"))
	fmt.Println("db tag:", field0.Tag.Get("db"))
}

func main() {
	// a := 3.4
	// Type_reflect(a)
	// Value_reflect(a)
	// Set_reflect(&a)
	// fmt.Println(a)

	s := User{
		Name: "alex",
		Age:  22,
	}
	// Struct_reflect(s)

	Set_struct_reflect(&s)

	Method_struct_reflect(&s)

	Tag_struct_reflect(s)
}
