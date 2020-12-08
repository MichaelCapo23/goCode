package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

// func main() {
// 	c := make(chan string)
// 	go count("sheep", c)

// 	for msg := range c {
// 		fmt.Println(msg)
// 	}
// }

// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}

// 	close(c)
// }

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		for {
// 			c1 <- "Every 500"
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			c2 <- "Every 2 seconds"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()

// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

// func main() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)

// 	go worker(jobs, results)
// 	go worker(jobs, results)
// 	go worker(jobs, results)
// 	go worker(jobs, results)

// 	for i := 0; i < 100; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)
// 	for j := 0; j < 100; j++ {
// 		fmt.Println(<-results)
// 	}
// }

// func worker(jobs <-chan int, results chan<- int) {
// 	for n := range jobs {
// 		results <- fib(n)
// 	}
// }

// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	return fib(n-1) + fib(n-2)
// }

// func mergeSort(items []int) []int {
// 	if len(items) == 1 {
// 		return items
// 	}

// 	middle := int(len(items) / 2)
// 	var (
// 		left  = make([]int, middle)
// 		right = make([]int, len(items)-middle)
// 	)
// 	for i := 0; i < len(items); i++ {
// 		if i < middle {
// 			left[i] = items[i]
// 		} else {
// 			right[i-middle] = items[i]
// 		}
// 	}

// 	return merge(mergeSort(left), mergeSort(right))
// }

// func merge(left, right []int) (result []int) {
// 	result = make([]int, len(left)+len(right))

// 	i := 0
// 	for len(left) > 0 && len(right) > 0 {
// 		if left[0] < right[0] {
// 			result[i] = left[0]
// 			left = left[1:]
// 		} else {
// 			result[i] = right[0]
// 			right = right[1:]
// 		}
// 		i++
// 	}

// 	for j := 0; j < len(left); j++ {
// 		result[i] = left[j]
// 		i++
// 	}
// 	for j := 0; j < len(right); j++ {
// 		result[i] = right[j]
// 		i++
// 	}
// 	return
// }

// Quick Sort in Golang

// func main() {

// 	slice := generateSlice(20)
// 	fmt.Println("\n--- Unsorted --- \n\n", slice)
// 	quicksort(slice)
// 	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
// }

// // Generates a slice of size, size filled with random numbers
// func generateSlice(size int) []int {

// 	slice := make([]int, size, size)
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < size; i++ {
// 		slice[i] = rand.Intn(999) - rand.Intn(999)
// 	}
// 	return slice
// }

// func quicksort(a []int) []int {
// 	if len(a) < 2 {
// 		return a
// 	}

// 	left, right := 0, len(a)-1

// 	pivot := rand.Int() % len(a)

// 	a[pivot], a[right] = a[right], a[pivot]

// 	for i := range a {
// 		if a[i] < a[right] {
// 			a[left], a[i] = a[i], a[left]
// 			left++
// 		}
// 	}

// 	a[left], a[right] = a[right], a[left]

// 	quicksort(a[:left])
// 	quicksort(a[left+1:])

// 	return a
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseList(head *ListNode) *ListNode {
// 	current := head
// 	var next *ListNode = nil
// 	var prev *ListNode = nil

// 	for current != nil {
// 		next = current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}
// 	return prev
// }

// type Node struct {
// 	prev *Node
// 	next *Node
// 	key  interface{}
// }

// type List struct {
// 	head *Node
// 	tail *Node
// }

// func (L *List) Insert(key interface{}) {
// 	list := &Node{
// 		next: L.head,
// 		key:  key,
// 	}

// 	if L.head != nil {
// 		L.head.prev = list
// 	}
// 	L.head = list

// 	l := L.head
// 	for l.next != nil {
// 		l = l.next
// 	}
// 	L.tail = l
// }

// func (l *List) Display() {
// 	list := l.head
// 	for list != nil {
// 		fmt.Printf("%+v ->", list.key)
// 		list = list.next
// 	}
// 	fmt.Println()
// }

// func Display(list *Node) {
// 	for list != nil {
// 		fmt.Printf("%v ->", list.key)
// 		list = list.next
// 	}
// 	fmt.Println()
// }

// func ShowBackwards(list *Node) {
// 	for list != nil {
// 		fmt.Printf("%v <-", list.key)
// 		list = list.prev
// 	}
// 	fmt.Println()
// }

// func (l *List) Reverse() {
// 	curr := l.head
// 	var prev *Node
// 	l.tail = l.head

// 	for curr != nil {
// 		next := curr.next
// 		curr.next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	l.head = prev
// 	Display(l.head)
// }

// func main() {
// 	link := List{}
// 	link.Insert(5)
// 	link.Insert(9)
// 	link.Insert(13)
// 	link.Insert(22)
// 	link.Insert(28)
// 	link.Insert(36)

// 	fmt.Println("\n==============================\n")
// 	fmt.Printf("Head: %v\n", link.head.key)
// 	fmt.Printf("Tail: %v\n", link.tail.key)
// 	link.Display()
// 	fmt.Println("\n==============================\n")
// 	fmt.Printf("head: %v\n", link.head.key)
// 	fmt.Printf("tail: %v\n", link.tail.key)
// 	link.Reverse()
// 	fmt.Println("\n==============================\n")
// }

// const arraySize = 7

// //HashTable structure
// type HashTable struct {
// 	array [arraySize]*bucket
// }

// //bucket structure (linked list)
// type bucket struct {
// 	Head *bucketNode
// }

// //bucketNode structure
// type bucketNode struct {
// 	Key  string
// 	Next *bucketNode
// }

// //Insert (take a key and insert it into the hash table array)
// func (h *HashTable) Insert(key string) {
// 	index := hash(key)
// 	h.array[index].insert(key)
// }

// //Search
// func (h *HashTable) Search(key string) bool {
// 	index := hash(key)
// 	return h.array[index].search(key)
// }

// //Delete
// func (h *HashTable) Delete(key string) {
// 	index := hash(key)
// 	h.array[index].delete(key)
// }

// //hash
// func hash(key string) int {
// 	sum := 0
// 	for _, v := range key {
// 		sum += int(v)
// 	}

// 	return sum % arraySize
// }

// //insert (insert a new bucketNode into the current index inside the hashTables array)
// func (b *bucket) insert(k string) {
// 	if !b.search(k) {
// 		current := &bucketNode{Key: k}
// 		current.Next = b.Head
// 		b.Head = current
// 	} else {
// 		fmt.Println("already exists")
// 	}
// }

// //search
// func (b *bucket) search(k string) bool {
// 	current := b.Head

// 	for current != nil {
// 		if current.Key == k {
// 			return true
// 		}
// 		current = current.Next
// 	}
// 	return false
// }

// //delete
// func (b *bucket) delete(k string) {

// 	if b.Head.Key == k {
// 		b.Head = b.Head.Next
// 		return
// 	}

// 	current := b.Head
// 	for current.Next != nil {
// 		if current.Next.Key == k {
// 			current.Next = current.Next.Next
// 			return
// 		}
// 		current = current.Next
// 	}
// }

// //Init (create a new bucket into each index of the HashTables indexes (currently nil))
// func Init() *HashTable {
// 	result := &HashTable{}

// 	for i := range result.array {
// 		result.array[i] = &bucket{}
// 	}
// 	return result
// }

// func main() {
// 	//set the HashTable value to the result of the init function (a hashMap that has each index populated with a bucket)
// 	HashTable := Init()

// 	List := []string{
// 		"ERIC",
// 		"KENNY",
// 		"KYLE",
// 		"STAN",
// 		"RANDY",
// 		"BUTTERS",
// 		"TOKEN",
// 	}

// 	for _, v := range List {
// 		HashTable.Insert(v)
// 	}

// 	HashTable.Delete("STAN")
// 	fmt.Println("STAN", HashTable.Search("STAN"))
// 	fmt.Println("KENNY", HashTable.Search("KENNY"))
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func isPalindrome(head *ListNode) bool {
// 	current := head

// 	if head == nil || head.Next == nil {
// 		return true
// 	}

// 	var (
// 		slow *ListNode = head
// 		fast *ListNode = head
// 	)

// 	for fast.Next != nil && fast.Next.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}

// 	slow.Next = reverse(slow.Next)
// 	slow = slow.Next

// 	for slow != nil {
// 		if slow.Val != current.Val {
// 			return false
// 		}

// 		slow = slow.Next
// 		current = current.Next
// 	}

// 	return true

// }

// func reverse(head *ListNode) *ListNode {
// 	var (
// 		prev    *ListNode
// 		next    *ListNode
// 		current *ListNode = head
// 	)
// 	for current != nil {
// 		next = current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}
// 	return prev
// }

// func (head *ListNode) Insert(val int) {

// 	for head != nil {
// 		if head.Next == nil {
// 			head.Next = &ListNode{Val: val}
// 			return
// 		}
// 		head = head.Next
// 	}

// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func (head *ListNode) Insert(val int) {

// 	current := head

// 	for current != nil {
// 		if current.Next == nil {
// 			current.Next = &ListNode{Val: val}
// 			return
// 		}
// 		current = current.Next
// 	}

// }

// func (head *ListNode) Reverse() *ListNode {
// 	var (
// 		next    *ListNode = nil
// 		prev    *ListNode = nil
// 		current *ListNode = head
// 	)

// 	for current != nil {
// 		next = current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}
// 	return prev
// }

// func main() {
// 	head := &ListNode{}
// 	head.Insert(1)
// 	head.Insert(2)
// 	head.Insert(3)
// 	head.Insert(4)
// 	head.Insert(5)
// 	head.Insert(6)
// 	res := head.Reverse()
// 	fmt.Println(res)
// }

// type plane struct {
// 	Fuel bool
// }

// func (p plane) AddFuel() {
// 	p.Fuel = true
// }

// //AddFuel
// func AddFuel(p plane) plane {
// 	p.Fuel = true
// 	return p
// }

// func main() {

// 	var fuel = plane{}
// 	fuel.AddFuel()
// 	fmt.Println(fuel)

// 	var new2 = plane{}
// 	res := AddFuel(new2)
// 	fmt.Println(res)
// }

// Design a service that is sent urls and must grab their contents.

// func GetAndParse(url string) error {
// 	return nil
// }

// func worker(ctx context.Context, ch chan string) {
// 	for {
// 		select {
// 		case val := <-ch:
// 			GetAndParse(val)
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }

// func main() {
// 	ch := make(chan string)
// 	ctx := context.Background()
// 	go worker(ctx, ch)
// 	go worker(ctx, ch)
// 	go worker(ctx, ch)
// 	go worker(ctx, ch)
// }

// func sleepAndTalk(ctx context.Context, s string) {
// 	time.Sleep(5 * time.Second)
// 	fmt.Println(s)
// }

// func main() {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)

// 	time.AfterFunc(time.Second, cancel)

// 	sleepAndTalk(ctx, "Hello")
// }

// func main() {
// ctx := context.Background()
// ctx := context.TODO()
// _ = context.WithValue(ctx, "myKey", "123")

//*******withCancel example*******

// ctx2, cancelFunc := context.WithCancel(ctx)
// log.Print(ctx2.Err())
// cancelFunc()
// log.Print(ctx2.Err() == context.Canceled)

//*******withDeadline example*******

// t := time.Now().Add(1 * time.Second)
// ctx2, cancelFunc := context.WithDeadline(ctx, t)
// cancelFunc() //can also use a cancel function with this too or just set a deadline either way
// time.Sleep(2 * time.Second)
// log.Print(ctx2.Err() == context.DeadlineExceeded)

//*******withTimeout  *******
// 	timeout := 1 * time.Second
// 	ctx2, _ := context.WithTimeout(ctx, timeout) //can again use a cancel function
// 	time.Sleep(2 * time.Second)

// 	ctx3 := context.WithValue(ctx2, "whatever", 456)
// 	fn(ctx2)
// 	fn(ctx3)
// }

// func fn(ctx context.Context) {
// 	select {
// 	case <-time.After(3 * time.Second):
// 		log.Print("after 3 sec")
// 	case <-ctx.Done():
// 		log.Print("Done")
// 	}
// }

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Podcast struct {
	Title    string  `redis:"title"`
	Creator  string  `redis:"creator"`
	Category string  `redis:"category"`
	Fee      float64 `redis:"membership_fee"`
}

func main() {
	// ctx := context.Background()
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkErr(err)
	defer conn.Close()

	_, err = conn.Do(
		"HMSET",
		"podcast:1",
		"title",
		"Teach Over Tea",
		"creator",
		"Brodie Robertson",
		"category",
		"technology",
		"membership_fee",
		9.99,
	)
	checkErr(err)
	title, err := redis.String(conn.Do("HGET", "podcast:1", "title"))
	checkErr(err)
	fmt.Println("Podcast Title: ", title)

	fee, err := redis.Float64(conn.Do("HGET", "podcast:1", "membership_fee"))
	checkErr(err)
	fmt.Println("Podcast membership fee: ", fee)

	values, err := redis.StringMap(conn.Do("HGETALL", "podcast:1"))
	checkErr(err)
	for k, v := range values {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", v)
	}

	reply, err := redis.Values(conn.Do("HGETALL", "podcast:1"))
	checkErr(err)
	var podcast Podcast

	err = redis.ScanStruct(reply, &podcast)
	checkErr(err)
	fmt.Printf("Podcast: %+v\n", podcast)

}
