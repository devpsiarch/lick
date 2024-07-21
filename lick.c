#include "lick.h"


int main(int argc,char **argv){
	if(argc < 2){
		perror("invalid command : lick <command>");	
		return 0;
	}
	if(strcmp(argv[1],"init") == 0) init();	
	if(strcmp(argv[1],"cat-file") == 0) cat_file(argv[2],argv[3]);	

	return 0;
}
