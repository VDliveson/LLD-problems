#pragma once
#include <iostream>
#include "Vehicle/vehicle.hpp"
using namespace std;


class ParkingSlot
{
    private:
    int slotNumber;
    bool isOccupied;
    string vehicleType;
    Vehicle* parkedVehicle;

public:
    ParkingSlot(int slotNumber, string vehicleType) : slotNumber(slotNumber), isOccupied(false), vehicleType(vehicleType) {}
    int getSlotNumber() { return slotNumber; }
    bool getIsOccupied() { return isOccupied; }
    string getVehicleType() { return vehicleType; }
    bool parkVehicle(Vehicle* vehicle) {
        if (!isOccupied && vehicle->getType() == vehicleType) {
            parkedVehicle = vehicle;
            isOccupied = true;
            return true;
        }
        return false;
    }

    void removeVehicle() {
        delete parkedVehicle;
        parkedVehicle = NULL;
        isOccupied = false;
    }

};
