#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>
using namespace std;

class Employee{
    string name = "";
    int id = -1;
    public:

    virtual string getName(){
        return name;
    }

    virtual int getID(){
        return id;
    }

    virtual int getAirplaneID(){
        return -1;
    }
};