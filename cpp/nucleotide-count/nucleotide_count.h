#if !defined(NUCLEOTIDE_COUNT_H)
#define NUCLEOTIDE_COUNT_H
#include <map>
#include <string>
using namespace std;

namespace nucleotide_count
{
    class counter
    {
        map<char, int> nucleotideCounts{{'A', 0}, {'T', 0}, {'C', 0}, {'G', 0}};

    public:
        counter(string strand);
        int count(char a) const;
        map<char, int> nucleotide_counts() const;
    };
} // namespace nucleotide_count

#endif // NUCLEOTIDE_COUNT_H