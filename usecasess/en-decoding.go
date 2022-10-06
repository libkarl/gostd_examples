package usecases

import (
	"encoding/json"
	"fmt"
	"time"
)

type Dog struct {
	ID     int
	Name   string
	Breed  string
	BornAt time.Time
}

type JSONDog struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	BornAt int64  `json:"born_at"`
}

// A constructor that takes in a Dog and returns a JSONDog we get the following code.
func NewJSONDog(dog Dog) JSONDog {
	return JSONDog{
		dog.ID,
		dog.Name,
		dog.Breed,
		dog.BornAt.Unix(),
	}
}

func EncodeDecEx() {
	dog := Dog{1, "bowser", "husky", time.Now()}
	fmt.Println(dog)
	b, err := json.Marshal(NewJSONDog(dog))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// unmarshall
	var jsonDog JSONDog
	err = json.Unmarshal(b, &jsonDog)
	fmt.Println(jsonDog.Dog())

}

func (jd JSONDog) Dog() Dog {
	return Dog{
		jd.ID,
		jd.Name,
		jd.Breed,
		time.Unix(jd.BornAt, 0),
	}
}
