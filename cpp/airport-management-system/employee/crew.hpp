#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>

#include "employee.hpp"
using namespace std;

class Crew : public Employee
{
    int id;
    string name;
    int flightId;
    public:
    Crew();

    Crew(int id, string name, int flightId){
        this->id = id;
        this->name = name;
        this->flightId = flightId;
    }

    string getName(){
        return name;
    }

    int getID(){
        return id;
    }

    int getAirplaneID(){
        return flightId;
    }

};