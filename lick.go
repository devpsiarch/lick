package main

import (
	"crypto/sha1"
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


func hash_object(option,filepath string){
	if(option != "-w"){
		fmt.Println("wrong option only <-w> for now!")
		return
	}
	//getting the input of the sha1 fuc
	
	file,err := os.ReadFile(filepath)
	if(err != nil){
		fmt.Println("unable to read file :",filepath)
		return
	}

	stat,err := os.Stat(filepath)
	if(err != nil){
		fmt.Println("unable to get stats of file :",filepath)
		return
	}
	content := string(file)
	header := fmt.Sprintf("blob %d\x00%s", stat.Size(), content)

	sha := (sha1.Sum([]byte(header)))
	hash := fmt.Sprintf("%x",sha)
	fmt.Println("the hash is :",hash)

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
			hash_object(os.Args[2],os.Args[3])
	}




}
