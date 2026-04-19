#pragma once
#include "slot.hpp"
#include <vector>
using namespace std;

class ParkingFloor
{
    private:
    int floorNumber;
    vector<ParkingSlot> parkingSlots;

public:
    ParkingFloor(int floorNumber, vector<ParkingSlot> parkingSlots) : parkingSlots(parkingSlots), floorNumber(floorNumber) {}

    int getFloorNumber() { return floorNumber; }
    vector<ParkingSlot> getParkingSlots() { return parkingSlots; }
    ParkingSlot* findAvailableSlot(string vehicleType) {
        for (auto &slot : parkingSlots) {
            if (!slot.getIsOccupied() && slot.getVehicleType() == vehicleType) {
                return &slot;
            }
        }
        return NULL;
    }

    void addParkingSlot(ParkingSlot slot) {
        parkingSlots.push_back(slot);
    }

    bool freeSlot(int slotNumber) {
        for (ParkingSlot &slot : parkingSlots) {
            if (slot.getSlotNumber() == slotNumber) {
                slot.removeVehicle();
                return true;
            }
        }
        return false;
    }
    
};