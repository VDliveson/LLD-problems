#pragma once
#include "slot.hpp"
#include <vector>
using namespace std;

class Ticket{
    private:
    string ticketNumber;
    string vehicleNumber;
    int slotId;
    time_t timeOfEntry;
    time_t timeOfExit = -1;
public:
    Ticket(string ticketNumber, string vehicleNumber, int slotId) : ticketNumber(ticketNumber), vehicleNumber(vehicleNumber), slotId(slotId), timeOfEntry(time(0)) {}
    string getTicketNumber() { return ticketNumber; }
    string getVehicleNumber() { return vehicleNumber; }
    int getSlotId() { return slotId; }
    time_t getTimeOfEntry() { return timeOfEntry; }
    time_t getTimeOfExit() { return timeOfExit; }
    void setTimeOfExit() { timeOfExit = time(0); }
};