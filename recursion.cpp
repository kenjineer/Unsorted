#include <iostream>
using namespace std;

int pow(int base, int exp)
{
    return (exp > 0) ? base * pow(base, exp - 1) : 1;
}

int main(void)
{
    int base, exp;
        cout << "Enter the base number: ";
        cin >> base;
    loop:
        cout << "Enter the power of the number: ";
        cin >> exp;
        if (exp < 0)
        {
            cout << "OutOfBounds: Only positive exponents are accepted." << endl;
            goto loop;
        }
    cout << base << "^" << exp << " = " << pow(base, exp);
}