#pragma once

#include <iostream>
using namespace std;

class Vehicle
{
    protected:
    string licensePlateNumber;

public:
    Vehicle(string licensePlateNumber) : licensePlateNumber(licensePlateNumber) {}
    virtual ~Vehicle() = default;
    virtual string getLicensePlateNumber() = 0;
    virtual string getType() = 0;
};