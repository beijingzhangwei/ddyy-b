package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Topic string
	Age   int
}

func main() {
	//s := new(Student)
	s := &Student{Name: "222"}
	fmt.Println(reflect.TypeOf(s))        //  ----> *main.Student
	fmt.Println(reflect.TypeOf(s).Elem()) //  ----> main.Student
	fmt.Println(reflect.TypeOf(*s))       //  ----> main.Student
	v := reflect.ValueOf(s).Elem()        //  (依靠不安全的操作存储器对齐来强制转换)方法将运行时类型和变量转换为反射类型和变量
	v.Field(0).SetString("66")
	// 最后，调用 Interface ()方法将变量从反射状态转换回运行时状态。
	fmt.Printf("%#v\n", v.Interface()) //  ----> main.Student{Name:"66", Age:0}

	t := Teacher{}
	fmt.Println(reflect.TypeOf(&t).Elem())
	//fmt.Println(reflect.TypeOf(t).Elem()) panic
	field, _ := reflect.TypeOf(&t).Elem().FieldByName("Topic")
	fmt.Println("field.PkgPath--->", field.PkgPath)
	fmt.Println("field--->", field)
	fmt.Println(reflect.TypeOf(&t).Elem().FieldByName("Age"))
	fmt.Println(reflect.TypeOf(&t))
	fmt.Println(reflect.TypeOf(t).Field(0))
	fmt.Println(reflect.TypeOf(t).Field(0))
	fmt.Println(reflect.TypeOf(t).Name())
	fmt.Println(reflect.TypeOf(t).PkgPath())
	fmt.Println(reflect.TypeOf(t).FieldByName("Topic"))

	// 通过指针反射操作 内部实现 *(*string)(v.ptr) = x
	// 即需要通过地址来完成操作
	reflect.ValueOf(&t).Elem().Field(0).SetString("99999")
	fmt.Println("field Value SetString --->", t)
	reflect.ValueOf(t).Field(1)
	fmt.Println("field Value get --->", reflect.ValueOf(t).Field(0))

	// 运行时变量
	//type Value struct {
	//	typ Type
	//	ptr uintptr
	//}
	//
	//type Type interface {
	//	Name() string         // by all type
	//	Index(int) Value      // by Slice Array
	//	MapIndex(Value) Value // by Map
	//	Send(Value)           // By Chan
	//}
	// 首先，我们根据 Golang 的规则简单地定义一个变量类型 Value。
	// 值具有两种类型的成员属性 type 和 ptr。
	// Type 是指示这个变量是什么对象的类型，ptr 是指向这个变量的地址的地址。

	// 在操作变量时，我们根据 Type 类型进行操作，操作对象的数据位于内存的 ptr 位置。
	// 变量类型 Type 定义一个接口，因为不同的类型有不同的操作方法。
	// 例如，在 Map 中获取/设置值的方法，
	// 在 Slice 和 Array 中获取索引的方法，
	// 发送和接收一个 Chan 类型的对象的方法，
	// 在 Struct 获取结构属性的方法，以及属性也可以有一个标记。
	// 这样，不同的类型有不同的唯一操作方法。如果 Map 类型无法实现 Index 方法，则会引起恐慌。
	// 了解一个变量的本质是一个数据地址和一种数据类型，然后根据这两个变量进行操作即可体现出来。

	// 反射类型
	// 让我们来看看源文件 return/type.go 中的一段代码。
	// type rtype struct {
	//    size    uintptr
	//    ptrdata uintptr
	//    kind    uint8
	//    ...
	//}
	// Rtype 对象是 Type 接口的简化实现，kind 是这种类型的类型，
	// 然后其他复合类型(Ptr、 Slice、 Map 等)具有其他属性和方法。
	// // ptrType represents a pointer type.
	// type ptrType struct {
	//    rtype
	//    elem *rtype // pointer element (pointed at) type
	// }

}

type Value struct {
	typ Type
	ptr uintptr
}

type Type interface {
	Name() string         // by all type
	Index(int) Value      // by Slice Array
	MapIndex(Value) Value // by Map
	Send(Value)           // By Chan
}
