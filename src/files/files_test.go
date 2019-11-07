// # Files
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

var (
	tmpRoot = os.TempDir() + string(os.PathSeparator)

//	fileHandle *os.File
//	fileInfo   os.FileInfo
//	err        error
)

// ## Create, Get Metadata, Rename, Copy, and Delete File
func Test_Create_Verify_Delete(t *testing.T) {

	// Create file (equivalent to "touch")
	fileHandle, err := os.Create(tmpRoot + "example.txt")
	if err != nil {
		t.Error("Can't create file: ", err)
	}
	fileHandle.Close()

	// Get File Metadata
	fileInfo, err2 := os.Stat(tmpRoot + "example.txt")
	if err2 != nil {
		t.Error("Can't access file: ", err2)
	}
	assert.Equal(t, "example.txt", fileInfo.Name())
	assert.Equal(t, int64(0), fileInfo.Size())
	assert.Equal(t, os.FileMode(0x1b6), fileInfo.Mode())
	assert.Equal(t, true, fileInfo.ModTime().Year() >= 2019)
	assert.Equal(t, false, fileInfo.IsDir())

	// Rename File
	err3 := os.Rename(tmpRoot+"example.txt", tmpRoot+"deleteme.txt")
	if err3 != nil {
		t.Error("Can't delete file: ", err3)
	}

	// Delete File
	err4 := os.Remove(tmpRoot + "deleteme.txt")
	if err4 != nil {
		t.Error("Can't delete file: ", err4)
	}

}

// ## OS: Write File
func Test_Write_File(t *testing.T) {
	// Open File for Writing
	fileHandle, err := os.Create(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer fileHandle.Close()
	// Write File
	bytes := []byte("Hello 😀")
	count, err2 := fileHandle.Write(bytes) // WriteString() also available
	if err2 != nil {
		t.Error("Can't write to file: ", err2)
	}
	fileHandle.Sync() // Flush changes
	assert.Equal(t, 10, count)
}

// ## OS: Read File
func Test_Read_File(t *testing.T) {

	Test_Write_File(t)

	// Open File
	fileHandle, err := os.Open(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer fileHandle.Close()

	// Successful Read
	bytes := make([]byte, 10)
	count, err2 := fileHandle.Read(bytes)
	if err2 != nil {
		t.Error("Can't open file: ", err2)
	}
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, count)

	// EOF Error (No more bytes left to read)
	var oneByteArray = []byte{0}
	count, err3 := fileHandle.Read(oneByteArray)
	assert.Equal(t, io.EOF, err3)
}

// ## IOUtil: Write File (Quick)
func Test_Write_File_IOUtil_Quick(t *testing.T) {
	err := ioutil.WriteFile(tmpRoot+"hello.txt", []byte("Hello 😀"), 0666)
	if err != nil {
		t.Error("Can't open file: ", err)
	}
}

// ## IOUtil: Read File (Using File Handle)
func Test_Read_File_IOUtil(t *testing.T) {

	Test_Write_File_IOUtil_Quick(t)

	// Open file for reading
	fileHandle, err := os.Open(tmpRoot + "hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read File
	bytes, err2 := ioutil.ReadAll(fileHandle)
	if err2 != nil {
		t.Error("Can't open file: ", err2)
	}
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, len(bytes))

}

// ## IOUtil: Read File (Quick)
func Test_Read_File_IOUtil_Quick(t *testing.T) {

	Test_Write_File_IOUtil_Quick(t)

	bytes, err := ioutil.ReadFile(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, len(bytes))

}

// ## BufIO: Write File
func Test_Write_File_BufIO(t *testing.T) {

	// Open File For Writing
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
	writer.Flush()
	assert.Equal(t, 10, count)
}

// ## BufIO: Read File
func Test_Read_File_BufIO(t *testing.T) {

	// Open File For Reading
	fileHandle, err := os.Open(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer fileHandle.Close()

	// Read File
	reader := bufio.NewReader(fileHandle)
	line, err2 := reader.ReadString('o')
	if err2 != nil {
		t.Error("Can't open file: ", err2)
	}
	assert.Equal(t, "Hello", line)

	// EOF Error (No more lines left to read)
	_, err3 := reader.ReadString('o')
	assert.Equal(t, io.EOF, err3)
}

// ## BufIO: Scanner
func Test_BufIO_Scanner(t *testing.T) {

	// First generate example file
	contents := []byte("line1\nline2\nline3\n")
	err := ioutil.WriteFile(tmpRoot+"lines.txt", contents, 0666)
	if err != nil {
		t.Error("Can't open file: ", err)
	}

	// Open File For Reading
	fileHandle, err := os.Open(tmpRoot + "lines.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer fileHandle.Close()

	// Use Scanner
	input := bufio.NewScanner(fileHandle)
	var lineBuffer = ""
	for input.Scan() {
		lineBuffer = lineBuffer + input.Text()
	}
	assert.Equal(t, "line1line2line3", lineBuffer)
}
