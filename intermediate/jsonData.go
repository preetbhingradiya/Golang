package intermediate

import (
	"encoding/json"
)

type person struct {
	Name     string  `json:"name" db:"first_name"` // This field will be included in JSON output and mapped to the "first_name" column in the database
	Age      int     `json:"age,omitempty"`        // This field will be omitted/removed if it has a zero/empty value
	Password string  `json:"-"`                    // This field will be ignored in JSON output
	Address  address `json:"address"`
}

type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func JsonData() {

	person1 := person{
		Name: "John Doe",
		// Age:      30,
		Password: "wmdi3d9nddjj39j",
		Address: address{
			City:  "New York",
			State: "NY",
		},
	}

	personData, _ := json.Marshal(person1)

	jsonString := string(personData)
	infoLogger.Println("JSON String:", jsonString)

	jsonData := `{"name":"Jane Smith","age":25,"address":{"city":"Los Angeles","state":"CA"}}`

	var person2 person
	err := json.Unmarshal([]byte(jsonData), &person2)
	if err != nil {
		panic(err)
	}

	infoLogger.Println("Person 2:", person2)

	var person3 map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &person3)
	if err != nil {
		panic(err)
	}

	infoLogger.Println("Person 3:", person3)
	infoLogger.Println("Person 3 Name:", person3["name"])
	infoLogger.Println("Person 3 City:", person3["address"].(map[string]interface{})["city"])
}
