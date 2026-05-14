package collection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

	fmt.Printf("left length:%d,right length:%d", len(diffResult.LeftOnly), len(diffResult.RightOnly))
}

func TestIntersectBy(t *testing.T) {
	left := []User{
		{Id: 1, Name: "A"},
		{Id: 2, Name: "B"},
		{Id: 3, Name: "C"},
	}
	right := []User{
		{Id: 2, Name: "B2"},
		{Id: 4, Name: "D"},
	}

	intersected := IntersectBy(left, right, func(u User) int { return u.Id }, func(u User) int { return u.Id })

	assert.Equal(t, 1, len(intersected))
	assert.Equal(t, 2, intersected[0].Id)
	assert.Equal(t, "B", intersected[0].Name)
}

func TestJoinBy(t *testing.T) {
	left := []User{
		{Id: 1, Name: "A"},
		{Id: 2, Name: "B"},
	}
	right := []User{
		{Id: 2, Name: "B2"},
		{Id: 2, Name: "B3"},
		{Id: 3, Name: "C"},
	}

	joined := JoinBy(left, right, func(u User) int { return u.Id }, func(u User) int { return u.Id }, func(l User, r User) string {
		return fmt.Sprintf("%s-%s", l.Name, r.Name)
	})

	assert.Equal(t, []string{"B-B2", "B-B3"}, joined)
}

func TestInterceptBy(t *testing.T) {
	left := []User{
		{Id: 1, Name: "A"},
		{Id: 2, Name: "B"},
	}
	right := []User{
		{Id: 2, Name: "B2"},
	}

	intersected := InterceptBy(left, right, func(u User) int { return u.Id }, func(u User) int { return u.Id })

	assert.Equal(t, 1, len(intersected))
	assert.Equal(t, 2, intersected[0].Id)
}

func TestSelect(t *testing.T) {
	var users []User
	user1 := User{
		Id:   1,
		Name: "LoremIpSumXu",
	}

	user2 := User{
		Id:   2,
		Name: "are Q",
	}
	users = append(users, user1)
	users = append(users, user2)
	names := Select(users, func(u User) string {
		return u.Name
	})

	fmt.Printf("length of names : %d", len(names))

}

type User struct {
	Id   int
	Name string
}

func TestRange(t *testing.T) {
	slice := Range(0, 6)
	assert.True(t, len(slice) == 7)
}
