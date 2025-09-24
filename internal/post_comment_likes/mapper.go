package post_comment_likes

func PostCommentLikeToPostCommentLikeResponse(postCommentLike *PostCommentLike) *PostCommentLikeResponse {
	return &PostCommentLikeResponse{
		UserID:    postCommentLike.UserID,
		CommentID: postCommentLike.CommentID,
	}
}
