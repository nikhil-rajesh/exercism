#include "space_age.h"

namespace space_age {
    space_age::space_age(long int seconds) {
        _seconds = seconds;
        noOfEarthYears = double(seconds)/31557600;
    }
    long int space_age::seconds() const {
        return _seconds;
    }
    double space_age::on_mercury() const {
        return noOfEarthYears/planets.at("Mercury");
    }
    double space_age::on_venus() const {
        return noOfEarthYears/planets.at("Venus");
    }
    double space_age::on_earth() const {
        return noOfEarthYears;
    }
    double space_age::on_mars() const {
        return noOfEarthYears/planets.at("Mars");
    }
    double space_age::on_jupiter() const {
        return noOfEarthYears/planets.at("Jupiter");
    }
    double space_age::on_saturn() const {
        return noOfEarthYears/planets.at("Saturn");
    }
    double space_age::on_uranus() const {
        return noOfEarthYears/planets.at("Uranus");
    }
    double space_age::on_neptune() const {
        return noOfEarthYears/planets.at("Neptune");
    }
}  // namespace space_age
