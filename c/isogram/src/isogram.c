#include "isogram.h"

#include <stdio.h>
#include <search.h>
#include <stddef.h>
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

bool is_isogram(const char phrase[]){
    if(phrase == NULL) return false;
    if (strlen(phrase) == 0) return true;

    hdestroy();
    hcreate(100);

    ENTRY e, *found;
    for (int idx = 0; phrase[idx] != '\0'; idx++){
        if(strncmp(&phrase[idx], "-", 1) == 0 || strncmp(&phrase[idx], " ", 1) == 0){
            continue;
        }
        e.key = (char*) malloc(2);//strdup(&c);
        e.key[0] = tolower(phrase[idx]);
        e.key[1] = '\0';
        found = hsearch(e, FIND);

        if (found != NULL){
            hdestroy();
            return false;
        }

        ENTRY* entered = hsearch(e, ENTER);
        if(!entered){
            printf("ERROR!!");
            hdestroy();
            return false;
        }
    }

    hdestroy();
    return true;

}
