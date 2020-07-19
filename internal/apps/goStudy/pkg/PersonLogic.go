package mypkg

// SetPerson Personインスタンスのプロパティに値を設定します
func (person *PersonModel) SetPerson(name string, age int) {
	person.Name = name
	person.Age = age
}

// GetPerson Personインスタンスを取得します
func (person *PersonModel) GetPerson() (string, int) {
	return person.Name, person.Age
}

// ToString はPersonのプロパティを文字列化して返却します
func (person *PersonModel) ToString() string {
	return person.Name + "," + string(person.Age)
}
