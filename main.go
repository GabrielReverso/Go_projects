package main

import (
	tree "module/DataStructure/Tree"
)

func main() {
	arvore := tree.Tree_create()

	/* tree.Tree_add(arvore, 4)
	tree.Tree_add(arvore, 2)
	tree.Tree_add(arvore, 1)
	tree.Tree_add(arvore, 3)
	tree.Tree_add(arvore, 5)
	tree.Tree_add(arvore, 5)
	tree.Tree_add(arvore, 6) */

	tree.Tree_add(arvore, 12)
	tree.Tree_add(arvore, 10)
	tree.Tree_add(arvore, 8)
	tree.Tree_add(arvore, 18)
	tree.Tree_add(arvore, 42)
	tree.Tree_add(arvore, 12)
	tree.Tree_add(arvore, 4)
	tree.Tree_add(arvore, 66)

	tree.Tree_pre_order(arvore)
	tree.Tree_order(arvore)
	tree.Tree_post_order(arvore)

}
