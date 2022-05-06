package principle

import (
	"fmt"
	"time"
)

// Interface 1
type AuthorDetails interface {
	Details()
	Sums() (int, error)
}

// Interface 2
type AuthorArticles interface {
	Articles()
}

// Interface 3

// Interface 3 embedded with
// interface 1 and 2's methods
type FinalDetails interface {
	AuthorDetails
	AuthorArticles
	Cleans()
}

type Author struct {
	A_name    string
	Branch    string
	College   string
	Year      int
	Salary    int
	Particles int
	Tarticles int
}

type NewPrincipleTwoInit struct {
	Author Author
}

func NewPrincipleTwo(author Author) FinalDetails {
	return &NewPrincipleTwoInit{
		Author: author,
	}
}

// Implementing method of
// the interface 1
func (a *NewPrincipleTwoInit) Details() {
	fmt.Printf("Author Name: %s", a.Author.A_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.Author.Branch, a.Author.Year)
	fmt.Printf("\nCollege Name: %s", a.Author.College)
	fmt.Printf("\nSalary: %d", a.Author.Salary)
	fmt.Printf("\nPublished articles: %d", a.Author.Particles)
}

// Implementing method
// of the interface 2
func (a *NewPrincipleTwoInit) Articles() {
	pendingarticles := a.Author.Tarticles - a.Author.Particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

func (a *NewPrincipleTwoInit) Sums() (int, error) {
	currentYear := time.Now().Year()
	diffYear := currentYear - a.Author.Year
	i := 0
	totalSalary := 0
	for i < diffYear {
		totalSalary += a.Author.Salary
		i++
	}
	fmt.Printf("\n Total Salary from year %d to current year %d : %d", a.Author.Year, currentYear, totalSalary)
	return totalSalary, nil
}

func (a *NewPrincipleTwoInit) Cleans() {
	fmt.Printf("\n-------------- %s --------------\n", "End Process")
}
