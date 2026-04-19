#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>

#include "node.h"

class Helper
{
    Node *head;
    Node *tail;

public:
    Helper() {}

    Helper(Node *&head, Node *&tail)
    {
        this->head = head;
        this->tail = tail;
        head->next = tail;
        tail->prev = head;
    }

    void addNodeAtEnd(Node *node)
    {
        Node *prevNode = tail->prev;
        prevNode->next = node;
        node->prev = prevNode;
        node->next = tail;
        tail->prev = node;
    }

    void removeNode(Node *node)
    {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }
};
