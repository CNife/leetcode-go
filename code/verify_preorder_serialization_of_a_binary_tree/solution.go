package verify_preorder_serialization_of_a_binary_tree

func IsValidSerialization(preorder string) bool {
	slots := 1
	for i := 0; i < len(preorder); {
		switch {
		case slots == 0:
			return false
		case preorder[i] == ',':
			i++
		case preorder[i] == '#':
			slots--
			i++
		default:
			for ; i < len(preorder) && preorder[i] != ','; i++ {
			}
			slots++
		}
	}
	return slots == 0
}
