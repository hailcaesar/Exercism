#include "resistor_color_duo.h"

int color_code(resistor_band_t* colors){
    int result = 0;
    for(int i = 0; i < 2; i++){
        result = result * 10 + colors[i];
    }

    return result;
}
