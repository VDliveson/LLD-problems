#pragma once
#include "vehicle.hpp"
#include "car.hpp"
#include "bike.hpp"
#include <string>
using namespace std;

class VehicleFactory {
public:
    static Vehicle* createVehicle(const string& licensePlateNumber, const string& type) {
        if (type == "car") {
            return new Car(licensePlateNumber);
        } else if (type == "bike") {
            return new Bike(licensePlateNumber);
        }
        return nullptr;
    }
};
