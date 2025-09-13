package post_comments

func ToPostComment(req *CreateCommentRequest) *PostComment {
	return &PostComment{
		PostID:    req.PostID.String(),
		CreatorID: req.UserID.String(),
		Content:   &req.Content,
	}
}

func ToUpdatePostComment(comment *PostComment, req *UpdateCommentRequest) *PostComment {
	comment.Content = &req.Content
	return comment
}
