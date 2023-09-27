package dog

import (
    "strings"
)

func WhenGrownUp(s string) string {
    //return "When it grows up, the Puppy says Not working!!"
    return "When it grows up, the Puppy says" + strings.ToUpper(s)
}

func WhenStillBaby() string {
    return "The Puppy says - am still a baby! Woof!\n"
}
