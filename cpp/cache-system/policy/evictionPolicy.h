#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>

class EvictionPolicy {
    public:

    virtual void AccessedKey(string key){}

    virtual string EvictKey(){
        return "";
    } 
};