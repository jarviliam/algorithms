/**
 * Merge Sort
 */
#include <stdio.h>

void merge(int in[], int left, int middle, int right) {
  int i, j, k;
  int lInd = middle - left + 1;
  int rInd = right - middle;
  int L[lInd], R[rInd];
  for (i = 0; i < lInd; i++)
    L[i] = in[left + i];
  for (j = 0; j < rInd; j++)
    R[j] = in[middle + 1 + j];
  i = 0;
  j = 0;
  k = left;

  while (i < lInd && j < rInd) {
    if (L[i] <= R[j]) {
      in[k] = L[i];
      i++;
    } else {
      in[k] = R[j];
      j++;
    }
    k++;
  }
  while (i < lInd) {
    in[k] = L[i];
    i++;
    k++;
  }
  while (j < rInd) {
    in[k] = R[j];
    j++;
    k++;
  }
}

void mergeSort(int in[], int l, int r) {
  if (l < r) {
    int middle = l + (r - l) / 2;
    mergeSort(in, l, middle);
    mergeSort(in, middle + 1, r);
    merge(in, l, middle, r);
  }
}

int main(int argc, char *argv[argc + 1]) {
  int in[] = {99, 33, 66, 11, 2, 64, 213};
  int size = sizeof(in) / sizeof(in[0]);
  mergeSort(in, 0, size - 1);

  for (int i = 0; i < size; i++) {
    printf("Value: %d\n", in[i]);
  }
  return 0;
}
