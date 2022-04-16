package collection

import (
	"fmt"
	"testing"
)

func TestDiff(t *testing.T) {
	var left []User
	var right []User

	user1 := User{
		Id:   1,
		Name: "LoremIpSumXu",
	}

	user2 := User{
		Id:   2,
		Name: "are Q",
	}

	left = append(append(left, user1), user2)
	right = append(right, user2)

	diffResult := Diff(left, right, func(u User) int { return u.Id }, func(u User) int { return u.Id })

	fmt.Printf("left length:%d,right length:%d",len(diffResult.LeftOnly),len(diffResult.RightOnly))
}

type User struct {
	Id   int
	Name string
}
