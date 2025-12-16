package domain

type Likes struct {
	ID        int    `json:"id" bson:"id"`
	UserId    string `json:"user_id" bson:"user_id"`
	PostId    *int   `json:"post_id" bson:"post_id"`
	CommentId *int   `json:"comment_id" bson:"comment_id"`
}
