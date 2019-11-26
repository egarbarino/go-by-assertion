// # Input/Output: Files and Streams
// This section covers file access and input/output/error streams. A `tmpRoot` variable (e.g., pointing
// to `/tmp` in Linux) is assumed for all examples.
//
// Ignore-On
package files

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpRoot = os.TempDir() + string(os.PathSeparator)

// ## Creating a File
// The `Create()` function from the `os` package is
// equivalent to the Unix `touch` command.

func Test_Create_File(t *testing.T) {
	file, err := os.Create(tmpRoot + "example.txt")
	if err != nil {
		t.Error("Can't create file: ", err)
	}
	file.Close()

	// Assertions
	assert.Equal(t, nil, err)
}

// ## Obtaining a File's Metadata
// The `Stat()` function from the `os` package
// returns a `fileInfo` struct which has a number
// of methods that help inspect the file's metadata.

func Test_Get_Metadata(t *testing.T) {
	// Create example.txt
	Test_Create_File(t)

	// Get File Metadata
	fileInfo, err := os.Stat(tmpRoot + "example.txt")
	if err != nil {
		t.Error("Can't access file: ", err)
	}

	// Assertions
	assert.Equal(t, "example.txt", fileInfo.Name())
	assert.Equal(t, int64(0), fileInfo.Size())
	assert.Equal(t, os.FileMode(0x1b6), fileInfo.Mode())
	assert.Equal(t, true, fileInfo.ModTime().Year() >= 2019)
	assert.Equal(t, false, fileInfo.IsDir())
}

// ## Renaming a File
// The `Rename()` function from the `os` package is
// equivalent to the Unix `rename` command.
func Test_Rename(t *testing.T) {
	// Create example.txt
	Test_Create_File(t)

	// Rename example.txt to deleteme.txt
	err := os.Rename(tmpRoot+"example.txt", tmpRoot+"deleteme.txt")
	if err != nil {
		t.Error("Can't rename file: ", err)
	}

	// Assertions
	_, err = os.Stat(tmpRoot + "deleteme.txt") // Check file's presence
	assert.Equal(t, nil, err)
}

// ## Deleting a File
// The `Remove()` function from the `os` package is
// equivalent to the Unix `rm` command.
func Test_Delete(t *testing.T) {
	// Create example.txt
	Test_Create_File(t)

	// Delete File
	err := os.Remove(tmpRoot + "example.txt")
	if err != nil {
		t.Error("Can't delete file: ", err)
	}

	// The example.txt should be absent
	_, err = os.Stat(tmpRoot + "example.txt")

	// Assertions
	assert.NotEqual(t, nil, err)
}

// ## Writing Bytes to a File
// Data may be written to  file by using the `Write()` or `WriteString()` methods
// offered by the `File` struct.
func Test_Write_File(t *testing.T) {
	// Open File for Writing
	file, err := os.Create(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer file.Close()

	// Write File
	bytes := []byte("Hello 😀")
	count, err2 := file.Write(bytes) // WriteString() also available
	if err2 != nil {
		t.Error("Can't write to file: ", err2)
	}

	// Flush changes!
	file.Sync()

	// Assertions
	assert.Equal(t, 10, count) // bytes written!
}

// ## Reading Bytes from a File
// Data may be read from a file by using the `Read()` method
// offered by the `File` struct.
func Test_Read_File(t *testing.T) {
	// Create hello.txt
	Test_Write_File(t)

	// Open File
	file, err := os.Open(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer file.Close()

	// Successful Read
	bytes := make([]byte, 10)
	count, err2 := file.Read(bytes)
	if err2 != nil {
		t.Error("Can't open file: ", err2)
	}

	// EOF Error (No more bytes left to read)
	var oneByteSlice = []byte{0}
	_, err3 := file.Read(oneByteSlice)

	// Assertions
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, count)
	assert.Equal(t, io.EOF, err3)
}

// ## Writing Bytes to a File (Single Statement)
// The `WriteFile()` function from the `ioutil` package allows creating and
// writing to a file using a single statement.
func Test_Write_File_IOUtil_Quick(t *testing.T) {
	err := ioutil.WriteFile(tmpRoot+"hello.txt", []byte("Hello 😀"), 0666)
	if err != nil {
		t.Error("Can't open file: ", err)
	}
}

// ## Read Bytes from a File Without Allocating a Slice
// The `ReadAll()` function from the `ioutil` package returns a byte
// slice with the contents of a file, without requiring to allocate
// the slice in advance.
func Test_Read_File_IOUtil(t *testing.T) {

	// Create hello.txt
	Test_Write_File_IOUtil_Quick(t)

	// Open file for reading
	file, err := os.Open(tmpRoot + "hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read File
	bytes, err2 := ioutil.ReadAll(file) // results in bytes!
	if err2 != nil {
		t.Error("Can't open file: ", err2)
	}

	// Assertions
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, len(bytes))

}

// ## Reading Bytes form a File Without a File Handle
// The `ReadFile()` function from the `ioutil` package allows reading the
// bytes form a file using a single statement, without requiring the allocation
// of a slice in advance and neither the obtention of a file handle.
func Test_Read_File_IOUtil_Quick(t *testing.T) {

	// Create hello.txt
	Test_Write_File_IOUtil_Quick(t)

	// Open file and read bytes in one go!
	bytes, err := ioutil.ReadFile(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}

	// Assertions
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, len(bytes))

}

// ## Writing to a File with Buffering
// The `Write()` and `WriteString()` functions, for bytes and strings,
// respectively, from the `bufio` package, allow to write files with
// buffering.
func Test_Write_File_BufIO(t *testing.T) {

	// Open file For writing
	fileHandle, err := os.Create(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer fileHandle.Close()

	// Write File
	writer := bufio.NewWriter(fileHandle)
	count, err2 := writer.WriteString("Hello 😀")
	if err2 != nil {
		t.Error("Can't open file: ", err2)
	}

	// Make sure changes are written to disk!
	writer.Flush()

	// Assertions
	assert.Equal(t, 10, count)
}

// ## Reading from a File with Buffering
// The `Read()` and `ReadString()` functions, for bytes and strings,
// respectively, from the `bufio` package, allow to read files with
// buffering.
func Test_Read_File_BufIO(t *testing.T) {

	// Open file for reading
	file, err := os.Open(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer file.Close()

	// Read File
	reader := bufio.NewReader(file)
	line, err2 := reader.ReadString('o')
	if err2 != nil {
		t.Error("Can't open file: ", err2)
	}

	// EOF Error (No more lines left to read)
	_, err3 := reader.ReadString('o')

	// Assertions
	assert.Equal(t, "Hello", line)
	assert.Equal(t, io.EOF, err3)
}

// ## Reading Lines from a File
// The `NewScanner()` function, from the `bufio` package, allows iterating
// through the lines contained in a file.
func Test_BufIO_Scanner(t *testing.T) {

	// First generate example file
	contents := []byte("line1\nline2\nline3\n")
	err := ioutil.WriteFile(tmpRoot+"lines.txt", contents, 0666)
	if err != nil {
		t.Error("Can't open file: ", err)
	}

	// Open File For Reading
	file, err := os.Open(tmpRoot + "lines.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer file.Close()

	// Use Scanner
	input := bufio.NewScanner(file)
	var lineBuffer = ""
	for input.Scan() {
		lineBuffer = lineBuffer + input.Text()
	}

	// Assertions
	assert.Equal(t, "line1line2line3", lineBuffer)
}
