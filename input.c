#include <stdio.h>
#include <stdlib.h>
#define MAXLENGTH 120

char nextChar(FILE *fpi, FILE *fpo, int lineIndex)
{
    printf("called\n");
    static char line[MAXLENGTH];

    if (lineIndex == -1)
    {
        if (fgets(line, MAXLENGTH, fpi) != NULL)
        {
            fputs(line, fpo);
            lineIndex = 0;
            /* code */
        }
        else
        {
            exit(1);
        }
    }
    char ch = line[lineIndex];
    lineIndex++;
    if (ch == '\n')
    {
        lineIndex = -1;
        return ' ';
        /* code */
    }
    return ch;
}

int main()
{
    char outputFile[] = "output.c";
    char inputFile[] = "input.c";
    static FILE *fpi;
    fpi = fopen(inputFile, "u");
    static FILE *fpo;
    fpo = fopen(outputFile, "w");
    int index = -1;
    char c = nextChar(fpi, fpo, index);
    printf("%c", c);
    return 1;
}
