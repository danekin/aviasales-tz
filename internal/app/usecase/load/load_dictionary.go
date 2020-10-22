package load

type Usecase struct {
	dict Dictionary
}

func NewLoadDictionaryUsecase(dict Dictionary) *Usecase {
	return &Usecase{
		dict: dict,
	}
}

func (u *Usecase) Do(dict []string) {
	u.dict.Set(dict)
}
