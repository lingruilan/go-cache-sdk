package get_animal

type Cat struct {
	Voice string
	Name  string
}

func NewCat(voice string, name string) *Cat {
	return &Cat{
		Voice: voice,
		Name:  name,
	}
}

func (c Cat) GetName() string {
	return c.Name
}

func (c Cat) GetVoice() string {
	return c.Voice
}
