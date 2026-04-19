#pragma once
#include "vehicle.hpp"

class Bike: public Vehicle{
    private:
    
    public:
    Bike(string licensePlateNumber) : Vehicle(licensePlateNumber) {}
    string getType() override { return "bike"; }
    string getLicensePlateNumber() override { return licensePlateNumber; }
};