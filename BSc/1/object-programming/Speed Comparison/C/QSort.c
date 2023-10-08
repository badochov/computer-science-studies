#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void qs(int tab[], int start, int size)
{
    if (size <= 1)
    {
        return;
    }
    int d = rand() % size + start;
    int smallerCount = 0;
    int biggerCount = 0;
    int smaller[size];
    int bigger[size];
    for (int i = start; i < start + size; i++)
    {
        if (i != d)
        {
            if (tab[i] < tab[d])
            {
                smaller[smallerCount] = tab[i];
                smallerCount++;
            }
            else
            {
                bigger[biggerCount] = tab[i];
                biggerCount++;
            }
        }
    }
    tab[start + smallerCount] = tab[d];
    for (int i = 0; i < smallerCount; i++)
    {
        tab[start + i] = smaller[i];
    }
    for (int i = 0; i < biggerCount; i++)
    {
        tab[start + smallerCount + 1 + i] = bigger[i];
    }
    qs(tab, start, smallerCount);
    qs(tab, start + smallerCount + 1, biggerCount);
}

void quickSort(int tab[], int size)
{
    qs(tab, 0, size);
}

void printArray(int tab[], int size)
{
    for (int i = 0; i < size; i++)
    {
        printf("%d ", tab[i]);
    }
    printf("\n\r");
}
int main()
{
    srand(time(NULL));
    FILE *f = fopen("./test.txt", "r");
    if (f == NULL)
    {
        exit(1);
    }
    int size;
    fscanf(f, "%d", &size);
    int tab[size];
    for (int i = 0; i < size; i++)
    {
        fscanf(f, "%d", &tab[i]);
    }
    quickSort(tab, size);
    // printArray(tab, size);
    return 0;
}