#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

// open the puzzle file and read the data
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

int calculate_distance(int **data, int rows) {
  int result = 0;

  for (int i = 0; i < rows; i++)
    // calculate distance between values and make it positive
    result += abs(data[i][1] - data[i][0]);

  return result;
}

int calculate_similarity(int **data, int rows) {
  // initialize result value
  int result = 0;

  // loop through rows
  for (int i = 0; i < rows; i++) {
    // get the left value and initialize how many appearances it has in the
    // right column
    int val = data[i][0];
    int appearances = 0;
    for (int j = 0; j < rows; j++) {
      // if left value = right value, +1
      if (val == data[j][1]) {
        appearances++;
      }
    }
    // add similarity
    result += val * appearances;
  }
  return result;
}

int main() {
  int **puzzle, **sample;
  int sample_rows, puzzle_rows;

  sample = open_puzzle("sample.txt", &sample_rows);
  smallest_to_largest(sample, sample_rows);
  int sample_result = calculate_distance(sample, sample_rows);
  printf("sample part1: %d\n", sample_result);

  puzzle = open_puzzle("puzzle.txt", &puzzle_rows);
  smallest_to_largest(puzzle, puzzle_rows);
  int puzzle_result = calculate_distance(puzzle, puzzle_rows);
  printf("puzzle part1: %d\n", puzzle_result);

  int sample_similarity = calculate_similarity(sample, sample_rows);
  printf("sample part2: %d\n", sample_similarity);

  int puzzle_similarity = calculate_similarity(puzzle, puzzle_rows);
  printf("puzzle part2: %d\n", puzzle_similarity);

  return 0;
}
