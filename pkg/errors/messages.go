package errors

type ErrorType struct {
	Code    string
	Message string
}

var (
	ErrBadRequest          = ErrorType{Code: "BAD_REQUEST", Message: "Bad request"}
	ErrValidationFailed    = ErrorType{Code: "VALIDATION_FAILED", Message: "Validation failed"}
	ErrUnauthorized        = ErrorType{Code: "UNAUTHORIZED", Message: "Unauthorized: User not authenticated."}
	ErrNotFound            = ErrorType{Code: "NOT_FOUND", Message: "Resource not found."}
	ErrInternalServerError = ErrorType{Code: "INTERNAL_SERVER_ERROR", Message: "Internal server error"}
	ErrSomethingWentWrong  = ErrorType{Code: "SOMETHING_WENT_WRONG", Message: "Something went wrong"}
	ErrExternalApiError    = ErrorType{Code: "EXTERNAL_API_ERROR", Message: "External API error."}
	ErrNoFileData          = ErrorType{Code: "NO_FILE_DATA_ERROR", Message: "No file data received"}
	ErrFileProcessingError = ErrorType{Code: "FILE_PROCESSING_ERROR", Message: "Error processing file"}
	ErrUsageLimitExceeded  = ErrorType{Code: "USAGE_LIMIT_EXCEEDED_ERROR", Message: "Usage limit exceeded."}

	// Post Errors
	FailedToCreatePost = ErrorType{Code: "FAILED_TO_CREATE_POST", Message: "Failed to create post"}
	InvalidPostID      = ErrorType{Code: "INVALID_POST_ID", Message: "Invalid post ID"}
	PostNotFound       = ErrorType{Code: "POST_NOT_FOUND", Message: "Post not found"}
	FailedToUpdatePost = ErrorType{Code: "FAILED_TO_UPDATE_POST", Message: "Failed to update post"}

	// Post Like Errors
	PostAlreadyLiked   = ErrorType{Code: "POST_ALREADY_LIKED", Message: "Post already liked by this user"}
	FailedToLikePost   = ErrorType{Code: "FAILED_TO_LIKE_POST", Message: "Failed to like post"}
	PostNotLiked       = ErrorType{Code: "POST_NOT_LIKED", Message: "Post not liked by this user"}
	FailedToUnlikePost = ErrorType{Code: "FAILED_TO_UNLIKE_POST", Message: "Failed to unlike post"}

	// Post Comment Like Errors
	CommentAlreadyLiked   = ErrorType{Code: "COMMENT_ALREADY_LIKED", Message: "Comment already liked by this user"}
	FailedToLikeComment   = ErrorType{Code: "FAILED_TO_LIKE_COMMENT", Message: "Failed to like comment"}
	CommentNotLiked       = ErrorType{Code: "COMMENT_NOT_LIKED", Message: "Comment not liked by this user"}
	FailedToUnlikeComment = ErrorType{Code: "FAILED_TO_UNLIKE_COMMENT", Message: "Failed to unlike comment"}

	// Follow/unfollow errors
	FollowedAlready   = ErrorType{Code: "FOLLOWED_ALREADY", Message: "You are already following this user"}
	FailedToFollow    = ErrorType{Code: "FAILED_TO_FOLLOW", Message: "Failed to follow user"}
	UnfollowedAlready = ErrorType{Code: "UNFOLLOWED_ALREADY", Message: "You are already unfollowing this user"}
	FailedToUnfollow  = ErrorType{Code: "FAILED_TO_UNFOLLOW", Message: "Failed to unfollow user"}
)

var (
	ErrNoToken                = ErrorType{Code: "NO_TOKEN_ERROR", Message: "No token provided"}
	ErrFailedCreateUpdateUser = ErrorType{Code: "FAILED_CREATE_UPDATE_USER_ERROR", Message: "Failed to create or update user"}
	ErrInvalidToken           = ErrorType{Code: "INVALID_TOKEN_ERROR", Message: "Authentication failed. Please log in again."}
	ErrExpiredToken           = ErrorType{Code: "EXPIRED_REFRESH_TOKEN_ERROR", Message: "Your session has expired. Please log in again."}
	ErrInvalidRefreshToken    = ErrorType{Code: "INVALID_REFRESH_TOKEN", Message: "Invalid refresh token"}
	ErrUserAgentMismatch      = ErrorType{Code: "USER_AGENT_MISMATCH_ERROR", Message: "Authentication failed. Please log in again."}
	ErrInvalidCredentials     = ErrorType{Code: "INVALID_CREDENTIALS", Message: "Invalid credentials"}
	ErrUserAlreadyExists      = ErrorType{Code: "USER_ALREADY_EXISTS", Message: "User already exists"}
	ErrFailedToCreateUser     = ErrorType{Code: "FAILED_TO_CREATE_USER", Message: "Failed to create user"}
	ErrFailedToGenerateToken  = ErrorType{Code: "FAILED_TO_GENERATE_TOKEN", Message: "Failed to generate token"}
)
