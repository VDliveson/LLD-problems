#pragma once

#include <iostream>
#include <unordered_map>
#include "storage.h"

using namespace std;

class MapStorage : public Storage{
    unordered_map<string,int> mp;
    public:

    void add(string key, int value){
        mp[key] = value;
    }

    int get(string key){
        if(mp.find(key)==mp.end()) return -1;
        return mp[key];
    }

    void remove(string key){
        mp.erase(key);
    }

};