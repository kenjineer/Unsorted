def PigConvert(word):
    word = list(word)
    vowel = ['a', 'e', 'i', 'o', 'u']
    i = 0
    if word[i] in vowel:
        pigadd = ["way"]
    else:
        pigadd = ["ay"]
    while word[i] not in vowel:
        letter = word.pop(i)
        word.append(letter)
    pigword = word + pigadd
    return ''.join(pigword)


word = input("Enter a word to be converted in Pig Latin: ").lower()
print(PigConvert(word))