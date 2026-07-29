#include <iostream>
using namespace std;

class Task{
    int id;
    string name;
    bool completed = false;
    public:
    Task(){}
    Task(int id, string name){
        this->id = id;
        this->name = name;
    }
    void markCompleted(){
        completed = true;
    }
    string getName(){
        return name;
    }
    int getId(){
        return id;
    }
    bool getIsCompleted(){
        return completed;
    }
};


class TaskManager{
    unordered_map<int, unordered_map<int, Task>> userTasks;
    int cnt = 0;
    public:
    TaskManager(){}
    int addTask(int userId, string name){
        cnt++;
        Task task = Task(cnt, name);
        userTasks[userId][task.getId()] = task;
        return cnt;
    }

    void markComplete(int userId, int taskId){
        if(userTasks.find(userId)!= userTasks.end() && userTasks[userId].find(taskId) != userTasks[userId].end()){
            userTasks[userId][taskId].markCompleted();
        }
    }

    void deleteTask(int userId, int taskId){
        if(userTasks.find(userId)!= userTasks.end() && userTasks[userId].find(taskId) != userTasks[userId].end()){
            userTasks[userId].erase(taskId);
        }
    }

    void listTasks(int userId){
        if(userTasks.find(userId)!= userTasks.end()){
            cout<<"Tasks for user id: "<<userId<<endl;
            for(auto it: userTasks[userId]){
                cout<<it.second.getName()<<" isCompleted ->"<<it.second.getIsCompleted()<<endl;
            }
        }
    }
};
int main(){
    TaskManager tm;
    tm.addTask(1, "task1");
    tm.addTask(2, "task2");
    tm.addTask(1, "task3");
    tm.listTasks(1);
    tm.markComplete(1, 1);
    tm.listTasks(1);
    tm.deleteTask(1, 1);
    tm.listTasks(1);
    return 0;
}