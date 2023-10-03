package post

type Package interface {
	Send()
}

type PostOfficeBranch struct {
	City string
	Name string
}

func (branch *PostOfficeBranch) Proccessing(pack Package) {
	pack.Send()
}
