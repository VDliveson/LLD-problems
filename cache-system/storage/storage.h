#pragma once

#include <iostream>
using namespace std;

class Storage{
    public:

    virtual void add(string key,int value){}

    virtual void remove(string key){}

    virtual int get(string key){
        return 0;
    }
};