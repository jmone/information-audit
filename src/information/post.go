package information

type Post struct{
	pid int32
	tid int32
	message string
}

func (p Post)GetInformations(page, size int32){
	println("post informtations")
}

func (p Post)Delete(id int32){
	println("delete")
}
