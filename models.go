package main

// AddItem adds a Post to an array of Posts
func (post *TopPosts) AddItem(item Post) []Post {
	post.Posts = append(post.Posts, item)
	return post.Posts
}
