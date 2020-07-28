#include "pangram.h"
#include <string.h>
#include <stdio.h>
#include <ctype.h>

bool is_pangram(const char *sentence){
    if(sentence == NULL) return false;
    bool exists[29] = {false};
    int len = strlen(sentence);
    int count = 0;

    for(int idx = 0; idx < len; idx++){
        int char_idx = tolower(sentence[idx]) - 'a';
        if(char_idx >= 0 && char_idx <= 25 && !exists[char_idx]){
            exists[char_idx] = true;
            count++;
        }
    }

    return count == 26;
}
