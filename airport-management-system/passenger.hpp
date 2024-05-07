#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>
using namespace std;

class Passenger{
    int id;
    string name;

    public:

    Passenger(int id, string name){
        this->id = id;
        this->name = name;
    }

    int getId(){
        return this->id;
    }

    string getName(){
        return this->name;
    }    

};