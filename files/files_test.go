package files

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tmpRoot    = os.TempDir() + string(os.PathSeparator)
	fileHandle *os.File
	fileInfo   os.FileInfo
	err        error
)

// ## Create, Get Metadata, Rename, Copy, and Delete File
func Test_Create_Verify_Delete(t *testing.T) {

	// Create file (equivalent to "touch")
	fileHandle, err = os.Create(tmpRoot + "example.txt")
	if err != nil {
		t.Error("Can't create file: ", err)
	}
	fileHandle.Close()

	// Get File Metadata
	fileInfo, err = os.Stat(tmpRoot + "example.txt")
	if err != nil {
		t.Error("Can't access file: ", err)
	}
	assert.Equal(t, "example.txt", fileInfo.Name())
	assert.Equal(t, int64(0), fileInfo.Size())
	assert.Equal(t, os.FileMode(0x1b6), fileInfo.Mode())
	assert.Equal(t, true, fileInfo.ModTime().Year() >= 2019)
	assert.Equal(t, false, fileInfo.IsDir())

	// Rename File
	err = os.Rename(tmpRoot+"example.txt", tmpRoot+"deleteme.txt")
	if err != nil {
		t.Error("Can't delete file: ", err)
	}

	// Delete File
	err = os.Remove(tmpRoot + "deleteme.txt")
	if err != nil {
		t.Error("Can't delete file: ", err)
	}

}

// ## OS: Write File
func Test_Write_File(t *testing.T) {
	fileHandle, err = os.Create(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer fileHandle.Close()
	var bytes = []byte("Hello 😀")
	var count int
	count, err = fileHandle.Write(bytes) // WriteString() also available
	if err != nil {
		t.Error("Can't write to file: ", err)
	}
	assert.Equal(t, 10, count)
}

// ## OS: Read File
func Test_Read_File(t *testing.T) {

	Test_Write_File(t)

	// Open File
	fileHandle, err = os.Open(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	defer fileHandle.Close()

	// Successful Read
	var bytes []byte = make([]byte, 10)
	var count int
	count, err = fileHandle.Read(bytes)
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, count)

	// EOF Error (No more bytes left to read)
	var oneByteArray = []byte{0}
	count, err = fileHandle.Read(oneByteArray)
	assert.Equal(t, io.EOF, err)
}

// ## IOUtil: Write File
func Test_Write_File_IOUtil(t *testing.T) {
	err := ioutil.WriteFile(tmpRoot+"hello.txt", []byte("Hello 😀"), 0666)
	if err != nil {
		t.Error("Can't open file: ", err)
	}
}

// ## IOUtil: Read File (Using File Handle)
func Test_Read_File_IOUtil(t *testing.T) {

	Test_Write_File_IOUtil(t)

	// Open file for reading
	fileHandle, err = os.Open(tmpRoot + "hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read File
	bytes, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, len(bytes))

}

// ## IOUtil: Read File (Quick)
func Test_Read_File_IOUtil_Quick(t *testing.T) {

	Test_Write_File_IOUtil(t)

	bytes, err := ioutil.ReadFile(tmpRoot + "hello.txt")
	if err != nil {
		t.Error("Can't open file: ", err)
	}
	assert.Equal(t, "Hello 😀", string(bytes))
	assert.Equal(t, 10, len(bytes))

}

/*
// ## BufIO: Write File
func Test_Read_File_IOUtil_Quick(t *testing.T) {
	io.Wr
	bufio.NewWriter()
}
*/
