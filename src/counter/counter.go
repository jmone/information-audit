package counter

type ItemCounter struct{
	ItemType string
	CounterId int32
	MaxDocId int32
}

func (ic ItemCounter)GetMaxDocId(itemType string, counterId int32) int32{
	//
	return 10024
}

func (ic ItemCounter)SetMaxDocId(itemType string, counterId int32, newDocId int32){
	println("Set max doc id success.")
}
