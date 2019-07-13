/* file dataStructs.go */

package main


import (
  "fmt"
  // "os"
  // "time"
)



/* Linked lists are linear collections of data el'ts where placement in memory
does not matter but each el't points to the next - streams of data - linked
lists with news feed
*/

/* Twitter Posts */
type Post struct {
  body string
  publishedDate int64
  next *Post // link to next post
}

/*
Feeds have lengths and linked posts
*/

type Feed struct{
  length int
  start *Post
  end *Post
}

/* one way to optimize insertions/append to a linked list - previously O(n)
complexity - is to include an end pointer (will take up more memory), but
appending a post to the feed (always added to the end)
--> If the feed is empty, then start and end are the same newPost
*/

func (f *Feed) Insert(newPost *Post){

}

func (f *Feed) Append(newPost *Post){
  if f.length == 0{
    f.start = newPost
    f.end = newPost
  } else {
    lastPost := f.end // lastPost is now the end pointing to the newPost
    lastPost.next = newPost
    f.end = newPost // update end of the Feed f
  }
}

// func (f *Feed) Delete(publishedDate int64){
//
// }

/*
Initializing a struct with assignment operations
*/

func main(){
  name := "blank"
  age := 18
  fmt.Printf("Hello my name is %s and I'm %d years old",name,age)
}
