package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
func FindClosetToZero(nums []int) int{
	closestNumber := math.MaxInt64 

	for _, num := range nums {
		if abs(num) < abs(closestNumber) || (abs(num) == abs(closestNumber) && num > 0 && closestNumber < 0) {
			closestNumber = num
		}
	}

	return closestNumber
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

// 56789 -> 68957
func MaxRot(n int64) int64 {
    var highValue int64
	numStr := strconv.FormatInt(n, 10)
	for i := 0; i < len(numStr)-1; i++{
		if i != len(numStr) - 1{
			firstChunk := numStr[:i]
			secondChunk := numStr[i + 1:]
			numStr = firstChunk + secondChunk + string(numStr[i:i+1])
			fmt.Println(numStr)
		}
		newNum, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			fmt.Println("not a number")
		}
		if highValue < newNum{
			highValue = newNum
		}
	}
	return highValue
}

func FirstNonRepeating(str string) string {
	if len(str) <= 1 {
		return str
	}
	repeatedStr := map[string]string{}
	for i := 0; i < len(str); i++ {
		char := str[i]
		answer := string(char)
		for j := i + 1; j < len(str); j++{
			if repeatedStr[string(char)] == string(char){
				answer = ""
			}
			if strings.EqualFold(string(char), string(str[j])) {
				answer = ""
				repeatedStr[string(char)] = string(char)
			}
		}
		if repeatedStr[answer] != answer{
			return answer
		}
	}
	return ""
}

// ".*?" ==> To(Equal(false)) ✔️
	// "a" ==> To(Equal(true)) ✔️
	// "Mazinkaiser" ==> To(Equal(true)) ✔️
	// "hello world_" ==> To(Equal(false)) ✔️
	// "     " ==> To(Equal(false)) ✔️
	// "PassW0rd" ==> To(Equal(true)) ✔️
	// "" ==> To(Equal(false)) ✔️
	// "\n\t\n" ==> To(Equal(false)) ✔️
	// "ciao\n$$_" ==> To(Equal(false)) ✔️
	// "__ * __" ==> To(Equal(false)) ✔️
	// "&)))(((" ==> To(Equal(false)) ✔️
	// "43534h56jmTHHF3k" ==> To(Equal(true)) ✔️

func Alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, val := range str {
		switch {
		case val < 48 || val > 57 && val < 65 || val > 90 && val < 97:
			return false
		default:
			continue
		}
	}

	return true
}

func WordCount(s string) map[string]int {
	countMap := map[string]int{}
	stringArr := strings.Split(s, " ")
	for _, word := range stringArr {
	w, ok := countMap[word]
	 if ok {
	 	countMap[word] = w + 1
		continue
	 }
	 countMap[word] = 1
	}
	return countMap
}

// func CreatePhoneNumber(arr []uint)string{
// 	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9])
	
// }

// func CreatePhoneNumber(numbers []uint) string {
//   var test string = strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
//   return fmt.Sprintf("(%s) %s-%s", test[0:3], test[3:6], test[6:10])  
// }

func CreatePhoneNumber(numbers []uint) string {
  tmp := make([]interface{}, len(numbers))
  for i, val := range numbers {
    tmp[i] = val
  }
  return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", tmp...)
}


// Expect(ToNato("If you can read")).To(Equal("India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"))
//      Expect(ToNato("Did not see that coming")).To(Equal("Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"))
//      Expect(ToNato("go for it!")).To(Equal("Golf Oscar Foxtrot Oscar Romeo India Tango !"))

var NATO = map[string]string{
	"A":"Alfa","B":"Bravo","C":"Charlie","D":"Delta","E":"Echo","F":"Foxtrot","G":"Golf","H":"Hotel","I":"India","J":"Juliett","K":"Kilo","L":"Lima","M":"Mike","N":"November","O":"Oscar","P":"Papa","Q":"Quebec","R":"Romeo","S":"Sierra","T":"Tango","U":"Uniform","V":"Victor","W":"Whiskey","X":"Xray","Y":"Yankee","Z":"Zulu",
}

func ToNato(words string) string {
	natoTranslatedText := ""
	for _, char := range words{
		if char == ' '{
			continue
		}
		if char >= 33 && char <= 47 || char >=58 && char <= 64 {
			natoTranslatedText += string(char)
		}
		natoTranslatedText += NATO[strings.ToUpper(string(char))] + " "
	}
	return strings.Trim(natoTranslatedText, " ")
}

func Int32ToIp(n uint32) string {
	num := uint32(n)
	iparray := []string{}
	binaryStr := fmt.Sprintf("%032b", num)

	for i:=1; i<=len(binaryStr); i++ {
		if i % 8 == 0 && i != 0 {
			binaryChunk := binaryStr[i-8: i]
			intChunk, _ := strconv.ParseUint(binaryChunk, 2, 8)
			iparray = append(iparray, strconv.Itoa(int(intChunk)))
		}
	}
	ip := strings.Join(iparray, ".")
	return ip
}


func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}

	result := []int{}

	for lng > 0 && wdth > 0 {
		if lng < wdth {
			result = append(result, lng)
			wdth -= lng
		} else {
			result = append(result, wdth)
			lng -= wdth
		}
	}

	return result
}

func ProperFractions(n int) int {
	if n <= 1 {
		return 0
	}

	result := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}


func JosephusSurvivor(n, k int) int {
	if n == 1 {
		return 1 
	}

	position := 0 

	for i := 2; i <= n; i++ {
		position = (position + k) % i
	}

	return position + 1
}

func Solution(nums []int) string { // convert num array to string shorten by adding - for consecutive numbers over 3
	if len(nums) == 0 {
		return ""
	}

	var result []string
	start := nums[0]
	end := start

	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
				end = nums[i]
		} else {
			if end > start+1 {
					result = append(result, fmt.Sprintf("%d-%d", start, end))
			} else if end > start {
					result = append(result, fmt.Sprintf("%d,%d", start, end))
			} else {
					result = append(result, strconv.Itoa(start))
			}
			start = nums[i]
			end = start
		}
	}

	if end > start+1 {
			result = append(result, fmt.Sprintf("%d-%d", start, end))
	} else if end > start {
			result = append(result, fmt.Sprintf("%d,%d", start, end))
	} else {
			result = append(result, strconv.Itoa(start))
	}

	return strings.Join(result, ",")
}


// ['a','b','c','d','f'] -> 'e'
func FindMissingLetter(chars []rune) rune {
	var missing rune
	for i,r := range chars {
		if r + 1 != chars[i+1]{
			missing = r + 1
			break
		}
	}
	return missing
}


func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	timeUnits := []struct {
		unit   string
		secondsInUnit int64
	}{
		{"year", 365 * 24 * 60 * 60},
		{"day", 24 * 60 * 60},
		{"hour", 60 * 60},
		{"minute", 60},
		{"second", 1},
	}

	var components []string
	for _, timeUnit := range timeUnits {
		if seconds >= timeUnit.secondsInUnit {
			unitCount := seconds / timeUnit.secondsInUnit
			seconds %= timeUnit.secondsInUnit
			if unitCount > 1 {
				timeUnit.unit += "s"
			}
			components = append(components, fmt.Sprintf("%d %s", unitCount, timeUnit.unit))
		}
	}

	if len(components) == 1 {
		return components[0]
	}
	return strings.Join(components[:len(components)-1], ", ") + " and " + components[len(components)-1]
}