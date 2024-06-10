package main

import "fmt"

// type AliasCountry
type CountryInterface interface {
	// getInfo() string
	// getCapital() string
	getCurrency() string
	SetName(name string)
}
type Country struct {
	CountryInterface
	Name    string
	Capital string
}

func (c Country) getInfo() string {
	return fmt.Sprintf("%s is the capital of %s", c.getCapital(), c.getName())
}

func (c Country) getCapital() string {
	return c.Capital
}

func (c Country) getName() string {
	return c.Name
}

func (c *Country) SetName(name string) {
	c.Name = name
}

func (c Country) getCurrency() string {
	return "rupiah"
}

func Newcountry() *Country {
	return &Country{}
}

func main() {
	country := Country{
		Name:    "Indonesia",
		Capital: "Jakarta",
	}

	fmt.Println(country.getInfo())
	fmt.Println(country.getCapital())
	fmt.Println(country.getName())

	var country2 = Newcountry()
	country2.SetName("Singapore")
	fmt.Println(country2.getInfo())
}
