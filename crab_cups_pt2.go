package main

import "fmt"

const (
    LIST_LEN = 1000000
)

var (
    lookup [LIST_LEN+1]*Node
    barred []int
)


type Node struct {
    Child *Node
    Val   int
}

func Construct_list (arr []int) (head *Node) {
    head = &Node {nil, arr[0]}
    lookup[arr[0]]=head

    itr := head
    for _, val := range arr[1:] {
        (*itr).Child = &Node{nil, val}
        lookup[val] = itr.Child

        itr = (*itr).Child
    }
    for i:=10;i<=LIST_LEN;i++ {
        itr.Child = &Node{nil, i}
        lookup[i] = itr
        itr=itr.Child
    }
    (*itr).Child = head
    return
}

func (self *Node) Move (iteration_no int) (*Node, *Node) {
    chain := self.Child
    itr := self.Child.Child.Child.Child
    self.Child.Child.Child.Child = nil
    self.Child = itr


    barred = []int{chain.Val,chain.Child.Val,chain.Child.Child.Val}

    for val:=self.GetNextCup(self.Val);;val=self.GetNextCup(val) {
        if Has(barred, val) {
            continue
        }
        itr = itr.Get(val)
        break
    }

    chain.Child.Child.Child = itr.Child
    itr.Child = chain


    return self.Child, self
}

func (self *Node) GetNextCup (itr int) (int) {
    if itr < 2 {
        return LIST_LEN
    } else {
        return itr - 1
    }
}

func Has (arr []int, val int) (bool) {
    for _, v := range arr {
        if v == val {
            return true
        }
    }

    return false
}

func (self *Node) Get (val int) (itr *Node) {
    itr = self
    return lookup[val]
}

func (self *Node) Println () {
    itr := self

    for iterations:=0;itr != nil && iterations < LIST_LEN+1; iterations ++ {
        fmt.Print(itr.Val);
        itr = itr.Child
    }
    fmt.Println("");
}

func main () {
    head := Construct_list([]int{7,3,9,8,6,2,5,4,1})

    itr := head
    for i:=0;i<10000000;i++ {
        fmt.Println("iteration:",i)
        itr, head = itr.Move(i)
    }
    itr=head.Get(1)
    fmt.Println((itr.Child.Val),(itr.Child.Child.Val))
    fmt.Println((itr.Child.Val)*(itr.Child.Child.Val))
}

