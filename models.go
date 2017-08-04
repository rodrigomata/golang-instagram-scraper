package main

func (post *TopPosts) AddItem(item Post) []Post {
  post.Posts = append(post.Posts, item)
  return post.Posts
}
