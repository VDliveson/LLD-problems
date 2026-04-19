#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>

class Node
{
public:
    string key;
    Node *next;
    Node *prev;

    Node(string key)
    {
        this->key = key;
        this->next = NULL;
        this->prev = NULL;
    }
};