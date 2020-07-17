#include "armstrong_numbers.h"
#include <math.h>

int int_len(int input){
    int len = 1;
    while (input >= 10){
        len++;
        input = input / 10;
    }

    return len;
}


bool is_armstrong_number(int candidate){
    int len = int_len(candidate);
    int result = 0;
    int input = candidate;

    while (candidate > 0){
        int digit = candidate % 10;
        result += pow(digit, len);
        candidate /= 10;
    }

    return result == input;   
}

