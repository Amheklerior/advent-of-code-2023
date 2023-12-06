package day1

func getFirst(nums []string) string {
	return nums[0]
}

func getLast(nums []string) string {
	return nums[len(nums)-1]
}

func findDigits(str string) []string {
	numbers := regexp
		.MustCompile(`\d`)
		.FindAllString(str, -1)
	return numbers
}

func getPrefixUntilFirstDigit(str string) (string, bool) {
	prefix := regexp
		.MustCompile(`^\D+`)
		.FindString(line)
	return (prefix, len(prefix) > 0)
}

func getSuffixFromLastDigit(str string) (string, bool) {
	suffix := regexp
		.MustCompile(`\D+$`)
		.FindString(line)
	return (suffix, len(prefix) > 0)
}

func replaceWithDigits(str string) string {
	if preserveLast {
		replacer := strings.NewReplacer(
			"oneight", "8", 
			"twone", "1", 
			"threeight", "8", 
			"fiveight", "8", 
			"sevenine", "9", 
			"eightwo", "2", 
			"eighthree", "3", 
			"nineight", "8", 
			"one", "1", 
			"two", "2", 
			"three", "3", 
			"four", "4", 
			"five", "5", 
			"six", "6", 
			"seven", "7", 
			"eight", "8", 
			"nine", "9"
		)	
	} else {
		replacer := strings.NewReplacer(
			"one", "1", 
			"two", "2", 
			"three", "3", 
			"four", "4", 
			"five", "5", 
			"six", "6", 
			"seven", "7", 
			"eight", "8", 
			"nine", "9"
		)
	}
	return replacer.Replace(str)
}
