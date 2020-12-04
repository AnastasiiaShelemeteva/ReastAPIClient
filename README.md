# RestAPIClient
Writing a Client for REST API

With this Client you can connect to JokeAPI (https://sv443.net/jokeapi/v2/) and call it from your terminal. 
By default you will get a joke from category "Programming" in English. 

To chage the language, use flag -l. You can choose from English (en), German (de), Czech (cs) and Spanish (es). 
For example: ./RestAPIClient -l de

To see the list of avaliable categories, use flag -showall.
For example:./RestAPIClietn -showall

To change the category, use flag -cat. 
For example: ./RestAPI -cat Any
You can choose more than one category at the same time. To do so, you need to write a list of Categories, seperated by commas, without whitespaces. 
For example: ./RestAPI -cat Christmas,Dark,Programming

