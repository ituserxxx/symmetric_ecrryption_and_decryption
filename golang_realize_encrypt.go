package encrypt

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// Position for padding string
const (
	PosLeft uint8 = iota
	PosRight
)

var (
	errConvertFail = errors.New("convert data type is failure")
	encryptSecret  = "0123456789ABCDEFGHJKMNPQRSTVWXYZabcdefghjkmnpqrstvwxyz"
)

// GetOffsetNumber 偏移数
func GetOffsetNumber(num int64) int64 {
	// 最新的抽奖id
	var theNewLotteryID int64 = 1041100
	var randNumbers [10]int64
	var oldRandNumber int64 = 123456789

	randNumbers[0] = 4874584950
	randNumbers[1] = 8978964530
	randNumbers[2] = 1581677540
	randNumbers[3] = 3543542440
	randNumbers[4] = 2534148840
	randNumbers[5] = 5545842750
	randNumbers[6] = 9454738480
	randNumbers[7] = 7354159310
	randNumbers[8] = 6191582590
	randNumbers[9] = 6458327540

	numStr := gconv.String(num)
	numLength := len(numStr)

	lastNum := gconv.Int(gstr.SubStr(numStr, numLength-1, 1))

	randNumber := randNumbers[lastNum]

	//这里只做解密处理，加密的话全部走新的加密算法
	//如果抽奖id小于1040766（代码更新前的最新抽奖id），则走之前的解密算法（123456789为以前加密固定数字）
	if num > oldRandNumber && num < theNewLotteryID+oldRandNumber {
		randNumber = oldRandNumber
	} else if num > 1100000+oldRandNumber && num < 38900000+oldRandNumber {
		//如果id大于1100000且小于38900000，该区间为用户uid，且排除了和抽奖lid重叠的uid，则走之前的解密算法
		randNumber = oldRandNumber
	}

	return randNumber
}

// NumToString 加密数字到字符串
func NumToString(num int64) string {
	num = num + GetOffsetNumber(num)
	destStr := PadLeft(gconv.String(num), "0", 10)
	destStrArray := ToArray(destStr, "")
	num1 := gconv.Int(gstr.TrimLeft(destStrArray[0]+destStrArray[2]+destStrArray[6]+destStrArray[9], "0"))
	num2 := gconv.Int(gstr.TrimLeft(destStrArray[1]+destStrArray[3]+destStrArray[4]+destStrArray[5]+destStrArray[7]+destStrArray[8], "0"))
	str1 := numToString(gconv.Int64(num1))
	str1 = gstr.Reverse(str1)
	str2 := numToString(gconv.Int64(num2))
	str2 = gstr.Reverse(str2)
	destStr = PadRight(str1, "U", 3) + PadRight(str2, "L", 4)
	return destStr
}

// numToString 数字转字符串
func numToString(num int64) string {
	baseStr := encryptSecret
	destStr := ""
	baseStrArray := ToArray(baseStr, "")
	for {
		if num == 0 {
			break
		}
		tempNum := num % 32
		destStr = destStr + baseStrArray[tempNum]
		num = num / 32
	}
	return destStr
}

// stringToNum 字符串转数字
func stringToNum(str string) float64 {
	baseStr := encryptSecret
	var num float64 = 0
	for i := 0; i < gstr.RuneLen(str); i++ {
		tempStr := gstr.SubStr(str, i, 1)
		index := gstr.PosIRune(baseStr, tempStr, 0)
		if index > 0 {
			num = num + (gconv.Float64(index) * math.Pow(32, gconv.Float64(gstr.RuneLen(str)-i-1)))
		}
	}
	return num
}

// StringToNum 解密字符串到数字
func StringToNum(str string) int64 {
	str1 := gstr.Trim(gstr.SubStr(str, 0, 3), "U")
	str2 := gstr.Trim(gstr.SubStr(str, 3, 4), "L")
	num1 := stringToNum(str1)
	num2 := stringToNum(str2)
	str1 = PadLeft(gconv.String(num1), "0", 4)
	str2 = PadLeft(gconv.String(num2), "0", 6)
	str1Array := ToArray(str1, "")
	str2Array := ToArray(str2, "")
	num := gconv.Int64(gstr.TrimLeft(str1Array[0]+str2Array[0]+str1Array[1]+
		str2Array[1]+str2Array[2]+str2Array[3]+str1Array[2]+
		str2Array[4]+str2Array[5]+str1Array[3], "0"))
	num -= GetOffsetNumber(num)
	return num
}

// Split string to slice. will clear empty string node.
func Split(s, sep string) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	for _, val := range strings.Split(s, sep) {
		if val = strings.TrimSpace(val); val != "" {
			ss = append(ss, val)
		}
	}
	return
}

// Trim string
func Trim(s string, cutSet ...string) string {
	if len(cutSet) > 0 && cutSet[0] != "" {
		return strings.Trim(s, cutSet[0])
	}

	return strings.TrimSpace(s)
}

// TrimLeft char in the string.
func TrimLeft(s string, cutSet ...string) string {
	if len(cutSet) > 0 {
		return strings.TrimLeft(s, cutSet[0])
	}

	return strings.TrimLeft(s, " ")
}

// TrimRight char in the string.
func TrimRight(s string, cutSet ...string) string {
	if len(cutSet) > 0 {
		return strings.TrimRight(s, cutSet[0])
	}

	return strings.TrimRight(s, " ")
}

// Substr for a string.
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	strLen := len(runes)

	// pos is to large
	if pos >= strLen {
		return ""
	}

	l := pos + length
	if l > strLen {
		l = strLen
	}

	return string(runes[pos:l])
}

// Padding a string.
func Padding(s, pad string, length int, pos uint8) string {
	diff := len(s) - length
	if diff >= 0 { // do not need padding.
		return s
	}

	if pad == "" || pad == " " {
		mark := ""
		if pos == PosRight { // to right
			mark = "-"
		}

		// padding left: "%7s", padding right: "%-7s"
		tpl := fmt.Sprintf("%s%d", mark, length)
		return fmt.Sprintf(`%`+tpl+`s`, s)
	}

	if pos == PosRight { // to right
		return s + Repeat(pad, -diff)
	}

	return Repeat(pad, -diff) + s
}

// PadLeft a string.
func PadLeft(s, pad string, length int) string {
	return Padding(s, pad, length, PosLeft)
}

// PadRight a string.
func PadRight(s, pad string, length int) string {
	return Padding(s, pad, length, PosRight)
}

// Repeat repeat a string
func Repeat(s string, times int) string {
	if times < 2 {
		return s
	}

	var ss []string
	for i := 0; i < times; i++ {
		ss = append(ss, s)
	}

	return strings.Join(ss, "")
}

// RepeatRune repeat a rune char.
func RepeatRune(char rune, times int) (chars []rune) {
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return
}

// ToInts alias of the ToIntSlice()
func ToInts(s string, sep ...string) ([]int, error) {
	return ToIntSlice(s, sep...)
}

// ToInt convert string to int
func ToInt(in interface{}) (iVal int, err error) {
	switch tVal := in.(type) {
	case int:
		iVal = tVal
	case int8:
		iVal = int(tVal)
	case int16:
		iVal = int(tVal)
	case int32:
		iVal = int(tVal)
	case int64:
		iVal = int(tVal)
	case uint:
		iVal = int(tVal)
	case uint8:
		iVal = int(tVal)
	case uint16:
		iVal = int(tVal)
	case uint32:
		iVal = int(tVal)
	case uint64:
		iVal = int(tVal)
	case float32:
		iVal = int(tVal)
	case float64:
		iVal = int(tVal)
	case string:
		iVal, err = strconv.Atoi(strings.TrimSpace(tVal))
	case nil:
		iVal = 0
	default:
		err = errConvertFail
	}
	return
}

// ToIntSlice split string to slice and convert item to int.
func ToIntSlice(s string, sep ...string) (ints []int, err error) {
	ss := ToSlice(s, sep...)
	for _, item := range ss {
		iVal, err := ToInt(item)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}

// ToArray alias of the ToSlice()
func ToArray(s string, sep ...string) []string {
	return ToSlice(s, sep...)
}

// ToSlice split string to array.
func ToSlice(s string, sep ...string) []string {
	if len(sep) > 0 {
		return Split(s, sep[0])
	}

	return Split(s, ",")
}
