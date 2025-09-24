package feeds

func MapPostWithAuthorToFeedPostDTO(data []PostWithAuthor) []FeedPostDTO {
	var feedPosts []FeedPostDTO
	for _, pwa := range data {
		feedPosts = append(feedPosts, FeedPostDTO{
			ID:        pwa.PostID,
			Content:   pwa.PostContent,
			CreatedAt: pwa.PostCreatedAt,
			Author: AuthorDTO{
				ID:         pwa.AuthorID,
				Username:   pwa.AuthorUsername,
				ProfilePic: pwa.AuthorProfilePic.String,
				Avatar:     pwa.AuthorAvatar.String,
			},
		})
	}
	return feedPosts
}
