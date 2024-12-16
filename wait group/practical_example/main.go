package main

import (
	"fmt"
	"sync"
	"time"
)

type SocialPost struct {
	Likes   int64
	Id      int64
	Comment []string
	Friends int64
}
type Response struct {
	err  error
	data interface{}
}

func (s *SocialPost) getLikes() (int64, error) {
	time.Sleep(time.Second * 2)
	return 64, nil
}

func (s *SocialPost) getComments() ([]string, error) {
	time.Sleep(time.Second * 2)
	return []string{
		"Fucking Hilarious",
		"Whassup",
	}, nil
}

func (s *SocialPost) getFriends() (int64, error) {
	time.Sleep(time.Second * 2)
	return 45, nil
}

func (s *SocialPost) handleSocialPost() (*SocialPost, error) {
	comments, err := s.getComments()
	if err != nil {
		return nil, err
	}
	friends, err := s.getFriends()
	if err != nil {
		return nil, err
	}
	likes, err := s.getLikes()
	if err != nil {
		return nil, err
	}
	return &SocialPost{
		Likes:   likes,
		Id:      s.Id,
		Comment: comments,
		Friends: friends,
	}, nil
}

func main() {
	var wg sync.WaitGroup
	post := SocialPost{}
	go post.handleSocialPost(&wg)
	fmt.Println(post)
	fmt.Println(wg)
}
