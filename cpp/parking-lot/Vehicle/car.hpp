#pragma once
#include "vehicle.hpp"

class Car: public Vehicle{
    private:
    
    public:
    Car(string licensePlateNumber) : Vehicle(licensePlateNumber) {}
    string getType() override { return "car"; }
    string getLicensePlateNumber() override { return licensePlateNumber; }
};