#include <iostream>
#include "lot.hpp"
using namespace std;

int main()
{
    ParkingLot parkingLot(100);
    unordered_map<string, int> vehicleTypes;
    vehicleTypes["car"] = 1;
    vehicleTypes["bike"] = 2;
    parkingLot.addParkingFloor(1, vehicleTypes);
    string vehicleNumber = "KA-01-1234";
    string type = "car";
    string ticket1 = parkingLot.parkVehicle(vehicleNumber, type);
    if (!ticket1.empty())
    {
        cout << "Vehicle "<< vehicleNumber << " parked with ticket number: " << ticket1 << endl;
    }
    else
    {
        cout << "Parking lot is full!" << endl;
    }

    string vehicleNumber2 = "KA-01-5678";   
    string type2 = "car";
    string ticket2 = parkingLot.parkVehicle(vehicleNumber2, type2);
    if (!ticket2.empty())
    {
        cout << "Vehicle "<< vehicleNumber2 << " parked with ticket number: " << ticket2 << endl;
    }
    else
    {
        cout << "Parking lot is full!" << endl;
    }
    bool unParked = parkingLot.unparkVehicle(ticket1);
    if (unParked) {
        cout << "Vehicle with ticket number " << ticket1 << " has been unparked" << endl;
    } else {
        cout << "Failed to unpark vehicle with ticket number " << ticket1 << endl;
    }

    string ticket3 = parkingLot.parkVehicle(vehicleNumber2, type2);
    if (ticket3 != "") {
        cout << "Vehicle "<< vehicleNumber2 << " parked with ticket number: " << ticket3 << endl;
    } else {
        cout << "Parking lot is full!" << endl;
    }

    return 0;
}