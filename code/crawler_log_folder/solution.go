package crawler_log_folder

func MinOperations(logs []string) int {
	folderDepth := 0
	for _, log := range logs {
		switch log {
		case "../":
			if folderDepth > 0 {
				folderDepth--
			}
		case "./":
			// do nothing
		default:
			folderDepth++
		}
	}
	return folderDepth
}
