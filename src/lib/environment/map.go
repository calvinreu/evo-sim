package environment

//Field contains the information of one square of the map
type Field struct {
	Creatures             []Creature
	Food, WaterPercentage float64
}

//Map contains the data about every part of the map
type Map struct {
	Fields [][]Field //[x][y]
}
