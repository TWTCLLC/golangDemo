package dto

type Messages struct {
	UserID  string   `bson:"userID"`
	Message []string `bson:"message"`
}
