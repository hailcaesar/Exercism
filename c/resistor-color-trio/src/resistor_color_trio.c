#include "resistor_color_trio.h"
#include <stdio.h>
#include <math.h>

resistor_value_t
color_code(resistor_band_t* colors){
    int value_ans = 0;
    for(int i = 0; i < 2; i++){
        value_ans = value_ans * 10 + colors[i];
    }

    unit_t unit_ans;
    unit_ans = OHMS;
    if (colors[2] > 0){
        int mult = pow(10, colors[2]);
        value_ans = value_ans * mult;
    }
    
    if (value_ans > 1000){
        value_ans = value_ans / 1000;
        unit_ans = KILOOHMS;
    }

    return (resistor_value_t){value_ans, unit_ans};
}
