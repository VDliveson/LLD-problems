#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>

#include "employee.hpp"
using namespace std;

class Ground : public Employee
{
    int id;
    string name;
    public:
    Ground();

    Ground(int id, string name){
        this->id = id;
        this->name = name;
    }

    string getName(){
        return name;
    }

    int getID(){
        return id;
    }

};