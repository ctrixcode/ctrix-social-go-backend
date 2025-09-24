package post_likes

func PostLikeToPostLikeResponse(postLike *PostLike) *PostLikeResponse {
	return &PostLikeResponse{
		UserID: postLike.UserID,
		PostID: postLike.PostID,
	}
}
