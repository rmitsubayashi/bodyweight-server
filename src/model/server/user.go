package server

type User struct {
	ID          int    `db:"id"`
	FirebaseUID string `db:"firebaseuid"`
	Points      int    `db:"points"`
	Cat1Exp     int    `db:"category1_experience"`
	Cat2Exp     int    `db:"category2_experience"`
	Cat3Exp     int    `db:"category3_experience"`
	Cat4Exp     int    `db:"category4_experience"`
	Cat5Exp     int    `db:"category5_experience"`
	Cat6Exp     int    `db:"category6_experience"`
}
