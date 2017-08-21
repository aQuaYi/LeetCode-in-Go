package main

func main() {
	categories := []string{
		"Algorithms",
		// "database",
		"Draft",
		// "operating-system",
		// "shell",
		// "system-design",
	}

	solvedProblems := Solveds{}
	cs := Categories{}
	total := &Category{
		Name: "Total",
	}

	for i := range categories {
		c := getCategory(categories[i])

		solvedProblems = append(solvedProblems, c.check()...)

		total.update(c)
		cs = append(cs, c)
	}
	cs = append(cs, total)

	makeREADME(cs, solvedProblems)
}
