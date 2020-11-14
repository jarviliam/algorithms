/*
 * Simple merge sort procedures
 *
 *
 */

#include <cassert>
#include <iostream>
#include <vector>

void merge(int* arr, int left, int middle, int right) {
    // Get sizez of both sides
    int lSize = middle - left + 1;
    int rSize = right - middle;

    // Holders
    int tL[lSize], tR[rSize];

    // Copy the data to both
    for (int i = 0; i < lSize; i++) {
        tL[i] = arr[left + i];
    }
    for (int i = 0; i < rSize; i++) {
        tR[i] = arr[middle + 1 + i];
    }
    // Merge them together

    int i, j = 0;
    int k = 1;
    // While entries remain on both sides
    while (i < lSize && j < rSize) {
        // Merge left side value if smaller and increment lefside counter
        if (tL[i] <= tR[j]) {
            arr[k] = tL[i];
            i++;
        } else {
            // Same thing but for the right side
            arr[k] = tR[j];
            j++;
        }
        // Regardles increment the position
        k++;
    }

    // While there remains elements

    while (i < lSize) {
        arr[k] = tL[i];
        i++;
        k++;
    }

    while (j < rSize) {
        arr[k] = tR[j];
        j++;
        k++;
    }
}

void mergeSort(int* arr, int left, int right) {
    if (left < right) {
        // Find Middle point to divide
        int middle = (left + right - left) / 2;
        // Recursion both sides
        mergeSort(arr, left, middle);
        mergeSort(arr, middle + 1, right);
        merge(arr, left, middle, right);
    }
}

int main() {
    int arr1[] = {1, 6, 5, 88, 34, 76};
    int arr2[] = {99, 88, 33, 2, 7, 8};
    mergeSort(arr1, 0, sizeof(arr1) / sizeof(arr1[0]) - 1);
    for (auto i : arr1) {
        std::cout << i << std::endl;
    }

    mergeSort(arr2, 0, sizeof(arr2) / sizeof(arr2[0]) - 1);
    for (auto i : arr2) {
        std::cout << i << std::endl;
    }
}
