package levenshtein

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {

	Convey("sort should return shorter word first", t, func() {

		Convey("it should work with words with ascii on them", func() {

			shorter, longer := sort("longerword", "short")

			So(longer, ShouldEqual, "longerword")
			So(shorter, ShouldEqual, "short")

		})

		Convey("it should work with words which have unicode in them", func() {

			shorter, longer := sort("أسد الصحراء", "word")

			So(longer, ShouldEqual, "أسد الصحراء")
			So(shorter, ShouldEqual, "word")

		})
	})
}

func TestMinimum(t *testing.T) {
	Convey("minimum should return smaller number", t, func() {
		t1 := minimum(3, 100)
		So(t1, ShouldEqual, 3)

		t2 := minimum(100, 3)
		So(t2, ShouldEqual, 3)

		t3 := minimum(-3, 3)
		So(t3, ShouldEqual, -3)

		t4 := minimum(0, 0)
		So(t4, ShouldEqual, 0)

		t5 := minimum(-100, -3)
		So(t5, ShouldEqual, -100)
	})
}

func TestRowVector(t *testing.T) {
	Convey("rowVector should return a vector with 1 row and len(shorterstring)+1 column", t, func() {
		s1 := []rune("longer")
		s2 := []rune("short")      //len(s2) = 5
		x := rowVector(s1, s2)     // x = [0,1,2,3,4,5]
		So(len(x), ShouldEqual, 6) //5+1 =6
		So(x[0], ShouldEqual, 0)
		So(x[5], ShouldEqual, 5)
	})
}

func TestDistance(t *testing.T) {
	Convey("distance should return the levenshtein distance", t, func() {
		So(Distance("shroud", "sound"), ShouldEqual, 3)
		So(Distance("lit", "lid"), ShouldEqual, 1)
		So(Distance("LID", "lid"), ShouldEqual, 0)
		So(Distance("", "lid"), ShouldEqual, 3)
		So(Distance("eating salad", "eating pizza"), ShouldEqual, 5)
		So(Distance("maka n", "makan"), ShouldEqual, 1)
		So(Distance("", ""), ShouldEqual, 0)
		So(Distance("الصحراء", "route"), ShouldEqual, 7)
		So(Distance("الصحراء", "الصحرا"), ShouldEqual, 1)
	})
}
