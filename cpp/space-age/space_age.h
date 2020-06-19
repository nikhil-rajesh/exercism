#if !defined(SPACE_AGE_H)
#define SPACE_AGE_H
#include <string>
#include <map>
using namespace std;

namespace space_age
{
    class space_age
    {
        const map<string, double> planets {
            {"Mercury", 0.2408467},
            {"Venus", 0.61519726},
            {"Earth", 1.0},
            {"Mars", 1.8808158},
            {"Jupiter", 11.862615},
            {"Saturn", 29.447498},
            {"Uranus", 84.016846},
            {"Neptune", 164.79132}
        };
        double noOfEarthYears;
        long int _seconds;

    public:
        space_age(long int seconds);
        long int seconds() const;
        double on_mercury() const;
        double on_venus() const;
        double on_earth() const;
        double on_mars() const;
        double on_jupiter() const;
        double on_saturn() const;
        double on_uranus() const;
        double on_neptune() const;
    };
} // namespace space_age

#endif // SPACE_AGE_H