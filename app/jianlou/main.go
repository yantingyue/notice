package main

import "notice/app/jianlou/api"

func init() {
	//cli.InitRedisClient()
}

func main() {
	api.Begin()
}

//func main() {
//	list := LinkedList{}
//	list.AddNode("A")
//	list.AddNode("B")
//	list.AddNode("C")
//	fmt.Println(list)
//	//遍历链表
//	list.Traverse(func(data interface{}) {
//		fmt.Println(data)
//	})
//	//
//	////删除节点并再次遍历链表
//	//list.RemoveNode("B")
//	//list.Traverse(func(data interface{}) {
//	//	fmt.Println(data)
//	//})
//}
