package file

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Reader struct {
	io.Reader
	Total   float64
	Current float64
}

func (r *Reader) Read(b []byte) (n int, err error) {
	n, err = r.Reader.Read(b)
	r.Current += float64(n)
	fmt.Printf("当前进度:\n%.2f%%", r.Current/r.Total*100)
	return
}

func DownLoadFilePercent(url, fileName string) {

	//readUrl, err := http.Get(url)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//readUrl.ContentLength()

	file, err := ioutil.ReadFile(url)
	if err != nil {
		return
	}

	create, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = create.Close()
	}()
	//readObj := Reader{
	//	Reader: bytes.NewReader(file),
	//	Total:  float64(len(file)),
	//}
	//n, err := readObj.Reader.Read(file)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("n", n)
	io.Copy(create, bytes.NewReader(file))

}

func DownLoadFile(url, fileName string) {

	getResponse, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = getResponse.Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	written, err := io.Copy(file, getResponse.Body)
	fmt.Println(written, err)
}
