#include <stdio.h>

void quickSort(int in[], int low, int high) {
  int i, j, pivot, tmp;

  if (low < high) {
    pivot = low;
    i = low;
    j = high;

    while (i < j) {
      while (in[i] <= in[pivot] && i < high) {
        i++;
      }
      while (in[j] > in[pivot]) {
        j--;
      }
      if (i < j) {
        tmp = in[i];
        in[i] = in[j];
        in[j] = tmp;
      }
    }
    tmp = in[pivot];
    in[pivot] = in[j];
    in[j] = tmp;
    quickSort(in, low, j - 1);
    quickSort(in, j + 1, high);
  }
}

int main() {
  int in[] = {100, 1, 44, 22, 55, 323, 92, 12, 15, 48};
  int count = sizeof(in) / sizeof(in[0]);
  quickSort(in, 0, count);

  for (int i = 0; i < count; i++) {
    printf("Value %d\n", in[i]);
  }
}
