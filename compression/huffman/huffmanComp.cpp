/*
 * HuffMan Compression Algorithm
 * 10/25/2020
 * TODO:
 * Output properly (bits)
 */

#include <cassert>
#include <fstream>
#include <iostream>
#include <queue>
#include <string>
#include <unordered_map>
#include <vector>

// LeafNode Struct
struct LeafNode {
    char symbol;
    int weight;
    LeafNode *l, *r;
    LeafNode(char s, int weight) : symbol(s), weight(weight){};
    LeafNode(char s, int weight, LeafNode* l1, LeafNode* l2)
        : symbol(s), weight(weight), l(l1), r(l2){};
    friend std::ostream& operator<<(std::ostream& os, const LeafNode& ln);
};

// For making the Priority Queue. We Have the Lowest Frequency on Top
struct CompareWeight {
    bool operator()(const LeafNode* l1, const LeafNode* l2) {
        return l1->weight > l2->weight;
    };
};

// Debug Purposes
std::ostream& operator<<(std::ostream& os, const LeafNode* ln) {
    os << "{" << ln->symbol << ", " << ln->weight << "}" << std::endl;
    return os;
};

// Encodes the HuffMan Tree (Assigns 0,1 to Characters based on left or right
// pos)
void encode(const LeafNode* root, std::string str,
            std::unordered_map<char, std::string>& code) {
    if (root == nullptr) return;

    if (!root->l && !root->r) {
        code[root->symbol] = str;
    }
    encode(root->l, str + "0", code);
    encode(root->r, str + "1", code);
};

int main() {
    std::ifstream file("./text.txt");
    std::unordered_map<char, int> freq;
    char ch{};
    // Lambda has no default cntr so can't be used as comparison for Priority
    // Queue (Requires Type) Apparently I can use Decltype but Im not too sure
    // about any of that yet auto cmp = [](LeafNode const& l1, LeafNode const&
    // l2) { return l1.weight < l2.weight;};
    std::priority_queue<LeafNode*, std::vector<LeafNode*>, CompareWeight> pq;
    // Populate Frequencies
    if (file.is_open()) {
        while (file.get(ch)) {
            freq[ch]++;
        }
    }
    // Populate Leaf Nodes
    for (auto& e : freq) {
        pq.push(new LeafNode(e.first, e.second));
        std::cout << "{" << e.first << ", " << e.second << "}"
                  << "\n";
    }
    // Make HuffMan Tree
    while (pq.size() != 1) {
        LeafNode* l1 = pq.top();
        pq.pop();
        LeafNode* l2 = pq.top();
        pq.pop();
        pq.push(new LeafNode('\0', (l1->weight + l2->weight), l1, l2));
    }

    const LeafNode* top = pq.top();
    std::unordered_map<char, std::string> code;
    encode(top, "", code);

    file.close();
    std::ofstream output("./compressedText.txt",
                         std::ios::out | std::ios::binary);
    // Now We can Write the Encoded text
    std::ifstream f2("./text.txt");
    std::string test;
    ch = {};
    if (f2.is_open() && output.is_open()) {
        std::cout << "Both Open " << std::endl;
        while (f2.get(ch)) {
            std::cout << "Character : " << ch << std::endl;
            test += code[ch];
            output << code[ch];
        }
        std::cout << "Closing " << std::endl;
        output.close();
    }
    std::cout << "Size of String: " << test.size() << std::endl;
}

