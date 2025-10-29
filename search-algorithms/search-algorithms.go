package searchalgorithms

func BinarySearch_Int(key int, list []int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		//Calculate the middle
		middle := left + ((right - left) / 2)

		if list[middle] == key {
			return middle
		}

		if list[middle] > key {
			right = middle - 1

		} else if list[middle] < key {
			left = middle + 1
		}
	}

	return -1
}

func BinarySearch_String(key string, list []string) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		middle := left + ((right - left) / 2)

		if list[middle] == key {
			return middle
		}

		if list[middle] < key {
			left = middle + 1
		} else if list[middle] > key {
			right = middle - 1
		}
	}
	return -1
}
