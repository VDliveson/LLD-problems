#pragma once

#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <set>

using namespace std;

class Flight
{
    int id;
    int startTime;
    int endTime;
    string start;
    string end;

public:
    Flight() {}

    Flight(int id, int startTime, int endTime, string start, string end) {
        this->id = id;
        this->startTime = startTime;
        this->endTime = endTime;
        this->start = start;
        this->end = end;
    }

    int getId() { return id; }

    int getStartTime() { return startTime; }

    int getEndTime() { return endTime; }

    string getStart() { return start; }

    string getEnd() { return end; }

    void changeStartTime(int time)
    {
        startTime = time;
    }

    void changeEndTime(int time)
    {
        endTime = time;
    }
};