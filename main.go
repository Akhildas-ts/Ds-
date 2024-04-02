package main

// // Stack
// type Stack struct {
// 	items []int
// }

// // push element into the Stack
// func (s *Stack) Push(i int) {

// 	s.items = append(s.items, i)

// }

// //Pop will remove the value at the end and
// // return the removed value

//		func (s *Stack) Pop() int {
//			l := len(s.items) - 1
//			removedElement := s.items[l]
//			s.items = s.items[:l]
//			return removedElement
//	}
func main() {

	// 1.STACK

	// mystack := Stack{}
	// fmt.Println(mystack)
	// mystack.Push(100)
	// mystack.Push(200)
	// mystack.Push(300)
	// fmt.Println(mystack)
	// mystack.Pop()
	// fmt.Println(mystack)

	//2.QUEUE AND DEQUEUE

	// q := Queue.Queue{}

	// q.Queue(2)
	// q.Queue(20)
	// q.Queue(40)
	// fmt.Println(q)

	// 3.LINKED LIST

	// linkedList := &Linklist.Linked{}

	// node1 := &Linklist.Node{Data: 10}
	// node2 := &Linklist.Node{Data: 20}
	// node3 := &Linklist.Node{Data: 30}
	// node4 := &Linklist.Node{Data: 50}
	// node5 := &Linklist.Node{Data: 50}
	// node6 := &Linklist.Node{Data: 60}
	// node7 := &Linklist.Node{Data: 70}
	// linkedList.Prepend(node1)
	// linkedList.Prepend(node2)
	// linkedList.Prepend(node3)
	// linkedList.Prepend(node4)

	// linkedList.Prepend(node5)
	// linkedList.Prepend(node6)
	// linkedList.RemoveNElement(2)
	// linkedList.PrinDataOnNode()
	// linkedList.Prepend(node7)

	// Linklist.DisplayList(linkedList.Head)

	// linkedList.PrinDataOnNode()

	// linkedList.DeleteLinkedListData(70)
	// linkedList.PrinDataOnNode()

	//4.BINARY SERACH

	// tree := &binaryserach.Node{Key: 39}

	// tree.Insert(400)
	// tree.Insert(40)
	// tree.Insert(80)
	// tree.Insert(4)
	// tree.Insert(30)
	// tree.Insert(90)

	// fmt.Println(tree.Serach(39))

	//5.STRING
	// text := "aaaawawaa"
	// substring := "aa"

	// fmt.Println(string_test.CountSubstringOccurrences(text, substring))

	//6.CONVERT ARRAY INTO LINKED LIST

	// values := []int{3, 4, 6, 7, 2, 7}
	// list := Linklist.ConvertIntoArray(values)
	// list.PrintListInArray()

	// 7.INSERT AFTER THE NODE

	// linkedList := &Linklist.Linked{}
	// // node1 := &Linklist.Node{Data: 10}
	// // node2 := &Linklist.Node{Data: 20}
	// // node3 := &Linklist.Node{Data: 30}
	// // node4 := &Linklist.Node{Data: 40}

	// // linkedList.Prepend(node1)
	// // linkedList.Prepend(node2)
	// // linkedList.Prepend(node3)
	// // linkedList.Prepend(node4)

	// // linkedList.InsertBefor(10,220)

	// // linkedList.PrintReverseData()

	// //8. REMOVE DUPILCATE ELEMENT'S

	// head := &Linklist.Node{Data: 1}
	// head.Next = &Linklist.Node{Data: 9}
	// head.Next.Next = &Linklist.Node{Data: 2}
	// head.Next.Next.Next = &Linklist.Node{Data: 3}
	// head.Next.Next.Next.Next = &Linklist.Node{Data: 3}

	// // Remove duplicates
	//  linkedList.RemoveDuplicate()

	// Linklist.DisplayList(linkedList.Head)

	// // 9. DOUBULE LINKEDLIST CREATED

	// new := double.D_LinkedList{}

	// new.D_Append(10)
	// new.PrintDouble()

	// new := double.doubleLink

	// dobuleLinked.D_Append(2)
	// dobuleLinked.D_Append(35)
	// dobuleLinked.D_Append(27)
	// dobuleLinked.PrintDouble()

	//10 . remove last nth element in singel linked list

	// 	linkedList.RemoveNTHelement(linkedList.Head,4)
	//    linkedList.PrintLinkedList(linkedList.Head)

	//A1. HASH INSERT
	// hashtable := hashs.Init()
	// // fmt.Println(hashtable)

	// inHash := hashs.Bucket{}
	// k := "RANDY"
	// inHash.Insert(k)
	// fmt.Println(inHash)
	// inHash.DeleteHash(k)
	// fmt.Println(inHash.SerachHash("RANDY"))

	// A3 bubble sort algoritham

	// arr := []int{2, 4, 6, 2, 6, 7, 29}
	// var store []int
	// store = bubblesort.BubbleSort(arr)
	// fmt.Println(store)

	// A4 berfor n Element (bubble sort)

	// bubblesort.CheckItDivisible(39)

	//A5 insertion sort
	// arr := []int{3, 6, 2, 4, 5, 10, 9}

	// insertionsort.InsertionSort(arr)

	// fmt.Println(arr)

	// A6 . Quick sort

	// here is the algoritham , other are problems i solved with quick sort ..

	// arr := []int{2, 3, 7, 2, 4, 7, 8, 4}
	// quicksort.QuickSort(arr,0,len(arr)-1)
	// fmt.Println(arr)

	// quicksort.LargestElementK(arr,3)

	// student := []quicksort.Student{

	// 	{Name: "akhil ",Age: 18,Grade: 10.00},
	// 	{Name: "arun",Age: 20,Grade: 5.00},
	// 	{Name: "anil",Age: 15,Grade: 20.00},
	// 	{Name: "remesh",Age: 12,Grade: 30.00},
	// 	{Name: "neran",Age: 10,Grade: 110.0},
	// }

	// quicksort.StudentSort(student)

	// A7 problem solving ..
	//  with hash table basic ..

	// h1 := problem.Init1()
	// fmt.Println(h1)

	// nums := []int{2, 36,  7,  3,3}
	// problem.ContainsDuplicate(nums)

}
