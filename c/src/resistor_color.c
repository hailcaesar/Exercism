#include "resistor_color.h"

static resistor_band_t cols[WHITE - BLACK + 1] = {0};

resistor_band_t* colors(){
    if (cols[0] != 0) return cols;
    
    for (int i = 0; i < WHITE - BLACK + 1; i++){
        cols[i] = (resistor_band_t)i;
    }
    return cols;
}

int color_code(resistor_band_t color){
    return (int)color;
}
