#ifndef LICK_H
#define LICK_H

#define inithead "ref: refs/heads/main\n"
#include <stdio.h>
#include <string.h>
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
	char dirname[2];
	strncpy(dirname,hash,2);
	dirname[2] = '\0';
	char *filename = hash+2;
	chdir(".lick/objects");
	if(chdir(dirname) == -1){
		printf("error opening folder in objects\n");
		return;
	}
}

#endif
