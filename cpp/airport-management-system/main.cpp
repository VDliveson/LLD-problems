#include "airport-system.hpp"


#include <iostream>
using namespace std;

int main(){
    AirportSystem s;
    Passenger *p = new Passenger(12,"vd");
    s.addPassenger(p);
    Luggage *l = new Luggage(4,12);
    s.addLuggage(l);
    Flight *f = new Flight(1,12,14,"Delhi","Mumbai");
    s.addFlight(f);

    Crew *c = new Crew(1,"Raj",1);
    s.addEmployee(c);
    Ground *g = new Ground(1,"Raj");
    s.addEmployee(g);

    Counter *counter = new Counter(23);
    s.addCounter(counter);
    int ticketId = s.bookTicket(counter,p,f);
    cout<<"Booked ticket with id "<<ticketId<<endl;
    bool flag = s.checkInPassenger(ticketId,counter);
    if(!flag){
        cout<<"Unable to check in passenger with id "<<ticketId<<endl;
    }
    return 0;
}