#include <math.h>
#include "darts.h"

uint8_t score(coordinate_t lp){
    uint8_t ans = 0;
    float distance = sqrt(pow(lp.x, 2) + pow(lp.y, 2));

    if (distance <= 1)
        ans = 10;
    else if (distance <= 5)
        ans = 5;
    else if (distance <= 10)
        ans = 1;

    return ans;
}
