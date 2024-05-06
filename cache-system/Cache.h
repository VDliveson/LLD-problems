#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <algorithm>

#include "storage/mapStorage.h"
#include "policy/LRUPolicy.h"

class Cache
{
public:
    EvictionPolicy *policy;
    Storage *storage;
    int capacity;

    int current = 0;
    Cache() {}

    Cache(EvictionPolicy *cpolicy, Storage *cstorage, int cap) : policy(cpolicy), storage(cstorage), capacity(cap) {}

    int get(string key)
    {
        if (storage->get(key) == -1)
            return -1;
        int val = storage->get(key);
        policy->AccessedKey(key);
        return val;
    }

    void put(string key, int value)
    {
        if (current == capacity)
        {
            string key = policy->EvictKey();
            cout << "Removing key " << key << endl;
            storage->remove(key);
            current--;
        }
        storage->add(key, value);
        policy->AccessedKey(key);

        current++;
    }

    ~Cache()
    {
        delete policy;
        delete storage;
    }
};