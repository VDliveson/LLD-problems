#pragma once
#include "slot.hpp"
#include "floor.hpp"
#include "ticket.hpp"
#include "Vehicle/vehicleFactory.hpp"
#include <string>
#include <vector>

class ParkingLot
{
    private:
    vector<ParkingFloor> parkingFloors;
    vector<Ticket> tickets;
    int capacity;
public:
    ParkingLot(int capacity): capacity(capacity) {}
    void addParkingFloor(int floorNumber, unordered_map<string, int> vehicleTypes) {
        vector<ParkingSlot> parkingSlots;
        for (auto &vehicleType : vehicleTypes) {
            for (int i = 0; i < vehicleType.second; i++) {
                parkingSlots.push_back(ParkingSlot(rand(), vehicleType.first));
            }
        }
        ParkingFloor floor(floorNumber, parkingSlots);
        parkingFloors.push_back(floor);
    }

    bool addParkingSlot(int floorNumber, string vehicleType) {
        for(auto &floor: parkingFloors) {
            if(floor.getFloorNumber() == floorNumber && capacity > 0) {
                floor.addParkingSlot(ParkingSlot(rand(), vehicleType));
                capacity--;
                return true;
            }
        }
        return false;
    }

    string parkVehicle(string licensePlateNumber, string vehicleType) {
        Vehicle* vehicle = VehicleFactory::createVehicle(licensePlateNumber, vehicleType);
        if (vehicle == NULL) return "";

        for (auto &floor : parkingFloors) {
            ParkingSlot* slot = floor.findAvailableSlot(vehicleType);
            if (slot != NULL) {
                slot->parkVehicle(vehicle);
                string ticketNumber = to_string(rand());
                Ticket ticket(ticketNumber, vehicle->getLicensePlateNumber(), slot->getSlotNumber());
                tickets.push_back(ticket);
                return ticket.getTicketNumber();
            }
        }
        delete vehicle;
        return "";
    }

    bool unparkVehicle(string ticketNumber) {
        for (Ticket &ticket : tickets) {
            if (ticket.getTicketNumber() == ticketNumber) {
                for (ParkingFloor &floor : parkingFloors) {
                    if (floor.freeSlot(ticket.getSlotId())) {
                        ticket.setTimeOfExit();
                        return true;
                    }
                }
            }
        }
        return false;
    }

    int getLeftCapacity() {
        return capacity;
    }
};