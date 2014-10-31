package main

import(
	"information"
	"counter"
	"fmt"
)

func loadDict(path string){

}

func init(){
	fmt.Println("欢迎使用Discuz敏感词过滤系统！")
}

func main(){
	var i information.Information
	var p information.Post
	var t information.Thread
	var counter counter.ItemCounter

	i = p
	i.GetInformations(1, 10)
	i = t
	i.GetInformations(1, 10)
	maxDocId := counter.GetMaxDocId("post", 1)
	counter.SetMaxDocId("post", 1, maxDocId+1000)
}
