#include "difference_of_squares.h"
#include <cmath>

namespace difference_of_squares {
    int square_of_sum(int n) {
        return pow((n*(n+1)/2),2);
    }
    int sum_of_squares(int n) {
        return n*(n+1)*(2*n + 1)/6;
    }
    int difference(int n) {
        return pow((n*(n+1)/2),2) - n*(n+1)*(2*n + 1)/6;
    }
}  // namespace difference_of_squares
