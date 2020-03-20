# MemDB

In-memory database with O(1) insert, delete and search.

**Note** for now it's not safe for concurrent use

### Usage

First, we define model to be used. It should implement `memdb.Entry` interface:

```Go
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
```

Key function required to implement Entry interface.

Initialize database and insert an entry:

```Go
user := User{
    Username:  "user1",
    FirstName: "user1 first name",
    LastName:  "user1 last name",
}

db := NewMemDb()
db.AddEntry(&user)
```

Search and delete:

```Go
// search for entry by key
entry1 := db.GetEntry("user1")
// cast entry to User type
entry1user := entry1.(*User)
// delete an entry
db.DeleteEntry("user1")
```

#### Groups

It's also possible to group entries. Let's say we have another model called Post:

```go
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
```

Using `AddGroupEntry`, `GetGroupEntry`, `DeleteGroupEntry` and `ListGroupEntry` we can manipulate items in a group.

```go
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
```

`GetGroupEntry` returns an entry in a group:

```go
db.GetGroupEntry("users", "user1")
db.GetGroupEntry("posts", "post1")
```

Returned interface can be casted to the origin type:

```go
entry := db.GetGroupEntry("users", "user1")
user := entry.(*User)
```

It's also possible to add different types to database without specifying any group:

```go
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
```
