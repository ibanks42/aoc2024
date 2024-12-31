#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

// Function to open the puzzle file and read the data
int **open_puzzle(const char *filename, int *rows) {
  FILE *fptr;
  int **arr = NULL;
  int cols = 2;
  int i;
  char line[256];

  fptr = fopen(filename, "r");

  if (fptr == NULL) {
    printf("Error opening file!\n");
    return NULL;
  }

  *rows = 0;

  while (fgets(line, sizeof(line), fptr) != NULL) {
    arr = (int **)realloc(arr, (*rows + 1) * sizeof(int *));
    arr[*rows] = (int *)malloc(cols * sizeof(int));

    sscanf(line, "%d %d", &arr[*rows][0], &arr[*rows][1]);
    (*rows)++;
  }

  fclose(fptr);

  return arr;
}

void smallest_to_largest(int **data, int rows) {
  for (int i = 0; i < rows; i++) {
    // loop through all values
    // loop again through all left and rightvalues and see if i is less than j.
    // if so, swap i and j.
    for (int j = 0; j < i; j++) {
      if (data[i][0] < data[j][0]) {
        int swap = data[i][0];
        data[i][0] = data[j][0];
        data[j][0] = swap;
      }

      if (data[i][1] < data[j][1]) {
        int swap = data[i][1];
        data[i][1] = data[j][1];
        data[j][1] = swap;
      }
    }
  }
}

int calculate(int **data, int rows) {
  int result = 0;

  for (int i = 0; i < rows; i++)
    result += abs(data[i][1] - data[i][0]);

  return result;
}

int main() {
  int **puzzle, **sample;
  int sample_rows, puzzle_rows;

  sample = open_puzzle("sample.txt", &sample_rows);
  smallest_to_largest(sample, sample_rows);
  int sample_result = calculate(sample, sample_rows);
  printf("sample part1: %d\n", sample_result);

  puzzle = open_puzzle("puzzle.txt", &puzzle_rows);
  smallest_to_largest(puzzle, puzzle_rows);
  int puzzle_result = calculate(puzzle, puzzle_rows);
  printf("puzzle part1: %d\n", puzzle_result);

  return 0;
}