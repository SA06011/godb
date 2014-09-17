package data

// record represents a row in table
type record struct {
	TableId  int         // unique table id
	PrimeKey interface{} // unique primary value
	ObjId    uint64      // automated generation id for each row
	Pos      []int       // record position's offset in cache
	Codec     map[int]func

}
