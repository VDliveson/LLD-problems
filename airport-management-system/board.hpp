#pragma once

#include <iostream>
#include <unordered_map>
#include "flight.hpp"
using namespace std;

class Board{
    unordered_map<int,Flight*> mp;
    public:

    Board(){}
    
    void display(){
        for(auto it:mp){
            Flight *f = it.second;
            cout<<" Flight: "<<f->getId()<<" from "<<f->getStart()<<" to "<<f->getEnd()<<" departs at : "<<f->getStartTime()<<endl;
        }
    }

    void update(Flight *f){
        mp[f->getId()] = f;
        display();
    }

};