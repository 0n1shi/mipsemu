int sum = 0;

int main(void) {
    int i = 1;
    while (i <= 100) {
        sum = sum + i;
        i++;
        if (i > 100) break;
    }
    return sum;
}
