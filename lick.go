package main

import (
	"encoding/hex"
	"path/filepath"
	"crypto/sha1"
	"compress/zlib"
	"io"
	"os"
	"fmt"
	"strings"
	"bytes"
)

func bad(err error){
	if(err != nil){
		panic(err)
	}
}

func direxi(err error,path string){
	if(err != nil){
		fmt.Printf("the directory [%s] aleady exist, no panic!\n",path)
		return
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func clean_lick(){
	err := os.RemoveAll(".lick")
	bad(err)	
}

func init_lick(){
	err := os.Mkdir(".lick",700)
	direxi(err,".lick")

	err = os.Mkdir(".lick/objects",700)
	direxi(err,".lick/objects")	

	err = os.Mkdir(".lick/refs",700)
	direxi(err,".lick/refs")
	
	file,err := os.Create(".lick/HEAD")
	

	file.Close()
}

func cat_file(option,blob_sha string){
	//decompress the contents of a file
	//for now it has to be -p

	path := fmt.Sprintf(".lick/objects/%v/%v", blob_sha[0:2], blob_sha[2:])
	file,err := os.Open(path)
	if(err != nil){
		fmt.Println("cant open file ",path)
		return
	}
	r, err := zlib.NewReader(io.Reader(file))
	if(err != nil){
		fmt.Println("error while reading file");
		return
	}
	str, err := io.ReadAll(r)
	if(err != nil){
		fmt.Println("error while reading file");
		return
	}
	parts := strings.Split(string(str), "\x00")
	if(option == "-p"){
		fmt.Println(parts[1])
		return
	}
	parts = strings.Split(parts[0]," ")
	// p,1 => content || t,first of 0 || s,second of 0 |> still there is more option will be implimented later

	switch option {
	case "-t":
		fmt.Println(parts[0])
	case "-s":
		fmt.Println(parts[1])
	}
	r.Close()
	
}


func hash_object(option,FilePath string)([20]byte){
	if(option != "-w"){
		panic("wrong option only -w for now!!")
	}
	//getting the input of the sha1 fuc
	
	file,err := os.ReadFile(FilePath)
	bad(err)	

	stat,err := os.Stat(FilePath)
	bad(err)	

	content := string(file)
	header := fmt.Sprintf("blob %d\x00%s", stat.Size(), content)

	sha := (sha1.Sum([]byte(header)))
	hash := hex.EncodeToString(sha[:])


	blobpath := ".lick/objects/"

	for i, c := range hash {
		blobpath += string(c)
		if i == 1{
			blobpath += "/"
		}

	}

	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write([]byte(header))
	w.Close()

	os.MkdirAll(filepath.Dir(blobpath),os.ModePerm)
	bad(err)	
	
	File,err := os.Create(blobpath)
	bad(err)
	defer File.Close()

	File.Write(buf.Bytes())
	fmt.Println(hash)

	return sha 

}

func ls_tree(flage,tree_sha string){
	if(flage != "--name-only"){
		fmt.Println("invalid flage ")
		return
	}
	path := fmt.Sprintf(".lick/objects/%v/%v", tree_sha[0:2], tree_sha[2:])
	file,err := os.Open(path)
	if(err != nil){
		fmt.Println("cant open file ",path)
		return
	}
	r, err := zlib.NewReader(io.Reader(file))
	if(err != nil){
		fmt.Println("error while reading file");
		return
	}
	str, err := io.ReadAll(r)
	if(err != nil){
		fmt.Println("error while reading file");
		return
	}

	split := bytes.Split(str, []byte("\x00"))
	//there still an implimentation for displaying more stuff
	//i cant be bothered ...
	use := split[1 : len(split)-1]
    for _, dByte := range use {
        splitByte := bytes.Split(dByte, []byte(" "))[1]
	    fmt.Println(string(splitByte))
    }

}


func write_tree(dir string) {
    entries, err := os.ReadDir(dir)
    bad(err)

    for _, entry := range entries {
		if(entry.Name() == ".git" || entry.Name() == ".lick"){
			continue
		}
		fullPath := dir + "/" + entry.Name() // Construct the full path for the entry
        if entry.IsDir() {
            fmt.Println("directory :", entry.Name())
            write_tree(fullPath) // Now passing the full path
        } else {
            fmt.Println("file :", entry.Name())
        }
    }
}

func main(){
	if(len(os.Args) < 2){
		fmt.Println("invalid command : lick <command> <option>")
		return
	}
	switch os.Args[1] {
		case "init":
			//creating .lick 
			init_lick()
		case "clean":
			//cleaning .lick
			clean_lick()
		case "cat-file":
			//reading objects
			cat_file(os.Args[2],os.Args[3])
		case "hash-object":
			//writing object
			fmt.Sprintf("%x",hash_object(os.Args[2],os.Args[3]))
		case "ls-tree":
			//read tree objects
			ls_tree(os.Args[2],os.Args[3])
		case "write-tree":
			write_tree(".")
		default:
			panic("invalid command: lick <cmd> <parameters>")
	}




}
