package Petrovich

import (
	"fmt"
	"testing"
)

var p *rules
var err error

func Init(path string) {
	if p == nil {
		p, err = LoadRules(path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func TestRules_Firstname(t *testing.T) {

	Init("rules.json")

	if p.InfFirstname("Дамир", Genitive, "male") != "Дамира" {
		t.Error("Firstname false")
	}

	if p.InfFirstname("Анна-Мария", Dative, "female") != "Анне-Марии" {
		t.Error("Firstname false")
	}

}

func TestRules_Lastname(t *testing.T) {
	Init("rules.json")

	if p.InfLastname("Кочубей", Dative, "male") != "Кочубею" {
		t.Error("Lastname false")
	}

	if p.InfLastname("Козлов", Dative, "male") != "Козлову" {
		t.Error("Lastname false")
	}

	if p.InfLastname("Салтыков-Щедрин", Dative, "male") != "Салтыкову-Щедрину" {
		t.Error("Lastname false")
	}

	if p.InfLastname("Дюма", Dative, "male") != "Дюма" {
		t.Error("Lastname false")
	}

	if p.InfLastname("Воробей", Dative, "male") != "Воробью" {
		t.Error("Lastname false")
	}

	if p.InfLastname("Плевако", Dative, "male") != "Плевако" {
		t.Error("Lastname false")
	}

}

func TestRules_Middlename(t *testing.T) {
	Init("rules.json")

	if p.InfMiddlename("Борух-Бендитовна", Dative, "female") != "Борух-Бендитовне" {
		t.Error("Middlename false")
	}

	if p.InfMiddlename("Георгиевна-Авраамовна", Dative, "female") != "Георгиевне-Авраамовне" {
		t.Error("Middlename false")
	}

}

func TestRules_InfFio(t *testing.T) {
	Init("rules.json")

	if p.InfFio("Цой Виктор Робертович", Accusative, true) != "Цоя В.Р." {
		t.Error("Fio short version false")
	}

	if p.InfFio("Белякова Ирина Сергеевна", Dative, false) != "Беляковой Ирине Сергеевне" {
		t.Error("Fio long version false")
	}
}

// Other tests

type testCase struct {
	name     string
	gender   string
	infNames []string
}

func TestFirstName(t *testing.T) {
	cases := []testCase{
		{"Анна-Мария", "female", []string{
			"Анны-Марии",
			"Анне-Марии",
			"Анну-Марию",
			"Анной-Марией",
			"Анне-Марии"}},
		{"Василий", "male", []string{
			"Василия",
			"Василию",
			"Василия",
			"Василием",
			"Василии"}},
		{"Кочубей", "male", []string{
			"Кочубея",
			"Кочубею",
			"Кочубея",
			"Кочубеем",
			"Кочубее"}},
		{"Лев", "male", []string{
			"Льва",
			"Льву",
			"Льва",
			"Львом",
			"Льве"}},
		{"Маша", "female", []string{
			"Маши",
			"Маше",
			"Машу",
			"Машей",
			"Маше"}},
		{"John", "male", []string{
			"John",
			"John",
			"John",
			"John",
			"John"}},
	}

	for _, c := range cases {
		for i, in := range c.infNames {
			if p.InfFirstname(c.name, i, c.gender) != in {
				t.Errorf("Fail in %s - %s", c.name, in)
			}
		}
	}
}

func TestMiddleName(t *testing.T) {
	cases := []testCase{
		{"Георгиевна-Авраамовна", "female", []string{
			"Георгиевны-Авраамовны",
			"Георгиевне-Авраамовне",
			"Георгиевну-Авраамовну",
			"Георгиевной-Авраамовной",
			"Георгиевне-Авраамовне"}},
		{"Петрович", "male", []string{
			"Петровича",
			"Петровичу",
			"Петровича",
			"Петровичем",
			"Петровиче"}},
		{"Петровна", "female", []string{
			"Петровны",
			"Петровне",
			"Петровну",
			"Петровной",
			"Петровне"}},
	}

	for _, c := range cases {
		for i, in := range c.infNames {
			if p.InfMiddlename(c.name, i, c.gender) != in {
				t.Errorf("Fail in %s - %s", c.name, in)
			}
		}
	}
}

func TestLastName(t *testing.T) {
	cases := []testCase{
		{"Андрейчук", "male", []string{
			"Андрейчука",
			"Андрейчуку",
			"Андрейчука",
			"Андрейчуком",
			"Андрейчуке"}},
		{"Воробей", "male", []string{
			"Воробья",
			"Воробью",
			"Воробья",
			"Воробьем",
			"Воробье"}},
		{"Андрейчук", "female", []string{
			"Андрейчук",
			"Андрейчук",
			"Андрейчук",
			"Андрейчук",
			"Андрейчук"}},
		{"Дюма", "male", []string{
			"Дюма",
			"Дюма",
			"Дюма",
			"Дюма",
			"Дюма"}},
		{"Петров-Водкин", "male", []string{
			"Петрова-Водкина",
			"Петрову-Водкину",
			"Петрова-Водкина",
			"Петровым-Водкиным",
			"Петрове-Водкине"}},
		{"Салтыков-Щедрин", "male", []string{
			"Салтыкова-Щедрина",
			"Салтыкову-Щедрину",
			"Салтыкова-Щедрина",
			"Салтыковым-Щедриным",
			"Салтыкове-Щедрине"}},
	}

	for _, c := range cases {
		for i, in := range c.infNames {
			if p.InfLastname(c.name, i, c.gender) != in {
				t.Errorf("Fail in: have %s wait %s", c.name, in)
			}
		}
	}
}
