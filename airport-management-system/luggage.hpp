#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>

class Luggage
{
    int id;
    int passengerId;

public:
    Luggage() {}

    Luggage(int id, int passengerId)
    {
        this->id = id;
        this->passengerId = passengerId;
    }

    int getId(){
        return this->id;
    }

    int getPassengerId(){
        return this->passengerId;
    }

};