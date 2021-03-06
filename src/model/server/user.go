package server

type User struct {
	ID          int    `db:"id"`
	FirebaseUID string `db:"firebaseuid"`
	Points      int    `db:"points"`
	Cat1Level   int    `db:"category1_level"`
	Cat2Level   int    `db:"category2_level"`
	Cat3Level   int    `db:"category3_level"`
	Cat4Level   int    `db:"category4_level"`
	Cat5Level   int    `db:"category5_level"`
	Cat6Level   int    `db:"category6_level"`
}

func (u *User) GetCatLevels() []int {
	var catLevels []int
	catLevels = append(catLevels, u.Cat1Level)
	catLevels = append(catLevels, u.Cat2Level)
	catLevels = append(catLevels, u.Cat3Level)
	catLevels = append(catLevels, u.Cat4Level)
	catLevels = append(catLevels, u.Cat5Level)
	catLevels = append(catLevels, u.Cat6Level)
	return catLevels
}

func GetCatLevelColName(catID int) string {
	switch catID {
	case 1:
		return "category1_level"
	case 2:
		return "category2_level"
	case 3:
		return "category3_level"
	case 4:
		return "category4_level"
	case 5:
		return "category5_level"
	case 6:
		return "category6_level"
	default:
		return ""
	}
}
