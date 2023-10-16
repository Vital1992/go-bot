package enums

// Category - Custom type to hold value for category ranging from 1-7
type Category int

// Declare related constants for each weekday starting with index 1
const (
	Kitchen    Category = iota + 1 // EnumIndex = 1
	Computers                      // EnumIndex = 2
	Sound                      	   // EnumIndex = 3
	TV                    		   // EnumIndex = 4
	CellPhones                     // EnumIndex = 5
	Tablets                        // EnumIndex = 6
	Games                     	   // EnumIndex = 7
)

// String - Creating common behavior - give the type a String function
func (c Category) String() string {
	return [...]string{"Kitchen", "Computers", "Sound", "TV", "CellPhones", "Tablets", "Games"}[c-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (c Category) EnumIndex() int {
	return int(c)
}
