word_set = [
   "apple red round fruit yellow green",
   "banana yellow long fruit  SWEET",
   "orange orange round fruit SOUR AND SWEET",
   "grape purple small fruit SOUR VIOLET  Favourite"
]

#Expected ouput
 
#"apple red" : ['round', 'fruit', 'Blue', 'green']
#"banana yellow" : ['long', 'fruit',  'SWEET']
#"orange orange" : ['round', 'fruit', 'SOUR', 'AND', 'SWEET']
#"grape purple" : ['small', 'fruit' , 'SOUR', 'VIOLET',  'Favourite']

for words in word_set:
    word = words.split(" ")
    print(word[0], word[1])
    print(word[2:])

def process_words(word_set):
    pass
    

process_words(word_set)