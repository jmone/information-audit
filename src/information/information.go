package information

type Information interface{
	GetInformations(page, size int32)
	Delete(id int32)
}
