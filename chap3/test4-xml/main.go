package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Customers struct {
	XMLName   xml.Name   `xml:"customers"`
	Customers []Customer `xml:"customer"`
}

type Customer struct {
	XMLName xml.Name `xml:"customer"`
	Address string   `xml:"address"`
	Name    string   `xml:"name"`
	Company Company  `xml:"company"`
}

type Company struct {
	XMLName  xml.Name `xml:"company"`
	Category string   `xml:"category,attr"`
	Name     string   `xml:"name"`
	Location string   `xml:"location"`
}

type XMLUtil struct{}

func NewXMLUtil() *XMLUtil {
	return &XMLUtil{}
}

func (xmlUtil *XMLUtil) readXML(path string) {
	xmlFormalFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFormalFile.Close()

	xmlBytes, err := io.ReadAll(xmlFormalFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var customers Customers
	err = xml.Unmarshal(xmlBytes, &customers)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	fmt.Println(customers)

	for i := 0; i < len(customers.Customers); i++ {
		fmt.Println("Customer Name is " + customers.Customers[i].Name)
		fmt.Println("Customer Address is " + customers.Customers[i].Address)
		fmt.Println("Customer Company Name is " + customers.Customers[i].Company.Name)
		fmt.Println("Customer Company Category is " + customers.Customers[i].Company.Category)
	}
}

func (xmlUtil *XMLUtil) writeXML(path string) {
	customers := Customers{
		Customers: []Customer{
			{
				Name:    "Jack Donner",
				Address: "Chicago",
				Company: Company{
					Name:     "Marigold",
					Location: "Denver",
					Category: "Silver",
				},
			},
		},
	}

	file, err := xml.MarshalIndent(customers, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	err = os.WriteFile(path, file, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func main() {
	xmlUtil := NewXMLUtil()
	xmlUtil.writeXML("customer.xml")
	xmlUtil.readXML("customer.xml")
}
