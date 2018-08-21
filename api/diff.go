package api

// DiffSet describes the difference between two companies
type DiffSet struct {
	Added, Modified, Removed Company
}

// Diff compares two company lists and returns a DiffSet
func Diff(oldC, newC Company) DiffSet {
	ds := DiffSet{
		Added:    Company{},
		Modified: Company{},
		Removed:  Company{},
	}

	for name, newVal := range newC {
		oldVal, found := oldC[name]
		if !found {
			ds.Added[name] = newVal
		} else if oldVal.Title != newVal.Title {
			ds.Modified[name] = newVal
		}
	}
	for name, oldVal := range oldC {
		_, found := newC[name]
		if !found {
			ds.Removed[name] = oldVal
		}
	}

	return ds
}
