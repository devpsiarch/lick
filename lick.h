#ifndef LICK_H
#define LICK_H

#define inithead "ref: refs/heads/main\n"

#include <zlib.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <dirent.h>

void init(){
	struct stat s;
    int err = stat(".lick",&s);
	if(err != -1){
		printf(".lick repo already exists\n");
		return;
	}
	mkdir(".lick",0777);
    mkdir(".lick/objects",0777);
    mkdir(".lick/refs",0777);
    FILE* headfile;
    headfile = fopen(".lick/HEAD","w");
    fputs(inithead,headfile);
    fclose(headfile);	
} 


void cat_file(char *option,char *hash){
	if(!option || !hash){
		printf("invalid command : lick <cmd> <option> <target>\n");
		return;
	}
	//handle overflow later
	char buffer[64];	
	char dirname[2];
	strncpy(dirname,hash,2);
	dirname[2] = '\0';
	char *filename = hash+2;
	chdir(".lick/objects");
	if(chdir(dirname) == -1){
		printf("error opening folder : .lick/objects/%s\n",dirname);
		return;
	}
	FILE *file = fopen(filename,"rb");
	if(file == NULL){
		printf("error opening file\n");
		return;
	}
	int count = fread(buffer,sizeof(buffer),1,file);
	for(int i = 0 ; i < 64 ; i++){
		printf("%u ",buffer[i]);
	}
	
	char content[64];
	count = uncompress(content,(uLongf*)sizeof(content),buffer,(uLongf)sizeof(buffer));
	
}

#endif
