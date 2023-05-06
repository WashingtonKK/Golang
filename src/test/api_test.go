package api_test

//Test to compare contents of the yaml with what we have

//Import the necessary packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

type data struct {
	User struct {
		Name  string `Washington`
		Email string `123.com`
		Error string `wkk`
		id    int    `3`
	}
	Requires struct {
		Order string `food`
	}
}

type ref struct {
	User struct {
		Name  string `yaml:"Name"`
		Email string `yaml:"Email"`
		Error string `yaml:"Error"`
		id    int    `yaml:"id"`
	}
	Requires struct {
		Order string `yaml:"Order"`
	}
}

// Implement the required packages.
func TestResponse(t *testing.T) {
	//Read the yaml file
	dataFile, err := ioutil.ReadFile("data.yaml")
	if err != nil {
		log.Fatalf("Unable to read from file %v", err)
	}

	dataRead := data{}
	Ref := ref{}
	err = yaml.Unmarshal(dataFile, &dataRead)

	if err != nil {
		fmt.Printf("Error reading file %v", err)
	}
	assert.Equal(t, dataRead.User.id, Ref.User.id, "Invalid id")
	assert.Equal(t, dataRead.User.Name, Ref.User.Name, "Invalid Name")
	//fmt.Printf("Tuko sete")
}
