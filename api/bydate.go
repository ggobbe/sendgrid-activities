package api

// ByDate implements sort.Interface for []Activity based on the Created field.
type ByDate []Activity

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Created < a[j].Created }
