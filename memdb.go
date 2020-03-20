package memdb

import "log"

// Entry defines model for in-memory database.
type Entry interface {
	Key() string
}

// NewMemDb creates and returns a new instance of MemDb.
func NewMemDb() *MemDb {
	return &MemDb{
		containers: map[string]*MemContainer{},
	}
}

// MemContainer is a in-memory storage model.
type MemContainer struct {
	entries    []Entry
	kvsEntries map[string]int
}

func newMemContainer() *MemContainer {
	return &MemContainer{
		entries:    []Entry{},
		kvsEntries: map[string]int{},
	}
}

// MemDb is a in-memory database.
type MemDb struct {
	containers map[string]*MemContainer
}

func (db *MemDb) getContainer(g string) *MemContainer {
	container, ok := db.containers[g]
	if !ok {
		db.containers[g] = newMemContainer()
		container = db.containers[g]
	}

	return container
}

// AddEntry entry to unassigned group.
func (db *MemDb) AddEntry(e Entry) Entry {
	log.Println("AddEntry", e)

	return db.AddGroupEntry("", e)
}

// AddGroupEntry entry to the given group.
func (db *MemDb) AddGroupEntry(g string, e Entry) Entry {
	log.Println("AddGroupEntry", g, e)

	container := db.getContainer(g)

	ind, ok := container.kvsEntries[e.Key()]
	if !ok {
		container.kvsEntries[e.Key()] = len(container.entries)
		container.entries = append(container.entries, e)
		return container.entries[len(container.entries)-1]
	}

	container.entries[ind] = e
	return container.entries[ind]
}

// GetEntry returns an entry in unassigned group.
func (db *MemDb) GetEntry(key string) Entry {
	log.Println("GetEntry", key)

	return db.GetGroupEntry("", key)
}

// GetGroupEntry returns an entry in the given group.
func (db *MemDb) GetGroupEntry(group, key string) Entry {
	log.Println("GetGroupEntry", group, key)

	container := db.getContainer(group)

	ind, ok := container.kvsEntries[key]
	if ok {
		return container.entries[ind]
	}

	return nil
}

// DeleteEntry deletes entry in unassigned group.
func (db *MemDb) DeleteEntry(key string) {
	log.Println("DeleteEntry", key)

	db.DeleteGroupEntry("", key)
}

// DeleteGroupEntry entry in the given group.
func (db *MemDb) DeleteGroupEntry(group, key string) {
	log.Println("DeleteGroupEntry", group, key)

	container := db.getContainer(group)

	if ind, ok := container.kvsEntries[key]; ok {
		container.entries[ind] = container.entries[len(container.entries)-1]
		container.entries = container.entries[:len(container.entries)-1]

		delete(container.kvsEntries, key)
	}
}

// ListEntries returns all entires in unassigned group.
func (db *MemDb) ListEntries() []Entry {
	log.Println("ListEntries")

	return db.ListGroupEntries("")
}

// ListGroupEntries returns all entries in the given group.
func (db *MemDb) ListGroupEntries(group string) []Entry {
	log.Println("ListGroupEntries", group)

	container := db.getContainer(group)

	return container.entries
}
