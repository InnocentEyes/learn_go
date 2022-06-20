package main

import "fmt"

//map

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) //还没有初始化 (没有在内存中开辟空间)
	m1 = make(map[string]int,10) //要估算该map容量,避免在程序运行期间再动态扩容
	m1["理想"] = 18
	m1["jiwuming"] = 35
	fmt.Println(m1)
	//约定成俗用ok接收返回的布尔值
	fmt.Println(m1["nz"]) //如果不存在这个key拿到对应值类型的零值
	value,ok := m1["nz"]
	if !ok{
		fmt.Println("查无此key")
	}else {
		fmt.Println(value)
	}
	
	//map的遍历
	for k, v := range m1 {
		fmt.Println(k,v)
	}

	//key的遍历
	for k := range m1 {
		fmt.Println(k)
	}

	//value的遍历
	for _, v := range m1 {
		fmt.Println(v)
	}

	//删除
	delete(m1,"jiwuling")
	fmt.Println(m1)
	delete(m1,"gz") //删除不存在的key

}
