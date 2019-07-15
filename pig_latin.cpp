#include <iostream>
#include <algorithm>
using namespace std;

string PigLatin(string word)
{
    string vowel = "aeiou";
    string pigadd;
    string letter;
    string pigword = word;
    int i = 0;

    if (vowel.find(word[i]) != string::npos)
        pigadd = "way";
    else pigadd = "ay";
    for (i = 0; i < word.length(); i++)
    {
        if (vowel.find(word[i]) != string::npos) 
        {
            break;
        }
        else
        {
            letter = word.at(i);
            pigword.erase(0, 1);
            pigword.append(letter);
        }
    }
    return pigword.append(pigadd);
}

int main(void)
{
    string word;

    cout << "Enter a word to be converted into Pig Latin: ";
    cin >> word;
    transform(word.begin(), word.end(), word.begin(), ::tolower);
    cout << PigLatin(word);    
}