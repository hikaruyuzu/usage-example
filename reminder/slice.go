package main

import "fmt"

func main() {
	// slice
	var fruits = []string {"Pisang", "Anggur", "Mangga", "Apel","Semangka"}
	fmt.Println(fruits)

	//var test1 = [3]string {"test1", "test2"."test3"} // array
	//var test2 = [...]string {"test1", "test2"."test3"} // array

	fmt.Println(fruits[0:2])
	fmt.Println(fruits[2:])
	fmt.Println(fruits[0:0])
	fmt.Println(fruits[:])

	var buah = [] string {"Jeruk", "Mangga", "Apel", "Pear", "Jambu"}
	aBuah := buah[0:3]
	bBuah := buah[1:4]

	aaBuah := aBuah[0:2]
	bbBuah := bBuah[0:1]

	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(bbBuah)
	fmt.Println(buah)

	bbBuah[0] = "Wakaranai Desu"
	fmt.Println(aaBuah)
	fmt.Println(buah)
	fmt.Println(bBuah)

	var month = [] string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(month)
	fmt.Println(len(month))

	birtday := month[0:9]
	fmt.Println(cap(birtday))

	birtday2 := month[8:]
	fmt.Println(len(birtday2))
	fmt.Println(cap(birtday2))

	birtday3 := month[8:9]
	fmt.Println(len(birtday3))
	fmt.Println(cap(birtday3))

	birtday4 := month[10:]

	birtday5 := birtday4
	birtday4[0] = "Kaguya sama"
	fmt.Println(birtday5)
	fmt.Println(birtday4)
	fmt.Println(month)

	birtday6 := birtday4
	birtday6[1] = "Tsukasa Chan"


	fmt.Println(birtday6)
	fmt.Println(month)

	test := month[10:]
	testAppend := append(test, "shino aki", "Yukinoshita")
	fmt.Println(testAppend)

	fmt.Println(test)

	test2 := month[8:10]
	testAppend2 := append(test2, "Gura")
	fmt.Println(testAppend2)
	fmt.Println(test2)
	fmt.Println(month)

	appType := make([]string, 3, 5) // 5 is capacity
	appType[0] = "ReadManga"
	appType[1] = "Streaming Anime"
	appType[2] = "BuyManga"

	rsc := make([]string, 4)
	var src = []string {"Clannad", "Kaguya Sama", "Plastic Memories", "Higehiro", "Shigatsu wakimi no uso"}

	resCopy := copy(rsc, src)
	fmt.Println(resCopy)
	fmt.Println(rsc)

	var name = []string {"Kaguya sama", "Tsukasa chan", "Yukinoshita"}

	waifu := make([]string, len(name))
	getName := copy(waifu, name)

	fmt.Printf("Lenght of getName = %d\n", getName)
	fmt.Println(waifu)

	var animals = []string {"Buaya", "Singa", "Anjing", "Kucing", "Nekomimi"}
	// last : : :__ is a lenght of slice
	animal := animals[3:4:5]
	fmt.Println(cap(animal))
	fmt.Println(len(animal))





}
