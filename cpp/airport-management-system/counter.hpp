#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>
#include "ticket.hpp"
using namespace std;

class Counter{
    int id;
    public:

    Counter(){}

    Counter(int id){
        this->id = id;
    }

    int getId(){
        return this->id;
    }

    Ticket* createTicket(int customerId, int flightId, int price){
        int ticketId = rand();
        return new Ticket(ticketId, customerId, flightId, price);
    }

    void checkInTicket(Ticket* ticket){
        ticket->checkIn();
        cout << "Ticket " << ticket->getId() << " has been checked in." << endl;
    }

};