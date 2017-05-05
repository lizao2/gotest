package main

import (
	"bufio"
	"bytes"
	"example"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
)

func dropCR(data []byte) []byte {

	if len(data) > 0 && data[len(data)-1] == '\r' {

		return data[0 : len(data)-1]

	}

	return data

}

func Scanspace(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ' '); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil

	}
	// If we're at EOF, we have a final, non-terminated line. Return it.

	if atEOF {
		return len(data), dropCR(data), nil
	}

	// Request more data.
	return 0, nil, nil

}

func main() {
	log.Print("enter main")
	//	test scanner
	var sc *bufio.Scanner

	if f, err := os.Open("test.txt"); err != nil {
		log.Print("open failed ", err)
		return
	} else {
		sc = bufio.NewScanner(f)
		defer f.Close()
	}

	sc.Split(Scanspace)
	for sc.Scan() {
		text := sc.Text()
		log.Print(text)
		log.Print("read onece")
	}

	return
	//  end test scanner

	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}

	log.Print(test)

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	log.Print(data)

	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Print(newTest)

	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	log.Print("leave main")
}
