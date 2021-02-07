#define NULL (void *)0

struct student {
    int no;
    char name[16];
    int year;
    char student_class[16];
};

struct student mike = {0, "Mike", 13, "a1"};

char new_name[16] = "John";

char *strcpy(char *dst, const char *src);

int main(void) {
    mike.no = 10;
    strcpy(mike.name, new_name);
    return 0;
}

char *strcpy(char *dst, const char *src) {
    if (dst == NULL) return NULL;
    char *ptr = dst;
    while (*src != '\0')
    {
        *dst = *src;
        dst++;
        src++;
    }
    *dst = '\0';
    return ptr;
}
