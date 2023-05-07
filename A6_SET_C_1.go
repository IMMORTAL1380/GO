/*
ASSIGNMENT 6
SET C 
1. WAP in Go language to read a XML file into structure and display structure
*/
/*
--------------------------------------------------------------------
The xml package includes Unmarshal() function that supports decoding data from a byte slice into values. The xml.Unmarshal function is used to decode the values from the XML formatted file into a Notes struct.

The notes.xml file is read with the ioutil.ReadFile() function and a byte slice is returned, which is then decoded into a struct instance with the xml.Unmarshal() function. The struct instance member values are used to print the decoded data.
------------------------------------------------------------------
*/

<note>
<to>Tove</to>
<from>Jani</from>
<heading>Reminder</heading>
<body>Don't forget me this weekend!</body>
</note>
package main
 
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)
 
type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}
 
func main() {
	data, _ := ioutil.ReadFile("notes.xml")
 
	note := &Notes{}
 
	_ = xml.Unmarshal([]byte(data), &note)
 
	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)
}