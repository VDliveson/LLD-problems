#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <set>

#include "counter.hpp"
#include "flight.hpp"
#include "employee/employee.hpp"
#include "employee/crew.hpp"
#include "employee/ground.hpp"
#include "luggage.hpp"
#include "counter.hpp"
#include "ticket.hpp"
#include "passenger.hpp"
#include "board.hpp"

class AirportSystem
{
    unordered_map<int, Flight *> _flight;
    unordered_map<int, Counter *> _counter;
    unordered_map<int, Employee *> _employee;
    unordered_map<int, int> _prices;
    unordered_map<int, Ticket *> _tickets;
    unordered_map<int, Luggage *> _luggage;
    unordered_map<int, Passenger *> _passenger;
    Board b;

public:
    AirportSystem() : b() {}

    int bookTicket(Counter *c, Passenger *p, Flight *f)
    {
        int flightId = f->getId();
        int cost = _prices[flightId];
        Ticket *ticket = c->createTicket(p->getId(), flightId, cost);
        int ticketId = ticket->getId();
        _tickets[ticketId] = ticket;
        return ticketId;
    };

    void addFlight(Flight *f)
    {
        int flightId = f->getId();
        _flight[flightId] = f;
        b.update(f);
    };

    void addCounter(Counter *c)
    {
        int counterId = c->getId();
        _counter[counterId] = c;
    };

    void addEmployee(Employee *e)
    {
        int employeeId = e->getID();
        _employee[employeeId] = e;
    };

    void addPassenger(Passenger *p)
    {
        int passengerId = p->getId();
        _passenger[passengerId] = p;
    };

    void addLuggage(Luggage *l)
    {
        int luggageId = l->getId();
        _luggage[luggageId] = l;
    };

    bool checkInPassenger(int ticketId, Counter *counter)
    {
        for (auto it : _tickets)
        {
            if (it.first == ticketId)
            {
                counter->checkInTicket(it.second);
                return true;
            }
        }
        return false;
    }
};