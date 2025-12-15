package domain

type Users struct {
	ID        string `json:"id" bson:"id,omitempty"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Address   string `json:"address" bson:"address"`
	Email     string `json:"email" bson:"email"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Age       int    `json:"age" bson:"age"`
	Role      string `json:"role" bson:"role"`
}
