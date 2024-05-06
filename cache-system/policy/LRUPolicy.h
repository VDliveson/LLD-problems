#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <set>
#include "evictionPolicy.h"
#include "../algo/helper.h"

using namespace std;

class LRUPolicy : public EvictionPolicy
{
    unordered_map<string, Node *> mp;
    Node *head;
    Node *tail;
    Helper helper;

public:
    LRUPolicy() : head(new Node("-1")), tail(new Node("-1")), helper(head, tail) {}

        void AccessedKey(string key)
    {
        if (mp.find(key) != mp.end())
        {
            Node *node = mp[key];
            helper.removeNode(node);
            helper.addNodeAtEnd(node);
        }
        else
        {
            mp[key] = new Node(key);
            helper.addNodeAtEnd(mp[key]);
        }
    }

    string EvictKey()
    {
        string key = head->next->key;
        helper.removeNode(head->next);
        mp.erase(key);
        return key;
    }
};