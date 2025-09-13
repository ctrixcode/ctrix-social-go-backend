package post_likes

func LikePostRequestToPostLike(req *LikePostRequest) *PostLike {
	return &PostLike{
		UserID: req.UserID,
		PostID: req.PostID,
	}
}

func UnlikePostRequestToPostLike(req *UnlikePostRequest) *PostLike {
	return &PostLike{
		UserID: req.UserID,
		PostID: req.PostID,
	}
}

func PostLikeToPostLikeResponse(postLike *PostLike) *PostLikeResponse {
	return &PostLikeResponse{
		UserID: postLike.UserID,
		PostID: postLike.PostID,
	}
}
