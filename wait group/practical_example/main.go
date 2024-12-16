package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type SocialPost struct {
	Likes   int64
	Id      int64
	Comment []string
	Friends int64
}

func (s *SocialPost) getLikes(wg *sync.WaitGroup) (int64, error) {
	time.Sleep(time.Second * 2)
	wg.Done()
	return 64, nil
}

func (s *SocialPost) getComments(wg *sync.WaitGroup) ([]string, error) {
	time.Sleep(time.Second * 2)
	wg.Done()
	return []string{
		"Fucking Hilarious",
		"Whassup",
	}, nil
}

func (s *SocialPost) getFriends(wg *sync.WaitGroup) (int64, error) {
	time.Sleep(time.Second * 2)
	wg.Done()
	return 45, nil
}

func (s *SocialPost) handleSocialPost(post *SocialPost, wg *sync.WaitGroup) {
	go func() {
		comments, err := s.getComments(wg)
		post.Comment = comments
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	go func() {
		friends, err := s.getFriends(wg)
		post.Friends = friends
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	go func() {
		likes, err := s.getLikes(wg)
		post.Likes = likes
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	wg.Add(3)
}

func main() {
	var wg sync.WaitGroup
	post := SocialPost{
		Id: 1,
	}
	post.handleSocialPost(&post, &wg)

	wg.Wait()
	fmt.Println(post)
}
