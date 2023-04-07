package main

import (
	"fmt"
	"os"

	"github.com/andreyvit/diff"
	gobible "github.com/gobible/gobible"
	bible "github.com/gobible/gobible/bible"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: gobible-compare <file1> <file2>")
		os.Exit(1)
	}

	file1 := args[0]
	file2 := args[1]

	// check if files exist

	verifyFiles(file1, file2)

	fmt.Println("Comparing", file1, "and", file2, "\n")

	b1 := gobible.New(file1)
	b2 := gobible.New(file2)
	bOut := gobible.New(file1)

	bOut.Version.Name = b1.Version.Name + " compared to " + b2.Version.Name
	bOut.Version.Abbrev = b1.Version.Abbrev + "vs" + b2.Version.Abbrev

	fmt.Println("\n\nComparing", b1.Version.Name, "and", b2.Version.Name, "\n")
	fmt.Println("We will use the books of the Protestant Canon to compare the two bibles. \n")

	for _, book := range bible.BooksTable {
		fail := false
		b1b := b1.GetBook(book.Name)
		b2b := b2.GetBook(book.Name)

		if b1b == nil {
			fmt.Println("error >", book.Name, "does not exist in", b1.Version.Name)
			fail = true
		}
		if b2b == nil {
			fmt.Println("error >", book.Name, "does not exist in", b2.Version.Name)
			fail = true
		}

		if fail {
			continue
		}

		for ci, b1c := range b1b.Chapters {
			b2c := b2b.GetChapter(b1c.Number)
			if b2c == nil {
				fmt.Println("error > Book", book.Name, "Chapter", b1c.Number, "does not exist in", b2.Version.Name)
				continue
			}

			if len(b1c.Verses) != len(b2c.Verses) {
				fmt.Println("warning> Book", book.Name, "Chapter", b1c.Number, "does not have the same number of verses in both bibles")
			}

			for vi, b1v := range b1c.Verses {
				b2v := b2c.GetVerse(b1v.Number)
				if b2v == nil {
					fmt.Println("error> ", b1v.Number, "does not exist in", b2.Version.Name)
					continue
				}

				d := diff.CharacterDiff(b1v.Text, b2v.Text)
				//fmt.Println("Book", bi, "Chapter", ci, "Verse", vi)
				bOut.GetBook(b1b.Name).Chapters[ci].Verses[vi].Text = d
			}

		}

	}

	fmt.Println("Saving output to out.json")

	err := bOut.Save("out.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func verifyFiles(file1, file2 string) {
	if file1 == file2 {
		fmt.Println("Files are the same")
	}

	if !fileExists(file1) {
		fmt.Println(file1 + " does not exist")
		os.Exit(1)
	}

	if !fileExists(file2) {
		fmt.Println(file2 + " does not exist")
		os.Exit(1)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
