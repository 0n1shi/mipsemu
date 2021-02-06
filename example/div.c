int c = 10;
int d = 20;

int main(void) {
    int a = 15;
    int b = 3;
    c = div(a, b);
    d = mod(a, b);
    return 0;
}

int div(int a, int b) {
    return a / b;
}

int mod(int a, int b) {
    return a % b;
}
