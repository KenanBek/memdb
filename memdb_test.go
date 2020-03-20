package memdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// User model.
type User struct {
	Username  string
	FirstName string
	LastName  string
}

// Key returns unique username of the user.
func (u *User) Key() string {
	return u.Username
}

// Post model.
type Post struct {
	Slug    string
	Author  string
	Title   string
	Content string
}

// Key returns unique slug of the post.
func (p *Post) Key() string {
	return p.Slug
}

func TestNewMemDb(t *testing.T) {
	user := User{
		Username:  "user1",
		FirstName: "user1 first name",
		LastName:  "user1 last name",
	}

	db := NewMemDb()
	db.AddEntry(&user)

	assert.NotNil(t, db.GetEntry("user1"))
}

func TestMemDb_AddEntry(t *testing.T) {
	user1 := User{
		Username:  "user1",
		FirstName: "user1 first name",
		LastName:  "user1 last name",
	}
	post1 := Post{
		Slug:    "post1",
		Author:  "user1",
		Title:   "post1 title",
		Content: "post1 content",
	}

	db := NewMemDb()
	db.AddEntry(&user1)
	db.AddEntry(&post1)

	assert.NotNil(t, db.GetEntry("user1"))
	assert.NotNil(t, db.GetEntry("post1"))
}

func TestMemDb_GetEntry(t *testing.T) {
	user1 := User{
		Username:  "user1",
		FirstName: "user1 first name",
		LastName:  "user1 last name",
	}
	user2 := User{
		Username:  "user2",
		FirstName: "user2 first name",
		LastName:  "user2 last name",
	}

	db := NewMemDb()
	db.AddEntry(&user1)
	db.AddEntry(&user2)

	entry1 := db.GetEntry("user1")
	entry1user := entry1.(*User)
	assert.Equal(t, "user1", entry1user.Username)
	assert.Equal(t, "user1 first name", entry1user.FirstName)
	assert.Equal(t, "user1 last name", entry1user.LastName)
}

func TestMemDb_DeleteEntry(t *testing.T) {
	user1 := User{
		Username:  "user1",
		FirstName: "user1 first name",
		LastName:  "user1 last name",
	}
	user2 := User{
		Username:  "user2",
		FirstName: "user2 first name",
		LastName:  "user2 last name",
	}

	db := NewMemDb()
	db.AddEntry(&user1)
	db.AddEntry(&user2)

	entry2 := db.GetEntry("user2")
	entry2user := entry2.(*User)
	assert.Equal(t, "user2", entry2user.Username)
	assert.Equal(t, "user2 first name", entry2user.FirstName)
	assert.Equal(t, "user2 last name", entry2user.LastName)

	db.DeleteEntry("user2")
	entry2 = db.GetEntry("user2")
	assert.Nil(t, entry2)
}

func TestMemDb_AddGroupEntry(t *testing.T) {
	user1 := User{
		Username:  "user1",
		FirstName: "user1 first name",
		LastName:  "user1 last name",
	}
	post1 := Post{
		Slug:    "post1",
		Author:  "user1",
		Title:   "post1 title",
		Content: "post1 content",
	}
	post2 := Post{
		Slug:    "post2",
		Author:  "user2",
		Title:   "post2 title",
		Content: "post2 content",
	}

	db := NewMemDb()
	db.AddGroupEntry("users", &user1)
	db.AddGroupEntry("posts", &post1)
	db.AddGroupEntry("posts", &post2)

	assert.NotNil(t, db.GetGroupEntry("users", "user1"))
	assert.NotNil(t, db.GetGroupEntry("posts", "post1"))
	assert.NotNil(t, db.GetGroupEntry("posts", "post2"))

	assert.Equal(t, "post2 title", db.GetGroupEntry("posts", "post2").(*Post).Title)

	db.DeleteGroupEntry("posts", "post1")

	assert.Nil(t, db.GetGroupEntry("posts", "post1"))
}
