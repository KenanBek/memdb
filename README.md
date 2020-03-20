# MemDB

In-memory database with O(1) insert, delete and search.

Note: for now it's not safe for concurrent use.

## Example

First, we define model to be used. It should implement `memdb.Entry` interface:

```Go
    // User model.
    type User struct {
        Username  string `json:"username"`
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
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
```
