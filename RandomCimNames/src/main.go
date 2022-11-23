package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	//	"golang.org/x/text/cases"
	//"golang.org/x/text/language"
)

func main() {
	//import word list
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	file, ferr := os.Open("wordList.txt")
	if ferr != nil {
		panic(ferr)
	}

	//extract words
	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		x := 0
		line := scanner.Text()
		if line != " " {
			line = line + "#"
			splitWords := strings.Split(line, " ")
			//	fmt.Println(line)
			for i := 1; i < 4; i++ {
				if len(splitWords) < 4 {
					for j := 0; j < len(splitWords)-1; j++ {
						words = append(words, splitWords[j])
					}
				} else {
					words = append(words, splitWords[i-1])
				}
			}
		}
		x++
	}

	//catagorise words
	var prefixes []string
	var suffixes []string
	var nouns []string
	var adjectives []string
	var verbs []string
	var adverbs []string

	for y := 1; y < len(words)-1; y++ {
		if words[y] == "##" {
			break
		}
		if words[y] == "prefix" {
			prefixes = append(prefixes, words[y-1])
		}

		if words[y] == "suffix" {
			suffixes = append(suffixes, words[y-1])
		}

		if words[y] == "n." || words[y] == "-n." || words[y] == "var." {
			nouns = append(nouns, words[y-1])
		}

		if words[y] == "adj." || words[y] == "-adj." {
			adjectives = append(adjectives, words[y-1])
		}

		if words[y] == "v." || words[y] == "-v." {
			verbs = append(verbs, words[y-1])
		}

		if words[y] == "adv." || words[y] == "-adv." {
			adverbs = append(adverbs, words[y-1])
		}
	}

	//generate a firstname
	rand.Seed(time.Now().UTC().UnixNano())
	selectGenerationMethod := rand.Intn(5-1) + 1
	var chosenFirstName string
	switch selectGenerationMethod {
	// adjective and suffix
	case 1:
		chosenAdj := adjectives[rand.Intn(len(adjectives)-1)+1]
		rand.Seed(time.Now().UTC().UnixNano())
		chosenSuff := suffixes[rand.Intn(len(suffixes)-1)+1]
		if len(chosenAdj) > 6 {
			chosenAdj = chosenAdj[0:4]
		}
		chosenSuff = chosenSuff[1 : len(chosenSuff)-1]
		chosenFirstName = chosenAdj + chosenSuff
		fmt.Println("adjective and suffix")
	//adjective with a wobble (start of random noun)
	case 2:
		chosenAdj := adjectives[rand.Intn(len(adjectives)-1)+1]
		wabble := nouns[rand.Intn(len(nouns)-1)+1][0:1]
		if len(chosenAdj) > 5 {
			wibble := chosenAdj[len(chosenAdj)-2 : len(chosenAdj)-1]
			wobble := chosenAdj[len(chosenAdj)-3 : len(chosenAdj)-2]

			chosenFirstName = chosenAdj[0:len(chosenAdj)-6] + wabble + "e" + wobble + wibble
		} else {
			chosenFirstName = chosenAdj + "e" + wabble
		}
		fmt.Println("adjective wibble")
	//adjective with an adjective
	case 3:
		chosenAdj1 := adjectives[rand.Intn(len(adjectives)-1)+1]
		rand.Seed(time.Now().UTC().UnixNano())
		chosenAdj2 := adjectives[rand.Intn(len(adjectives)-1)+1]
		if len(chosenAdj1) > 6 {
			chosenAdj1 = chosenAdj1[0:3]
		}
		if len(chosenAdj2) > 6 {
			chosenAdj2 = chosenAdj2[0:3]
		}
		chosenAdj1 = chosenAdj1[0 : len(chosenAdj1)-1]
		chosenAdj2 = chosenAdj2[0 : len(chosenAdj2)-1]
		chosenFirstName = chosenAdj1 + "a" + chosenAdj2
		fmt.Println("adjective adjective")
	//truncated verb
	//else add suffix to verb
	case 4:
		chosenVerb := verbs[rand.Intn(len(verbs)-1)+1]
		if len(chosenVerb) > 3 {
			chosenFirstName = chosenVerb[0 : len(chosenVerb)-2]
			chosenFirstName = chosenFirstName + "y"
			fmt.Println("truncated verb")
		} else {
			chosenSuff := suffixes[rand.Intn(len(suffixes)-1)+1]
			chosenSuff = chosenSuff[1 : len(chosenSuff)-1]
			chosenFirstName = chosenVerb + chosenSuff
			fmt.Println("suffixed verb")
		}
		if len(chosenFirstName) < 3 {
			chosenFirstName = chosenFirstName + "a"
		}
	}

	//generate a surname
	rand.Seed(time.Now().UTC().UnixNano())
	selectGenerationMethod = rand.Intn(5-1) + 1
	var chosenSurname string
	switch selectGenerationMethod {
	// noun
	case 1:
		chosenSurname = nouns[rand.Intn(len(nouns)-1)+1]
		fmt.Println("noun")
	//prefix noun
	case 2:
		chosenPrefix := prefixes[rand.Intn(len(prefixes)-1)+1]
		rand.Seed(time.Now().UTC().UnixNano())
		chosenNoun := nouns[rand.Intn(len(nouns)-1)+1]
		chosenPrefix = chosenPrefix[0 : len(chosenPrefix)-1]
		if len(chosenNoun) > 11 {
			chosenNoun = chosenNoun[0 : len(chosenNoun)-8]
		}
		chosenSurname = chosenPrefix + chosenNoun
		fmt.Println("prefix noun")
	//adverb + noun
	case 3:
		chosenAdv := adverbs[rand.Intn(len(adverbs)-1)+1]
		rand.Seed(time.Now().UTC().UnixNano())
		chosenNoun := nouns[rand.Intn(len(nouns)-1)+1]
		if len(chosenAdv) > 4 {
			chosenAdv = chosenAdv[0 : len(chosenAdv)-3]
		}
		if len(chosenAdv) > 7 {
			chosenAdv = chosenAdv[0 : len(chosenAdv)-5]
		}
		if len(chosenAdv) > 11 {
			chosenAdv = chosenAdv[0 : len(chosenAdv)-8]
		}
		if len(chosenNoun) > 7 {
			chosenNoun = chosenNoun[0 : len(chosenNoun)-5]
		}
		if len(chosenNoun) > 11 {
			chosenNoun = chosenNoun[0 : len(chosenNoun)-8]
		}

		chosenSurname = chosenAdv + chosenNoun
		fmt.Println("adverb noun")
	//firstname -son
	case 4:
		rand.Seed(time.Now().UTC().UnixNano())
		sgm := rand.Intn(5-1) + 1
		var cfn string
		switch sgm {
		// adjective and suffix
		case 1:
			chosenAdj := adjectives[rand.Intn(len(adjectives)-1)+1]
			rand.Seed(time.Now().UTC().UnixNano())
			chosenSuff := suffixes[rand.Intn(len(suffixes)-1)+1]
			if len(chosenAdj) > 6 {
				chosenAdj = chosenAdj[0:4]
			}
			if len(chosenSuff) > 2 {
				chosenSuff = chosenSuff[1 : len(chosenSuff)-2]
			}
			cfn = chosenAdj + chosenSuff
			fmt.Println("adjective and suffix")
		//adjective with a wobble
		case 2:
			chosenAdj := adjectives[rand.Intn(len(adjectives)-1)+1]
			wabble := nouns[rand.Intn(len(adjectives)-1)+1][0:1]
			if len(chosenAdj) > 3 {
				wibble := chosenAdj[len(chosenAdj)-2 : len(chosenAdj)-1]
				wobble := chosenAdj[len(chosenAdj)-3 : len(chosenAdj)-2]

				cfn = chosenAdj[0:len(chosenAdj)-4] + wabble + "e" + wobble + wibble
			} else {
				cfn = chosenAdj + "e" + wabble
			}
			fmt.Println("adjective wibble")
		//adjective with an adjective
		case 3:
			chosenAdj1 := adjectives[rand.Intn(len(adjectives)-1)+1]
			rand.Seed(time.Now().UTC().UnixNano())
			chosenAdj2 := adjectives[rand.Intn(len(adjectives)-1)+1]
			chosenAdj1 = chosenAdj1[len(chosenAdj1)-2 : len(chosenAdj1)-1]
			chosenAdj2 = chosenAdj2[len(chosenAdj2)-2 : len(chosenAdj2)-1]
			cfn = chosenAdj1 + "a" + chosenAdj2
			fmt.Println("adjective adjective")
		//truncated verb
		//else add suffix to verb
		case 4:
			rand.Seed(time.Now().UTC().UnixNano() + int64(len(chosenFirstName)))
			chosenVerb := verbs[rand.Intn(len(verbs)-1)+1]
			if len(chosenVerb) > 3 {
				cfn = chosenVerb[0 : len(chosenVerb)-2]
				fmt.Println("truncated verb")
			} else {
				chosenSuff := suffixes[rand.Intn(len(suffixes)-1)+1]
				chosenSuff = chosenSuff[1 : len(chosenSuff)-1]
				cfn = chosenVerb + chosenSuff
				fmt.Println("suffixed verb")
			}
		}
		if len(cfn) < 3 {
			cfn = cfn + "a"
		}
		fmt.Println("firstname -son")
		chosenSurname = cfn + "son"
	}

	//cleanup
	chosenFirstName = strings.ToLower(chosenFirstName)
	chosenSurname = strings.ToLower(chosenSurname)

	chosenFirstName = strings.ReplaceAll(chosenFirstName, "-", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "'", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, ".", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "1", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "2", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "3", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "4", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "5", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "6", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "7", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "8", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "9", "")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "aa", "a")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "eee", "ee")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "eea", "ea")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "df", "dof")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "hr", "har")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "llk", "lk")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "acl", "ack")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "ncl", "nk")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "bl", "l")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "tc", "t")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "scv", "sv")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "gc", "g")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "dp", "dap")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "ps", "sp")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "tyf", "t")
	chosenFirstName = strings.ReplaceAll(chosenFirstName, "wb", "w")
	if chosenFirstName[len(chosenFirstName)-1:] == "r" {
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "dr", "r")
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "rr", "r")
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "fr", "r")
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "tr", "r")
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "sr", "r")
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "pr", "r")
	}
	if chosenFirstName[len(chosenFirstName)-1:] == "q" {
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "q", "que")
	}
	if chosenFirstName[len(chosenFirstName)-1:] == "w" {
		chosenFirstName = strings.ReplaceAll(chosenFirstName, "w", "")
	}

	chosenSurname = strings.ReplaceAll(chosenSurname, "-", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "'", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, ".", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "1", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "2", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "3", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "4", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "5", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "6", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "7", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "8", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "9", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "hwp", "")
	chosenSurname = strings.ReplaceAll(chosenSurname, "dw", "d")
	chosenSurname = strings.ReplaceAll(chosenSurname, "crs", "s")
	chosenSurname = strings.ReplaceAll(chosenSurname, "crs", "s")
	chosenSurname = strings.ReplaceAll(chosenSurname, "crc", "crac")
	chosenSurname = strings.ReplaceAll(chosenSurname, "rwg", "rg")
	chosenSurname = strings.ReplaceAll(chosenSurname, "drs", "dors")
	chosenSurname = strings.ReplaceAll(chosenSurname, "prs", "pers")
	chosenSurname = strings.ReplaceAll(chosenSurname, "ps", "sp")
	chosenSurname = strings.ReplaceAll(chosenSurname, "sst", "ss")
	if chosenSurname[len(chosenSurname)-1:] == "r" {
		chosenSurname = strings.ReplaceAll(chosenSurname, "dr", "r")
		chosenSurname = strings.ReplaceAll(chosenSurname, "rr", "r")
		chosenSurname = strings.ReplaceAll(chosenSurname, "fr", "r")
		chosenSurname = strings.ReplaceAll(chosenSurname, "tr", "r")
		chosenSurname = strings.ReplaceAll(chosenSurname, "sr", "r")
	}
	if chosenSurname[len(chosenSurname)-1:] == "q" {
		chosenSurname = strings.ReplaceAll(chosenSurname, "q", "que")
	}
	if chosenSurname[len(chosenSurname)-1:] == "w" {
		chosenSurname = strings.ReplaceAll(chosenSurname, "w", "")
	}

	//result
	firstName := strings.Title(chosenFirstName)
	surName := strings.Title(chosenSurname)

	cimName := firstName + " " + surName

	fmt.Println(cimName)

}
