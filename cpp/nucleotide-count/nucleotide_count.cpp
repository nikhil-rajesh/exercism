#include <map>
#include "nucleotide_count.h"
#include <string>
using namespace std;

namespace nucleotide_count
{
    counter::counter(string strand)
    {
        for (const auto &c : strand)
        {
            nucleotideCounts[c]++;
        }
    }

    int counter::count(char a) const
    {
        return nucleotideCounts.at(a);
    }

    map<char, int> counter::nucleotide_counts() const
    {
        return nucleotideCounts;
    }

} // namespace nucleotide_count
