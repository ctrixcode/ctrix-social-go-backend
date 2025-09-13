package post_comment_likes

func LikeCommentRequestToPostCommentLike(req *LikeCommentRequest) *PostCommentLike {
	return &PostCommentLike{
		UserID:    req.UserID,
		CommentID: req.CommentID,
	}
}

func UnlikeCommentRequestToPostCommentLike(req *UnlikeCommentRequest) *PostCommentLike {
	return &PostCommentLike{
		UserID:    req.UserID,
		CommentID: req.CommentID,
	}
}

func PostCommentLikeToPostCommentLikeResponse(postCommentLike *PostCommentLike) *PostCommentLikeResponse {
	return &PostCommentLikeResponse{
		UserID:    postCommentLike.UserID,
		CommentID: postCommentLike.CommentID,
	}
}
