//二叉树
package btree

import (
	"fmt"
	"github.com/caoxiaolin/go-data-structure/queue"
)

//树结构
type BinaryTree struct {
	Data  interface{}
	Lnode *BinaryTree
	Rnode *BinaryTree
}

var str []byte

//构造二叉树，前序构建
func Build(s []byte) *BinaryTree {
	var t BinaryTree
	if len(s) == 0 {
		return nil
	}
	c := string(s[0])
	str = s[1:]
	if c != "#" {
		t.Data = c
		t.Lnode = Build(str)
		t.Rnode = Build(str)
	}
	return &t
}

//先序遍历
func (bt *BinaryTree) Preorder() {
	if bt == nil {
		return
	}
	if bt.Data != nil {
		fmt.Print(bt.Data.(string))
	} else {
		//fmt.Print("#")
	}
	bt.Lnode.Preorder()
	bt.Rnode.Preorder()
}

//中序遍历
func (bt *BinaryTree) Inorder() {
	if bt == nil {
		return
	}
	bt.Lnode.Inorder()
	if bt.Data != nil {
		fmt.Print(bt.Data.(string))
	} else {
		//fmt.Print("#")
	}
	bt.Rnode.Inorder()
}

//后序遍历
func (bt *BinaryTree) Postorder() {
	if bt == nil {
		return
	}
	bt.Lnode.Postorder()
	bt.Rnode.Postorder()
	if bt.Data != nil {
		fmt.Print(bt.Data.(string))
	} else {
		//fmt.Print("#")
	}
}

//层序遍历
func (bt *BinaryTree) Levelorder() {
	if bt == nil {
		return
	}
	queue := queue.InitQueue()
	queue.Enter(bt)
	for queue.IsEmpty() == false {
		n, _ := queue.Out()
		node := n.(*BinaryTree)
		if node.Data != nil {
			fmt.Print(node.Data.(string))
		} else {
			//fmt.Print("#")
		}
		if node.Lnode != nil {
			queue.Enter(node.Lnode)
		}
		if node.Rnode != nil {
			queue.Enter(node.Rnode)
		}
	}
}

func Example() {
	s := []byte("ABDH#K###E##CFI###G#J##")
	fmt.Printf("前序构造二叉树：%s\n", string(s))
	t := Build(s)
	fmt.Print("前序遍历：")
	t.Preorder()
	fmt.Print("\n中序遍历：")
	t.Inorder()
	fmt.Print("\n后序遍历：")
	t.Postorder()
	fmt.Print("\n层序遍历：")
	t.Levelorder()
}
