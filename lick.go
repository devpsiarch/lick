package main

import (
	"compress/zlib"
	"io"
	"os"
	"fmt"
	"strings"
)



func clean_lick(){
	err := os.RemoveAll(".lick")
	if(err != nil){
		fmt.Println("error cleaning .lick directory")
		return
	}
}

func init_lick(){
	err := os.Mkdir(".lick",700)
	if (err != nil){
		fmt.Println("error creating directory !<mother>")
		return
	}
	err = os.Mkdir(".lick/objects",700)
	if (err != nil){
		fmt.Println("error creating directory ! <objects>")
		return
	}
	err = os.Mkdir(".lick/refs",700)
	if (err != nil){
		fmt.Println("error creating directory ! <refs>")
		return
	}
	
	file,err := os.Create(".lick/HEAD")
	if(err != nil){
		fmt.Println("error creating file ! <HEAD>")
		return
	}
	file.Close()
}

func cat_file(option,blob_sha string){
	//decompress the contents of a file
	//for now it has to be -p

	path := fmt.Sprintf(".git/objects/%v/%v", blob_sha[0:2], blob_sha[2:])
	fmt.Println("the path is ",path)
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
	// p,1 => content || t,first of 0 || s,second of 0 |> still there is more option will be implimented later

	switch option {
	case "-p": 
		fmt.Print(parts[1])
	case "-t":
		fmt.Print(parts[0])
	case "-s":
		fmt.Print(parts[0])
	}
	r.Close()
	
}

func main(){
	if(len(os.Args) < 2){
		fmt.Println("invalid command : lick <command> <option>")
		return
	}
	//creating .lick and deleting it 
	if(os.Args[1] == "init"){
		init_lick()
	}
	if(os.Args[1] == "clean"){
		clean_lick()
	}
	//lick cat-file -p <hash>
	if(os.Args[1] == "cat-file"){
		cat_file(os.Args[2],os.Args[3])	
	}





}
