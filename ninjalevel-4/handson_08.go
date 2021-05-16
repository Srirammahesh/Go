package main
import "fmt"
func main(){
	m:=map[string][]string{
		`bond_james`:[]string{ `Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`:[]string{ `James Bond`, `Literature`, `Computer Science`},
		`no_dr`:[]string{ `Being evil`, `Ice cream`, `Sunsets`},

	}
		fmt.Println(m)

	for _,v :=range m{
		for k,v1 :=range v{
			fmt.Println(k,v1)
		}
	}


}
