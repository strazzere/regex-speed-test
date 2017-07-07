package main

import (
	"fmt"
	"regexp"
	"time"
)

var fails = []string{
	"test",
	"test1",
	"localhaus",
	//	"192.168.111.1213", Doesn't fail - though in reality it isn't even an ip address
	//	"192.168.1111", Same as above
	"::2",
	"fe80::2",
	"derp.rednaga.itt",
	"nah-this-aint-right.rednaga.orgs",
	"gdfgfdgeq4r23in409r0dni30h4i4n98hf985jd895j98j9e8rnggsidjrniucwnriuwcj4ciru3",
}

var successes = []string{
	"localhost",
	"wiki.rednaga.org",
	"wiki.rednaga.it",
	"fe80::1",
	"::1",
	"holymolygobilygookandalltheotherstuffblahblahblahblah.rednaga.it",
	"oiejfowiejfoiwejfoiewjoicmoimeocimsdocmdskcmsdlkcmsdlkmclkwjdoiqwjdowq.rednaga.org",
}

func testFail(regex *regexp.Regexp) {
	testSet(regex, fails, false)
}

func testSuccess(regex *regexp.Regexp) {
	testSet(regex, successes, true)
}

func testSet(regex *regexp.Regexp, set []string, expectation bool) {
	for i := 0; i < len(set); i++ {
		if expectation != regex.MatchString(set[i]) {
			panic("Error in test set for " + set[i])
		}
	}
}

func testRegex(regex string) []int64 {
	timing := []int64{-1, -1, -1}
	start := time.Now()
	regexp := regexp.MustCompile(regex)
	elapsed := time.Since(start)
	//log.Printf("Regex compilation took %s", elapsed)
	timing[0] = elapsed.Nanoseconds()

	// test failure speed
	start = time.Now()
	testFail(regexp)
	elapsed = time.Since(start)
	//log.Printf("Fail run took %s", elapsed)
	timing[1] = elapsed.Nanoseconds()

	// test success speed
	start = time.Now()
	testSuccess(regexp)
	elapsed = time.Since(start)
	//log.Printf("Success run took %s", elapsed)
	timing[2] = elapsed.Nanoseconds()

	return timing
}

func testSmart() []int64 {
	return testRegex(`^(?:localhost` +
		`|(?:127|192\.168|182\.(?:1[6-9]|2[0-9]|3[0-1])\.|10` +
		`|188\.27|193\.245)\.[0-9\.]+` +
		`|::1|fe80::1` +
		`|(?:[\w.]+\.)?(?:rednaga\.it|rednaga\.org)` +
		`)$`)
}

func testDumb() []int64 {
	return testRegex(`(localhost|` +
		`wiki.rednaga.org|` +
		`wiki.rednaga.it|` +
		`fe80::1|` +
		`::1|` +
		`holymolygobilygookandalltheotherstuffblahblahblahblah.rednaga.it|` +
		`oiejfowiejfoiwejfoiewjoicmoimeocimsdocmdskcmsdlkcmsdlkmclkwjdoiqwjdowq.rednaga.org)`)
}

func main() {
	runs := int64(100000)
	fmt.Println("Running test", runs, "times.")

	averageComp := int64(0)
	averageFail := int64(0)
	averageSucc := int64(0)
	averageTotal := int64(0)
	for i := 0; i < int(runs); i++ {
		start := time.Now()
		timing := testSmart()
		elapsed := time.Since(start)
		//log.Printf("Smart test took %s", elapsed)
		averageComp += timing[0]
		averageFail += timing[1]
		averageSucc += timing[2]
		averageTotal += elapsed.Nanoseconds()
	}
	fmt.Println("Smart Run -- compilation, fail, success, total")
	fmt.Println(averageComp/runs, averageFail/runs, averageSucc/runs,
		averageTotal/runs)

	averageComp = int64(0)
	averageFail = int64(0)
	averageSucc = int64(0)
	averageTotal = int64(0)

	for i := 0; i < int(runs); i++ {
		start := time.Now()
		timing := testDumb()
		elapsed := time.Since(start)
		//log.Printf("Dumb test took %s", elapsed)
		averageComp += timing[0]
		averageFail += timing[1]
		averageSucc += timing[2]
		averageTotal += elapsed.Nanoseconds()
	}
	fmt.Println("Dumb Run -- compilation, fail, success, total")
	fmt.Println(averageComp/runs, averageFail/runs, averageSucc/runs,
		averageTotal/runs)
}
