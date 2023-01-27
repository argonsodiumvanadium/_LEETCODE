package main

import "fmt"

type Node struct {
    Child *Node
    Val   int
}

func Construct_list (arr []int) (head *Node) {
    head = &Node {nil, arr[0]}

    itr := head
    for _, val := range arr[1:] {
        (*itr).Child = &Node{nil, val}
        itr = (*itr).Child
    }
    (*itr).Child = head

    return
}

func (self *Node) Move (iteration_no int) (*Node, *Node) {
    chain := self.Child
    itr := self.Child.Child.Child.Child
    self.Child.Child.Child.Child = nil
    self.Child = itr

    for val:=self.GetNextCup(self.Val);;val=self.GetNextCup(val) {
        if chain.Has(val) {
            continue
        }
        itr = itr.Get(val)
        break
    }

    chain.Child.Child.Child = itr.Child
    itr.Child = chain

    soln_head := self.IterateTo(iteration_no%9)
    return self.Child, soln_head
}

func (self *Node) GetNextCup (itr int) (int) {
    if itr < 2 {
        return 9
    } else {
        return itr - 1
    }
}

func (self *Node) Has (val int) (bool) {
    itr := self

    for itr != nil {
        if itr.Val == val {
            return true
        }

        itr=itr.Child
    }

    return false
}

func (self *Node) Get (val int) (itr *Node) {
    itr = self

    for i:=0;itr != nil && i < 10;i++ {
        if itr.Val == val {
            return
        }

        itr = itr.Child
    }

    return nil
}

func (self *Node) IterateTo (iteration_no int) (*Node) {
    itr := self
    for i:=0;i<iteration_no;i++ {
        itr = itr.Child
    }

    return itr
}

func (self *Node) Println () {
    itr := self

    for iterations:=0;itr != nil && iterations < 9; iterations ++ {
        fmt.Print(itr.Val);
        itr = itr.Child
    }
    fmt.Println("");
}

func main () {
    head := Construct_list([]int{7,3,9,8,6,2,5,4,1})

    itr := head
    for i:=0;i<100;i++ {
        itr, head = itr.Move(i)
    }
    fmt.Printf("Part 1: ")
    head.Get(1).Println()
}

