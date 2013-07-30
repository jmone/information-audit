package information

type Thread struct{
	tid int32
	subject string
	message string
}

func (t Thread)GetInformations(page, size int32){
	println("thread informations.")
}

func (t Thread)Delete(id int32){

}
