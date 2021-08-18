![Petrovich](petrovich.png)

Go библиотека для склонения русскоязычных фамилий, имён и отчеств. На вход ФИО в именительном падеже, на выходе
склоненное значение в указанном падеже. Порт библиотеки https://github.com/petrovich.

Golang port of petrovich library(https://github.com/petrovich) which inflects Russian names to a given grammatical case

## Installation

```
go get github.com/whimax/Petrovich-Go
```

Файл с правилами здесь  
https://github.com/petrovich/petrovich-rules
## Usage

```go
package main

import (
	"fmt"
	"github.com/whimax/Petrovich-Go"
)

func main() {
	p, err := Petrovich.LoadRules("Rules/rules.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	// Склонение составной фамилии
	// Салтыкову-Щедрину
	fmt.Println(p.InfLastname("Салтыков-Щедрин", Petrovich.Dative, "male"))
	// Склонение имени
	// Катюше
	fmt.Println(p.InfFirstname("Катюша", Petrovich.Dative, "female"))
	// Склонение отчества
	// Юрьевичем
	fmt.Println(p.InfMiddlename("Юрьевич", Petrovich.Instrumental, "male"))
	// Склонение ФИО, реузльтат в сокращенном виде
	// Беликовой И.П.
	fmt.Println(p.InfFio("Беликова Ирина Петровна", Petrovich.Dative, true))
	// Склонение ФИО, реузльтат в полном виде
	// Цоя Виктора Робертовича
	fmt.Println(p.InfFio("Цой Виктор Робертович ", Petrovich.Accusative, false))
}
```

Значения пола:
    "male" - мужской
    "female" - женский
    в остальных случаях пол будет считаться неопределенным "Androgynous"


| Case                      | Case (in Russian) | Question (in Russian)  |
|---------------------------|-------------------|------------------------|
| `Petrovich.Genitive`      | Родительный       | Кого?                  |
| `Petrovich.Dative`        | Дательный         | Кому?                  |
| `Petrovich.Accusative`    | Винительный       | Кого?                  |
| `Petrovich.Instrumental`  | Творительный      | Кем?                   |
| `Petrovich.Prepositional` | Предложный        | О ком?                 |
