package puppy

import (
	"fmt"

	"github.com/timerunnerx9/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func Fromv1() {
	fmt.Println("this is version 1.1")
}

func Fromv2() {
	fmt.Println("this is version 1.2")
}
