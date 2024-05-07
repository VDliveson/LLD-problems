#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>
using namespace std;

class Ticket{
    int id;
    int passengerId;
    int flightId;
    int price;
    bool isCheckedin = false;

    public:

    Ticket(){}
    
    Ticket(int id, int passengerId, int flightId, int price){
        this->id = id;
        this->passengerId = passengerId;
        this->flightId = flightId;
        this->price = price;
    }

    int getId(){
        return this->id;
    }

    int getPassengerId(){
        return this->passengerId;
    }

    int getFlightId(){
        return this->flightId;
    }

    int getPrice(){
        return this->price;
    }

    void setPrice(int price){
        this->price = price;
    }

    void checkIn(){
        if(!isCheckedin){
            isCheckedin = true;
        }
    }

};