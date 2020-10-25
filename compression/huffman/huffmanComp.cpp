#include <cassert>
#include <iostream>
#include <vector>
#include <string>
#include <fstream>
#include <unordered_map>
#include <queue>

struct LeafNode {
    private:
        std::string symbol;
        double weight;
        LeafNode *parent;
};

int main () {
    std::cout << "Test" << std::endl;
    std::ifstream file ("./text.txt");
    std::string line;
    std::unordered_map<char,int> freq;
    std::priority_queue<LeafNode> pq;
    char ch{};
    if(file.is_open()){
        while(file.get(ch)){
            std::cout << ch << "\n";
        }
    }
}
