#include "hamming.h"
#include <string>
using namespace std;

namespace hamming {
    int compute(string a, string b) {
        int dist = 0;
        for(int i=0; i<int(a.length()); i++) {
            if(a[i] != b[i])
                dist++;
        }  
        return dist;
    }
}  // namespace hamming
