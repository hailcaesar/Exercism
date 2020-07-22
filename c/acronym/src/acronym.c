#include "acronym.h"
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

int  count_words(const char* phrase){
    int count = 1; 
    int len = strlen(phrase);
    
    for(int idx = 0; idx < len; idx++){
        if(isalnum(phrase[idx]) && 
          (isspace(phrase[idx-1]) || phrase[idx-1] == '-')) count++;
    }

    return count;
}

char* abbreviate(const char *phrase){
    if(phrase == NULL || strlen(phrase) == 0) 
        return NULL;
    
    int words = count_words(phrase);
    int len = strlen(phrase);
    char* ans = (char*) calloc(words+1, sizeof(char));
    int ptr = 0;

    for(int idx = 0; idx < len; idx++){
        if(isalnum(phrase[idx]) && 
          (isspace(phrase[idx-1]) || idx == 0 || phrase[idx-1] == '-')){
            ans[ptr++] = toupper(phrase[idx]);
        }
    }

    return ans;
}
