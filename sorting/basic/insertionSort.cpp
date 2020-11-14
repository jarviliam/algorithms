/*
 * Basic Insertion Sort
 */

#include <array>
#include <cassert>
#include <iostream>

void insertionSort(int* arr, std::size_t size) {
    int key, j;
    for (int i = 1; i < size; i++) {
        key = arr[i];
        j = i - 1;
        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
}

int main() {
    int arr1[] = {5, 2, 3, 7, 9, 1};
    int arr2[] = {1, 2, 3, 4, 3, 2, 1};

    insertionSort(arr1, sizeof(arr1) / sizeof(arr1[0]));
    for (auto x : arr1) {
        std::cout << x << std::endl;
    }
    insertionSort(arr2, sizeof(arr2) / sizeof(arr2[0]));
    for (auto x : arr2) {
        std::cout << x << std::endl;
    }
}
