#include <iostream>
#include <unordered_map>


using namespace std;

#include "policy/LRUPolicy.h"
#include "storage/mapStorage.h"
#include "Cache.h"

int main(){
    EvictionPolicy* policy = new LRUPolicy();
    Storage* storage = new MapStorage();
    Cache cache(policy, storage,3);
    cache.put("1", 1);
    cache.put("2", 2);
    cache.put("3",3);
    cout<<cache.get("1")<<endl;
    cache.put("4",4);
    cout<<cache.get("2")<<endl;
    return 0;
}