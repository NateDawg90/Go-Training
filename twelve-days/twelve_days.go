package twelve

import "fmt"

const testVersion = 1

// Song prints out song
func Song() string {
	result := ""
	bundle := ""
	phrases := []string{
		"and a Partridge in a Pear Tree.",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}
	days := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}

	for i := 0; i < len(phrases); i++ {
		snippet1 := fmt.Sprintf("On the %s day of Christmas my true love gave to me, ", days[i])
		if i == 0 {
			bundle = "a Partridge in a Pear Tree."
		} else if i == 1 {
			bundle = phrases[i] + ", " + phrases[0]
		} else {
			bundle = phrases[i] + ", " + bundle
		}

		result = result + snippet1 + bundle + "\n"

	}
	return result
}

func Verse(i int) string {
	verses := []string{
		"On the first day of Christmas my true love gave to me, a Partridge in a Pear Tree.",

		"On the second day of Christmas my true love gave to me, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the third day of Christmas my true love gave to me, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the fourth day of Christmas my true love gave to me, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the fifth day of Christmas my true love gave to me, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the sixth day of Christmas my true love gave to me, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the seventh day of Christmas my true love gave to me, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the eighth day of Christmas my true love gave to me, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the ninth day of Christmas my true love gave to me, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the tenth day of Christmas my true love gave to me, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the eleventh day of Christmas my true love gave to me, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",

		"On the twelfth day of Christmas my true love gave to me, twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	}
	return verses[i-1]
}
